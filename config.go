package main

var config Config

type Config struct {
	VendingProviders struct {
		AmazonAccessKey string `json:"amazon_access_key"`
		AmazonSecretKey string `json:"amazon_secret_key"`
	} `json:"vending_providers"`
	Database struct {
		DatabaseName string `json:"database_name"`
		Username	string `json:"username"`
		Password	string `json:"password"`
		Sslmode		string `json:"sslmode"`
	} `json:"database"`
}
