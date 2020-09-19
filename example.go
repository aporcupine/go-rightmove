package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/aporcupine/go-rightmove/rightmove"
)

func main() {
	rightMoveURL := flag.String("rightmove-search-url", "", "the url of a RightMove search")
	flag.Parse()

	if *rightMoveURL == "" {
		log.Fatalln("You must provide a RightMove URL")
	}

	ctx := context.Background()
	properties, err := rightmove.GetPropertiesFromSearchURL(ctx, *rightMoveURL)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Retrieved %d properties from RightMove\n", len(properties))
}
