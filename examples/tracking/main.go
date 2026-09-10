package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/klikresi/go-sdk"
)

func main() {
	apiKey := os.Getenv("KLIKRESI_API_KEY")
	if apiKey == "" {
		log.Fatal("KLIKRESI_API_KEY is required")
	}

	client := klikresi.NewClient(apiKey)
	ctx := context.Background()

	tracking, err := client.Tracking.Get(ctx, "YOUR-AWB", klikresi.CourierJNE)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status: %s\n", tracking.Status)
	fmt.Printf("Origin: %s - %s\n", tracking.Origin.ContactName, tracking.Origin.Address)
	fmt.Printf("Destination: %s - %s\n\n", tracking.Destination.ContactName, tracking.Destination.Address)

	for _, h := range tracking.Histories {
		fmt.Printf("[%s] %s: %s\n", h.Date.Format("2006-01-02 15:04"), h.Status, h.Message)
	}
}
