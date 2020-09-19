// Package rightmove implements a very basic scraper to read property details from RightMove.
//
// This package extracts and unmarshals the JSON object that
// can be found at the bottom of a RightMove search page.
//
// NOTE: This implementation is a proof of concept only. The RightMove 'terms of use' states
// that the use of "scraping technology" is unauthorized. Do not use this library.
package rightmove

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

// GetPropertiesFromSearchURL returns a slice of properties from the search URL provided
func GetPropertiesFromSearchURL(ctx context.Context, searchURL string) ([]Property, error) {
	properties := []Property{}

	// loop over pages of the search
	for {
		pageURL := fmt.Sprintf("%s&index=%d", searchURL, len(properties))
		pageProperties, total, err := getPropertiesFromURL(ctx, pageURL)
		if err != nil {
			log.Printf("Error getting data from page (index %d) : %v\n", len(properties), err)
			return nil, err
		}
		properties = append(properties, pageProperties...)
		if len(properties) >= total {
			break
		}
	}

	return properties, nil
}

// getPropertiesFromURL reads and returns a slice of properties from the provided URL
func getPropertiesFromURL(ctx context.Context, url string) ([]Property, int, error) {
	content, err := getPageContent(ctx, url)
	if err != nil {
		return nil, 0, err
	}
	jsonModel, err := extractJSONModel(string(content))
	if err != nil {
		log.Printf("Error extracting JSON model from page: %v\n", err)
		return nil, 0, err
	}
	totalProperties, err := strconv.Atoi(jsonModel.ResultCount)
	if err != nil {
		log.Printf("Error converting result count to an int: %v\n", err)
		return nil, 0, err
	}

	return filterFeaturedProperties(jsonModel.Properties), totalProperties, nil
}

// getPageContent makes a GET request to the provided URL and returns it's body
func getPageContent(ctx context.Context, url string) ([]byte, error) {
	client := http.DefaultClient

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Error creating request: %v\n", err)
		return nil, err
	}

	rsp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making request to RightMove: %v\n", err)
		return nil, err
	}
	defer rsp.Body.Close()
	content, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		log.Printf("Error reading request body: %v\n", err)
		return nil, err
	}

	return content, nil
}

// extractJSONModel extracts the JSON model from the page content provided
func extractJSONModel(content string) (*jsonModel, error) {
	// the JSON string appears between 'window.jsonModel =' and '</script>'
	r := regexp.MustCompile("window\\.jsonModel = (.*)<\\/script>")

	results := r.FindStringSubmatch(content)
	if len(results) < 2 {
		return nil, errors.New("JSON object not found in provided page")
	}

	// unmarshal into JSONModel
	jsonModel := &jsonModel{}
	err := json.Unmarshal([]byte(results[1]), jsonModel)
	if err != nil {
		return nil, err
	}

	return jsonModel, nil
}

// filterFeaturedProperties removes properties marked as featured as this will be duplicates
func filterFeaturedProperties(properties []Property) []Property {
	filteredProperties := []Property{}
	for _, property := range properties {
		if !property.FeaturedProperty {
			filteredProperties = append(filteredProperties, property)
		}
	}
	return filteredProperties
}
