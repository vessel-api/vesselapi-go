// Package main demonstrates basic usage of the vesselapi Go SDK.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	vesselapi "github.com/vessel-api/vesselapi-go"
)

func main() {
	apiKey := os.Getenv("VESSELAPI_API_KEY")
	if apiKey == "" {
		log.Fatal("VESSELAPI_API_KEY environment variable is required")
	}

	client, err := vesselapi.NewVesselClient(apiKey)
	if err != nil {
		log.Fatalf("create client: %v", err)
	}
	ctx := context.Background()

	// Search for vessels by name.
	fmt.Println("--- Search Vessels ---")
	searchResult, err := client.Search.Vessels(ctx, &vesselapi.GetSearchVesselsParams{
		FilterName: vesselapi.Ptr("Ever Given"),
	})
	if err != nil {
		var apiErr *vesselapi.APIError
		if errors.As(err, &apiErr) {
			if apiErr.IsRateLimited() {
				log.Fatal("rate limited — try again later")
			}
			log.Fatalf("API error: %s (status %d)", apiErr.Message, apiErr.StatusCode)
		}
		log.Fatalf("search vessels: %v", err)
	}
	for _, v := range vesselapi.Deref(searchResult.Vessels) {
		fmt.Printf("Vessel: %s (IMO: %d)\n",
			vesselapi.Deref(v.Name),
			vesselapi.Deref(v.Imo),
		)
	}

	// Search for vessels by flag (e.g. Panama-flagged container ships).
	fmt.Println("\n--- Search Vessels by Flag ---")
	flagResult, err := client.Search.Vessels(ctx, &vesselapi.GetSearchVesselsParams{
		FilterFlag:       vesselapi.Ptr("PA"),
		FilterVesselType: vesselapi.Ptr("Container Ship"),
		PaginationLimit:  vesselapi.Ptr(5),
	})
	if err != nil {
		log.Fatalf("search vessels by flag: %v", err)
	}
	for _, v := range vesselapi.Deref(flagResult.Vessels) {
		fmt.Printf("Vessel: %s (IMO: %d, Country: %s)\n",
			vesselapi.Deref(v.Name),
			vesselapi.Deref(v.Imo),
			vesselapi.Deref(v.Country),
		)
	}

	// Search for ports by country.
	fmt.Println("\n--- Search Ports by Country ---")
	portSearch, err := client.Search.Ports(ctx, &vesselapi.GetSearchPortsParams{
		FilterCountry:   vesselapi.Ptr("NL"),
		PaginationLimit: vesselapi.Ptr(5),
	})
	if err != nil {
		log.Fatalf("search ports by country: %v", err)
	}
	for _, p := range vesselapi.Deref(portSearch.Ports) {
		fmt.Printf("Port: %s (%s)\n",
			vesselapi.Deref(p.Name),
			vesselapi.Deref(p.UnloCode),
		)
	}

	// Get a port by UNLOCODE.
	fmt.Println("\n--- Get Port ---")
	port, err := client.Ports.Get(ctx, "NLRTM")
	if err != nil {
		log.Fatalf("get port: %v", err)
	}
	if port.Port != nil {
		fmt.Printf("Port: %s (%s)\n",
			vesselapi.Deref(port.Port.Name),
			vesselapi.Deref(port.Port.UnloCode),
		)
	}

	// Handle a not-found port gracefully.
	fmt.Println("\n--- Not Found Handling ---")
	_, err = client.Ports.Get(ctx, "ZZZZZ")
	if err != nil {
		var apiErr *vesselapi.APIError
		if errors.As(err, &apiErr) && apiErr.IsNotFound() {
			fmt.Printf("Port ZZZZZ not found (status %d)\n", apiErr.StatusCode)
		} else {
			log.Fatalf("get port: %v", err)
		}
	}

	// Auto-paginate through all port events.
	fmt.Println("\n--- Port Events (paginated) ---")
	it := client.PortEvents.ListAll(ctx, &vesselapi.GetPorteventsParams{
		PaginationLimit: vesselapi.Ptr(10),
	})
	count := 0
	for it.Next() {
		event := it.Value()
		fmt.Printf("Event: %s at %s\n",
			vesselapi.Deref(event.Event),
			vesselapi.Deref(event.Timestamp),
		)
		count++
		if count >= 25 {
			break // Limit output for demo purposes.
		}
	}
	if it.Err() != nil {
		log.Fatalf("port events iteration: %v", it.Err())
	}
	fmt.Printf("Total events shown: %d\n", count)
}
