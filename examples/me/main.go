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

	me, err := client.Me.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s\n", me.ID)
	fmt.Printf("Name: %s\n", me.Name)
	fmt.Printf("Email: %s\n", me.Email)
	fmt.Printf("Balance: Rp %.0f\n", me.Balance)
}