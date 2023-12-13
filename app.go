package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Replace "YOUR_API_KEY" and "NEWS_API_ENDPOINT" with your actual API key and endpoint
	apiKey := "YOUR_API_KEY"
	apiEndpoint := "https://newsapi.org/v2/top-headlines"

	// Example request parameters
	country := "us"
	category := "general"

	// Build the request URL with parameters
	url := fmt.Sprintf("%s?country=%s&category=%s&apiKey=%s", apiEndpoint, country, category, apiKey)

	// Make the HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	// Parse the JSON response
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print the results
	fmt.Println("Status:", result["status"])
	fmt.Println("Total Results:", result["totalResults"])

	// Print the articles
	articles, ok := result["articles"].([]interface{})
	if !ok {
		fmt.Println("Error parsing articles")
		return
	}

	for i, article := range articles {
		articleMap, ok := article.(map[string]interface{})
		if !ok {
			fmt.Println("Error parsing article", i)
			continue
		}
		fmt.Printf("\nArticle %d:\n", i+1)
		fmt.Println("Title:", articleMap["title"])
		fmt.Println("Author:", articleMap["author"])
		fmt.Println("Description:", articleMap["description"])
		fmt.Println("URL:", articleMap["url"])
	}

	// Handle other information as needed
}
