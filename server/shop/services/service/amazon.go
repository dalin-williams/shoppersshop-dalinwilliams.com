package amazon

import (
	"os"
	"fmt"
	"net/url"
	"strings"
	"strconv"
	"container/heap"


	"github.com/pkg/errors"
	"github.com/dominicphillips/amazing"

	"github.com/funkeyfreak/vending-machine-api/server/shop"
	"github.com/funkeyfreak/vending-machine-api/caching"
	"github.com/funkeyfreak/vending-machine-api/etc"
)

type Amazon struct {
	AmazonAccessName string `json:"amazon_access_name"`
	AmazonAccessKey  string `json:"amazon_access_key"`
	AmazonSecretKey	 string `json:"amazon_secret_key"`

	client *amazing.Amazing
}

var (
	responseGroups = []string{
		"Request",
		"ItemIds",
		"Small",
		"Medium",
		"Large",
		"Offers",
		"OfferFull",
		"OfferSummary",
		"OfferListings",
		"PromotionSummary",
		"Variations",
		"VariationImages",
		"VariationSummary",
		"VariationMatrix",
		"VariationOffers",
		"ItemAttributes",
		"Tracks",
		"Accessories",
		"EditorialReview",
		"SalesRank",
		"BrowseNodes",
		"Images",
		"Similarities",
		"Reviews",
		"PromotionalTag",
		"AlternateVersions",
		"Collections",
		"ShippingCharges",
	}
	//TODO: Make dynamic
	categories = []string{
		"Automotive & Powersports",
	}
	categoryFile = "resources/amazon.json"
	ErrSearchNameMissing = "cannot create a search without name"
	ErrCreatingAmazonClient = "error in creating service client"
	dfsMem = map[string]string{} //will only exist from runtime to runtime
)

func init(){
	//load categories -
	//TODO: save in mongodb
	if _, err := os.Stat(categoryFile); os.IsNotExist(err) {
		// path/to/whatever does not exist
	}
}


func (a *Amazon) newClient() (*amazing.Amazing, error) {

	if a.client != nil {
		return a.client, nil
	}
	client, err := amazing.NewAmazing("US", a.AmazonAccessName, a.AmazonAccessKey, a.AmazonSecretKey)
	if err != nil {
		return nil, err
	}
	a.client = client
	return client, nil
}


func _toInventoryObj(item amazing.AmazonItem) (shop.InventoryObj){
	vsm := make([]string, len(item.ItemLinks))
	for idx, item := range item.ItemLinks {
		vsm[idx]= item.URL
	}

	res := shop.InventoryObj{
		Name: item.ItemAttributes.Title,
		ASIN: item.ASIN,
		Cost: shop.PayObj{
			Payment: float64(item.ItemAttributes.ListPrice.Amount),
			Currency: shop.PayCurrency{
				Name: "USD",
				Country: item.ItemAttributes.ListPrice.CurrencyCode,
				Format: item.ItemAttributes.ListPrice.FormattedPrice,
			},
		},
		Url: item.DetailPageURL,
		Resources: shop.InventoryResource{
			Images: vsm,
		},
	}
	return res

}

func _helperMapCategoriesToVisited(visited map[string]bool, categories []shop.CategoriesObj)(error){
	for _, item := range categories  {
		visited[item.Name] = true
	}
	return nil
}



// We use a bfs/dfs approach to find browser nodes we haven't yet seen. We only record the parent and root nodes
// by leveraging a dual visited approach, we are able to reduce possible duplicate traversals of the tree -
// if we traverse over any node towards a root, we record it - if we come across it again during the runtime
// of this app, we can auto-save the root
// TODO: Create a heap to get the
func (a *Amazon) bfsFindRootBrowserNode(node amazing.AmazonBrowseNode)(error){
	if node.Ancestors == nil  {
		return nil
	}

	pQueue := make(etc.PriorityQueue, len(node.Ancestors))
	pQIdx  := 0
	priority :=0

	queue := make(map[amazing.AmazonBrowseNode][]amazing.AmazonBrowseNode)


	//init map of bNodeAncestors
	for _, bNode := range node.Ancestors {
		queue[bNode] = bNode.Ancestors


		pQueue[0] =  &etc.Item{
			bNode,
			priority,
			pQIdx,
		}
		pQIdx++
	}

	heap.Init(&pQueue)

	top := node

	//register the parent
	categories, err := a.GetCategories()
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("trying to find root browser node %v - unable to get current categories", node))
	}

	visited := map[string]bool{top.Name: true}
	seen := []string{}

	//OPTIMIZE: Make this a bit better... && not as greedy. Maybe keep a list of these values around...
	_helperMapCategoriesToVisited(visited, categories)

	for len(queue) != 0 {
		//raise up the priority
		priority+=1

		//we've stored this, continue on
		if visited[top.Name] {
			caching.LocalCache.InsertIntoCache("amazon", "categories", shop.CategoriesObj{nil, top.Name, nil})
			item := heap.Pop(&pQueue).(*etc.Item)
			top = item.Value.(amazing.AmazonBrowseNode)
			delete(queue, top)
			continue
		}
		//we've ran across a node we've seen in our dfs algo
		if dfsMem[top.Name] != "" {
			visited[top.Name] = true
			caching.LocalCache.InsertIntoCache("amazon", "categories", shop.CategoriesObj{nil, top.Name, nil})
			item := heap.Pop(&pQueue).(*etc.Item)
			top = item.Value.(amazing.AmazonBrowseNode)
			delete(queue, top)
			continue
		}
		//Add all of top to queue
		if top.Ancestors != nil {
			//this will skip the first time -- we've initialized the queue to the children of the Ancestors already
			for idx, _ := range top.Ancestors {
				//if we don't have it stored, store it. If not "feget about et"
				if queue[top.Ancestors[idx]] == nil {
					tmpItem := &etc.Item{
						top.Ancestors[idx],
						priority,
						nil,
					}
					heap.Push(&pQueue, tmpItem)
					queue[top.Ancestors[idx]] = top.Ancestors[idx].Ancestors
					seen = append(seen, top.Ancestors[idx].Name)
				}
			}
		} else {
			//because we use a priority queue, we can "assume" seen will only contain nodes in a particular stack
			//this is cheaper than "checking" to see which queue we are in
			visited[top.Name] = true
			for idx, _ := range seen {
				dfsMem[seen[idx]] = top.Name
			}
			//clear seen
			seen = []string{}
		}


		//set top to the first element after deleting it from the queue
		//FIXME: consolidate this logic
		item := heap.Pop(&pQueue).(*etc.Item)
		top = item.Value.(amazing.AmazonBrowseNode)
		delete(queue, top)
	}
	return nil
}

func (a *Amazon) SearchProducts(query shop.InventoryQuery) ([]shop.InventoryObj, error){
	c, err := a.newClient()
	if err != nil {
		return nil, errors.WithMessage(err, ErrCreatingAmazonClient)
	}
	params := url.Values{

		"SearchIndex":   []string{"All"},
		"Condition": []string{"New"},
		"Operation":     []string{"ItemSearch"},
		"ResponseGroup": []string{strings.Join(responseGroups, ",")},
	}
	if query.Name == string(nil) {
		return nil, errors.New(ErrSearchNameMissing)
	} else {
		params.Add("Keywords", url.QueryEscape(query.Name))
	}

	if query.Price != float64(nil) {
		params.Add("MaximumPrice", url.QueryEscape(strconv.FormatFloat(query.Price, 'E', -1, 32)))
	}

	if query.Categories != nil {
		//TODO: Create category searching by BrowserNodes
		/*s := strings.Builder{}
		for _, value := range query.Categories {
			s.Grow(len(value.Name)+1)
			s.WriteString(value.Name+",")
		}
		//trim the excess ","
		//params.Add("BrowseNode", url.QueryEscape(query.Categories[0].Name))
		params.Add("SearchIndex", url.QueryEscape(strings.Trim(s.String(), ",")))*/

		params.Add("SearchIndex", url.QueryEscape(query.Categories[0].Name))
	}

	res, err := c.ItemSearch(params)
	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("error searching for item with params %v", params))
	}

	results := make([]shop.InventoryObj, len(res.AmazonItems.Items))
	for idx, val := range res.AmazonItems.Items {
		results[idx] = _toInventoryObj(val)
		//save categories
		//TEST: I have not added this feature because it has not been tested
		/*for i := 0; i < len(val.BrowseNodes); i++ {
			a.bfsFindRootBrowserNode(val.BrowseNodes[i])
		}*/
	}

	return results, nil
}


func (a *Amazon) GetCategories() ([]shop.CategoriesObj, error) {
	var categories = []shop.CategoriesObj{}

	err := caching.LocalCache.FetchObj("amazon", categories, []string{"categories"})
	if err != nil {
		return nil, errors.WithMessage(err, "could not parse categories object")
	}

	return 	categories, nil
}