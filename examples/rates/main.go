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

	// By district ID, filtered to JNE only.
	byID, err := client.Rates.CalculateByID(ctx, "33.08.20", "32.09.31", 1, klikresi.CourierJNE)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Rates from %s to %s:\n\n", byID.Origin.Name, byID.Destination.Name)
	for _, p := range byID.Pricing {
		fmt.Printf("- [%s] %s (%s): Rp %d, %s\n", p.CourierName, p.Service, p.Type, int(p.Price), p.Duration)
	}

	// By location name.
	byName, err := client.Rates.CalculateByName(ctx, "Secang, Kabupaten Magelang, Jawa Tengah", "Depok, Kabupaten Cirebon, Jawa Barat", 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%d offers found by name.\n", len(byName.Pricing))
}
