package config

import (
	"log"

	"github.com/joho/godotenv"
)

func GetApplicationConfig() AzureConfig {
	configMap, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}

	return AzureConfig{
		TenantId:     configMap["TENANT_ID"],
		ClientId:     configMap["CLIENT_ID"],
		ClientSecret: configMap["CLIENT_SECRET"],
	}
}

type AzureConfig struct {
	TenantId     string
	ClientId     string
	ClientSecret string
}
