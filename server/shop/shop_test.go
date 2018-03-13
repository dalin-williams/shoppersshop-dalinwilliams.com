package shop_test

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattes/migrate/database/postgres"
	"github.com/mattes/migrate"


	"github.com/funkeyfreak/vending-machine-api/etc"
	"github.com/funkeyfreak/vending-machine-api/server/shop/pg"
	"github.com/funkeyfreak/vending-machine-api/server/shop/services"
	"github.com/funkeyfreak/vending-machine-api/server/shop"
	"github.com/stretchr/testify/assert"
)

// a container for all of our test resources
type appTests struct {
	*shop.Shop
}

// Setups a test database named vendingmachine_test for testing our application
func setupTestDatabase() (*sqlx.DB, error) {
	//check for and re-create the test database
	db, err := sqlx.Connect("postgres", "postgres://localhost:5432/database?sslmode=disable")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`drop database if exists vendingmachine_test`)

	_, err = db.Exec(`create database vendingmachine_test`)
	if err != nil {
		return nil, err
	}
	db.Close()


	//run migrations on our test database
	url := "postgres://localhost/vendingmachine_test?sslmode=disable"
	db, err = sqlx.Connect("postgres", url)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`create extension citext`)
	if err != nil {
		return nil, err
	}

	m, err := migrate.New( "../db/migrate", url)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate test database: %v", err)
	}
	m.Steps(2)

	return db, nil
}

func TestApp(t *testing.T) {
	t.Run("pg", pgTests(t).run)
	t.Run("services", serviceTests(t).run)
}

func serviceTests(t *testing.T) *appTests {
	testObj := appTests{}

	conn, err := setupTestDatabase()
	if err != nil {
		t.Fatalf("Could not setup postgres database: %v", err)
		return &testObj
	}

	newShop, err  = shop.New(pg.Driver(conn))
	if err != nil {
		t.Errorf("could not create pg store: %v", err)
	}

	return &appTests{
		Shop: newShop
	}
}

func (s *appTests) run(t *testing.T) {
	if !assert.NotNil(t, s.Shop) {
		assert.FailNow(t, "The returned shop is nil")
	}

	t.Run("inventory", s.testInventory)
	t.Run("pay", s.testPay)
	t.Run("session", s.testSession)
	t.Run("category", s.testCategory)
	t.Run("vend", s.testVend)
}