package models

import (
	//"github.com/funkeyfreak/vending-machine-api/server/shop"
	"github.com/satori/go.uuid"
	"time"
)

type User struct {
	Id 			uuid.UUID	`json:"id"`
	UserName 	string		`json:"user_name"`
	FirstName	string		`json:"first_name"`
	LastName	string		`json:"last_name"`
	Email		string		`json:"email"`
	Password    string		`json:"password"`
	Phone		string		`json:"phone"`
	UserStatus	string		`json:"user_status"`
	IsActive	bool		`json:"is_active"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

func (u *User) Create(*User) error {

}

func (u *User) Read(*User) error {

}

func (u *User) Update(*User, func(*User)) error {

}

func (u *User) Delete(*User) error {

}