package main

import (
	"log"
	"os"
	"time"

	"github.com/awoldes/goanda"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func candles() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &goanda.ConnectionConfig{
		UserAgent: "goanda",
		Timeout:   10 * time.Second,
		Live:      false,
	}

	granularity := goanda.GranularityFiveSeconds

	key := os.Getenv("OANDA_API_KEY")
	accountID := os.Getenv("OANDA_ACCOUNT_ID")

	oanda, err := goanda.NewConnection(accountID, key, config)
	if err != nil {
		log.Fatalf("Error creating connection: %v", err)
	}

	// Example 1: Get latest 10 candles
	history, err := oanda.GetCandles("EUR_USD", 10, granularity)
	if err != nil {
		log.Fatalf("Error getting candles: %v", err)
	}
	spew.Dump("Latest 10 candles:", history)

	// Example 2: Get candles between specific time range
	from := time.Now().Add(-24 * time.Hour) // 24 hours ago
	to := time.Now().Add(-12 * time.Hour)   // 12 hours ago

	rangeHistory, err := oanda.GetTimeRangeCandles("EUR_USD", goanda.GranularityHour, from, to)
	if err != nil {
		log.Fatalf("Error getting time range candles: %v", err)
	}
	spew.Dump("Time range candles (24h-12h ago):", rangeHistory)
}
