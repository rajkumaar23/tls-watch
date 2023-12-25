package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/joho/godotenv"
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
	loadConfigFromEnv()

	for _, hostname := range config.Hostnames {
		WaitGroup.Add(1)
		go ProcessHostname(hostname)
	}

	WaitGroup.Wait()
}

func loadConfigFromEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("error loading .env file: %s\nproceeding to see if env variables are set already", err)
	}

	hostnames := strings.TrimSpace(os.Getenv("HOSTNAMES"))
	if hostnames == "" {
		log.Panicf("hostnames are empty")
	}

	config.Hostnames = strings.Split(hostnames, ",")
	if len(config.Hostnames) < 1 {
		log.Fatalf("failed parsing hostnames: %+v\n", hostnames)
	}

	threshold, err := strconv.Atoi(os.Getenv("THRESHOLD"))
	if err != nil {
		log.Printf("error parsing threshold value: %s\nfalling back to default value: 10\n", err)
		config.Threshold = 10
	} else {
		config.Threshold = threshold
	}

	telegramID, err := strconv.Atoi(os.Getenv("TELEGRAM_ID"))
	if err != nil {
		log.Fatalf("parsing telegramID failed : %s\n", err)
	} else {
		config.TelegramID = telegramID
	}

	config.TelegramBotToken = strings.TrimSpace(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if config.TelegramBotToken == "" {
		log.Fatal("telegram bot token cannot be empty")
	}
}
