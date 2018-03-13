package services

import "github.com/funkeyfreak/vending-machine-api/server/shop"

type inventory struct {
	services map[string]shop.Service
}

