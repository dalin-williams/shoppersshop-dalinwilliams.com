package models

import (
	"time"
	"github.com/satori/go.uuid"
)

type Session struct {
	Id			uuid.UUID	`json:"id"`
	IsActive	bool		`json:"is_active"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

func (s *Session) Create(*User) error {

}

func (s *Session) Read(*User) error {

}

func (s *Session) Update(*User, func(*User)) error {

}

func (s *Session) Delete(*User) error {

}
