/*
Copyright © 2024 Donald Gifford dgifford06@gmail.com
*/
package src

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Auth     ConfigAuth
	Nest     ConfigNest
	InfluxDB ConfigInfluxDB
}

type ConfigAuth struct {
	ClientId      string `yaml:"client_id"`
	ClientSecret  string `yaml:"client_secret"`
	RedirectUri   string `yaml:"redirect_uri"`
	TokenFileName string `yaml:"token_file_name"`
}

type ConfigNest struct {
	ProjectID string `yaml:"project_id"`
	DeviceID  string `yaml:"device_id"`
}

type ConfigInfluxDB struct {
	URL    string `yaml:"url"`
	Token  string `yaml:"token"`
	Org    string `yaml:"org"`
	Bucket string `yaml:"bucket"`
}

func NewConfig() {
	c1 := Config{
		Auth: ConfigAuth{
			ClientId:      "my-client-id",
			ClientSecret:  "my-client-secret",
			RedirectUri:   "http://localhost:8080",
			TokenFileName: "token.json",
		},
		Nest: ConfigNest{
			ProjectID: "my-project-id",
			DeviceID:  "my-device-id",
		},
		InfluxDB: ConfigInfluxDB{
			URL:    "http://localhost:8086",
			Token:  "my-token",
			Org:    "my-org",
			Bucket: "my-bucket",
		},
	}

	yamlData, err := yaml.Marshal(&c1)
	if err != nil {
		fmt.Println("Error while marshalling data")
		log.Fatal(err)
	}

	fileName := ".nester.yaml"
	err = os.WriteFile(fileName, yamlData, 0644)
	if err != nil {
		fmt.Println("Error while writing yaml file")
		log.Fatal(err)
	}
}
