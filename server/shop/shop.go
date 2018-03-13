package shop

import(
	"sync"
	"fmt"

	"github.com/funkeyfreak/vending-machine-api/etc"
	"encoding/json"
)

// driver lock
var dmu sync.Mutex
var drivers = make(map[string]Driver, 32)

type Shop struct {
	Pay 		`json:"pay"`
	Vend		`json:"vend"`
	Session		`json:"session"`
	Inventory 	`json:"inventory"`
	Category    `json:"cateogry"`
}

type Service interface {
	GetCategories()	([]CategoriesObj, error)
	SearchProducts(query InventoryQuery) ([]InventoryObj, error)
}

// Allows us to dynamically inject drivers, yo
func (s *Shop) UnmarshalJSON(b []byte) error {
	var conf struct {
		Drivers   	map[string]json.RawMessage 	`json:"drivers"`
		Inventory   string                      `json:"inventory"`
		Session    	string                       `json:"session"`
		Vend   		string                           `json:"vend"`
		Pay   		string                            `json:"pay"`
		Category 	string                         `json:"category"`
		Services 	map[string]json.RawMessage  	`json:"services"`
	}
	if err := json.Unmarshal(b, &conf); err != nil {
		return fmt.Errorf("unable to unmarshal shop config: %v", err)
	}
	for name, blob := range conf.Drivers {
		driver, ok := drivers[name]
		if !ok {
			return fmt.Errorf("bad shop config: no such driver: %s", name)
		}
		if err := json.Unmarshal(blob, driver); err != nil {
			return fmt.Errorf("unable to unmarshal config for shop driver %s: %v", name, err)
		}
	}

	var err error
	// pick the driver with the given name and then use it to update the
	// current shop. close over err to make an error-handling chain.
	withDriver := func(name string, fn func(Driver) error) {
		if err != nil {
			return
		}
		driver := drivers[name]
		if driver == nil {
			err = fmt.Errorf("no such driver: %s", name)
			return
		}
		err = fn(driver)
	}

	withDriver(conf.Vend, func(d Driver) error { return d.Vend(s) })
	withDriver(conf.Session, func(d Driver) error { return d.Session(s) })
	withDriver(conf.Pay, func(d Driver) error { return d.Pay(s) })
	withDriver(conf.Inventory, func(d Driver) error { return d.Inventory(s) })
	withDriver(conf.Category, func(d Driver) error { return d.Category(s) })

	return err
}


type Driver interface {
	Vend(*Shop) error
	Session(*Shop) error
	Pay(*Shop) error
	Inventory(*Shop) error
	Category(*Shop) error
}

func Open(driver string) (*Shop, error) {
	dmu.Lock()
	defer dmu.Unlock()

	d, ok := drivers[driver]
	if !ok {
		return nil, fmt.Errorf("no such driver: %s", driver)
	}

	return New(d)
}


func RegisterDriver(name string, d Driver) {
	dmu.Lock()
	defer dmu.Unlock()

	drivers[name] = d
}

//TODO: handle testing of all interfaces through services - set errors to nil or something
func New(d Driver) (*Shop, error) {
	s := new(Shop)
	err := etc.MergeErrors(
		d.Category(s),
		d.Inventory(s),
		d.Pay(s),
		d.Session(s),
		d.Vend(s),
		d.Devices(s),
	)
	return s, err
}








