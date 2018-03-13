package pg

import (
	"github.com/jmoiron/sqlx"

	"github.com/funkeyfreak/vending-machine-api/server/shop"
	"fmt"
	"github.com/RumbleMonkey/rumbled/store"
)

// A wrapper for our database instance
type driver struct {
	Db      string `json:"database_name"`
	User    string `json:"username"`
	Pass    string `json:"password"`
	SslMode string `json:"sslmode"`

	conn *sqlx.DB
}

func Driver(db *sqlx.DB) shop.Driver {
	return &driver{conn: db}
}

// Connect to the database given the connection string
func (d * driver) dial() error {
	if d.conn != nil {
		return nil
	}

	creds := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=%s", d.Db, d.User, d.Pass, d.SslMode)
	conn, err := sqlx.Connect("postgres", creds)
	if err != nil {
		return fmt.Errorf("pg store dial error: %v", err)
	}

	d.conn = conn
	return nil
}

func (d *driver) Pay(s *shop.Shop) error {
	if err := d.dial(); err != nil {
		return err
	}
	s.Pay = &pay{d.conn}
	return nil
}

func (d *driver) Inventory(s *shop.Shop) error {
	if err := d.dial(); err != nil {
		return err
	}
	s.Inventory = &pay{d.conn}
	return nil
}

func (d *driver) Vend(s *shop.Shop) error {
	if err := d.dial(); err != nil  {
		return err
	}
	s.Inventory = &vend{d.conn}
	return nil
}

func (d *driver) Session(s *shop.Shop) error {
	if err := d.dial(); err != nil {
		return err
	}
	s.Session = &session{d.conn}
	return nil
}

func (d *driver) Category(s *shop.Shop) error {
	if err := d.dial(); err != nil {
		return err
	}
	s.Category = &category{d.conn}
	return nil
}