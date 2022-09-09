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

func GetConfig() Configuration {
	config := Configuration{}
	file, err := os.Open("env.json")
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
