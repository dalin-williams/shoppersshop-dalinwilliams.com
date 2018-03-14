package shop

import "github.com/satori/go.uuid"

type CategoriesObj struct {
	Id 			uuid.UUID    			`json:"id"    bson:"_id,omitempty"`
	Name		string					`json:"name"  bson:"name"`
	ToService	map[string]uuid.UUID	`json:"to_service" bson:"-"`
}

type Category interface {

}