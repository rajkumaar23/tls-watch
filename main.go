package main

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"
)

type Config struct {
	Hostnames        []string `json:"hostnames"`
	TelegramID       int      `json:"telegram_id"`
	TelegramBotToken string   `json:"telegram_bot_token"`
	Threshold        int      `json:"threshold"`
}

var config Config
var WaitGroup sync.WaitGroup
var HttpClient = &http.Client{}

func main() {
	configFileContents, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(configFileContents, &config); err != nil {
		panic(err)
	}

	for _, hostname := range config.Hostnames {
		WaitGroup.Add(1)
		go ProcessHostname(hostname)
	}

	WaitGroup.Wait()
}
