package main

type Config struct {
	VendingProviders struct {
		AmazonAccessKey string `json:"amazon_access_key"`
		AmazonSecretKey string `json:"amazon_secret_key"`
	} `json:"vending_providers"`
}
