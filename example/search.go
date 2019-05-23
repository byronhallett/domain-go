package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	domain "github.com/byronhallett/domain-go"
	types "github.com/byronhallett/domain-go/types"
)

var session *domain.Session

func search() {
	locations := []types.SearchLocation{
		types.SearchLocation{
			State:  "NSW",
			Region: "Sydney Region",
		},
	}
	params := types.SearchParameters{
		ListingType: "Rent",
		MinBedrooms: 2,
		MinPrice:    300,
		MaxPrice:    420,
		Locations:   &locations,
	}
	results, err := session.ResidentialSearch(params)
	if err != nil {
		log.Fatalln(err)
	}
	for _, r := range *results {
		detail := &r.Listing.PropertyDetails
		fmt.Printf("%s, %s, %s, %s\n",
			detail.UnitNumber, detail.StreetNumber, detail.Street, detail.Suburb)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	session, err = domain.NewSession(
		os.Getenv("CLIENT"), os.Getenv("SECRET"),
	)
	if err != nil {
		log.Fatalln(err)
	}
	// Search and print results until last page
	search()
}
