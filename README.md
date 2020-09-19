# go-rightmove

Implements a very basic scraper to read property details from RightMove.

This package extracts and unmarshals the JSON object that can be found at the bottom of a RightMove search page.

## Usage

1. Go to rightmove.co.uk and search for whatever type of properties you're interested in
1. Add your required filters to the search
1. Execute the search and copy the URL of the first results page (e.g. https://www.rightmove.co.uk/property-for-sale/find.html?searchType=SALE&locationIdentifier=REGION%5E87490&maxPrice=200000)
1. Call the `GetPropertiesFromSearchURL` function with the search URL
    ```golang
      ctx := context.Background()
      searchURL := "https://www.rightmove.co.uk/property-for-sale/find.html?searchType=SALE&locationIdentifier=REGION%5E87490&maxPrice=200000"
      properties, err := rightmove.GetPropertiesFromSearchURL(ctx, searchURL)
      if err != nil {
        log.Fatalln(err)
      }

      fmt.Printf("Retrieved %d properties from RightMove\n", len(properties))
      // Output: Retrieved 817 properties from RightMove
    ```
1. Process the resulting slice as necessary

## Legal

This implementation is a proof of concept only. The RightMove 'terms of use' states that the use of "scraping technology" is unauthorized. Do not use this library.