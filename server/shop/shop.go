package shop

import(
	"sync"
)

var lock sync.Mutex

type Shop struct {
	Pay 		`json:"pay"`
	Vend		`json:"vend"`
	Session		`json:"session"`
	Inventory 	`json:"inventory"`
}






