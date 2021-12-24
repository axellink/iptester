package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Key string `json:"key"`
}

func GetConf() (*Config, error) {
	file, err := os.Open("key.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
