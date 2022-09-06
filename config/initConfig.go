package config

import (
	"encoding/json"
	"log"
	"os"
)

var Config Configuration

type Configuration struct {
	ConnectionString string
	Database         string
	Collection       string
	ApiKey           string
}

func GetConfig(env string) Configuration {
	var fname string
	config := Configuration{}
	switch env {
	case "prod":
		fname = "prod-env.json"
		return config
	case "dev":
		fname = "dev-env.json"
	default:
		log.Fatal("unexpected environment name")
	}
	file, err := os.Open("config/" + fname)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Println(err)
	}
	return config

}
