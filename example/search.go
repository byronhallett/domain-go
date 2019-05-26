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

var locations = []types.SearchLocation{
	types.SearchLocation{
		Area:   "Sydney City",
		Region: "Sydney Region",
	},
	types.SearchLocation{
		Area:   "Eastern Suburbs",
		Region: "Sydney Region",
	},
	types.SearchLocation{
		Area:   "Inner West",
		Region: "Sydney Region",
	},
}
var params = types.SearchParameters{
	ListingType: "Rent",
	MinBedrooms: 2,
	MinPrice:    340,
	MaxPrice:    450,
	Locations:   locations,
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
	results, err := session.ResidentialSearch(params)
	if err != nil {
		log.Fatalln(err)
	}

	for _, r := range *results {
		detail := &r.Listing.PropertyDetails
		fmt.Printf("%s %s, %s, %s\n",
			detail.StreetNumber, detail.Street, detail.Suburb, detail.State)
	}
}
