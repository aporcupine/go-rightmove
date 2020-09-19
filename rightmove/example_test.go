package rightmove_test

import (
	"context"
	"fmt"
	"log"

	"github.com/aporcupine/go-rightmove/rightmove"
)

func ExampleGetPropertiesFromSearchURL() {
	ctx := context.Background()
	searchURL := "https://www.rightmove.co.uk/property-for-sale/find.html?searchType=SALE&locationIdentifier=REGION%5E87490&maxPrice=200000"
	properties, err := rightmove.GetPropertiesFromSearchURL(ctx, searchURL)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Retrieved %d properties from RightMove\n", len(properties))
	// Output: Retrieved 817 properties from RightMove
}
