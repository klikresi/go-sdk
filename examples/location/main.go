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

	page, err := client.Location.Search(ctx, "depok", klikresi.WithLimit(10))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("First page of results for \"depok\" (next cursor: %s):\n", page.NextCursor)
	for _, l := range page.Data {
		fmt.Printf("- %s: %s\n", l.ID, l.Name)
	}

	all, err := client.Location.AllLocations(ctx, "depok")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%d total matches.\n", len(all))

	provinces, err := client.Location.AllProvinces(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d provinces loaded.\n", len(provinces))
}
