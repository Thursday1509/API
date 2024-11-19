/*
package main

import (

	"fmt"
	"io/ioutil"
	"log"
	"net/http"

)

	func main() {
		// Define the URL you want to call
		url := "https://api.hackertarget.com/dnslookup/?q=flipkart.com" // example endpoint

		// Send an HTTP GET request
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		// Check the response status
		if resp.StatusCode != http.StatusOK {
			log.Fatalf("Failed to get a valid response: %v", resp.Status)
		}

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read response body: %v", err)
		}

		// Print the response
		fmt.Println("Response Body:", string(body))
	}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define a struct to represent a todo item
type Todo struct {
	TXT       string `json:"TXT"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://api.hackertarget.com/dnslookup/?q=flipkart.com" // API returning a list of todos

	// Send an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check if we got a valid response
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to get a valid response: %v", resp.Status)
	}

	// Decode the JSON array into a slice of Todo structs
	var todos []Todo
	if err := json.NewDecoder(resp.Body).Decode(&todos); err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	// Filter and print only the completed todos
	fmt.Println("Completed Todos:")
	for _, todo := range todos {
		if todo.Completed {
			fmt.Printf("- %s\n", todo.TXT)
		}
	}
}
*/

package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()

	// Define what to do when a visited HTML element is found
	c.OnHTML("body", func(e *colly.HTMLElement) {
		// Example: Searching for specific text in the element
		if e.DOM.Find("p:contains('TXT')").Length() > 0 {
			// Extract and print the text or any data you want
			fmt.Println("Found the specific text!")
			fmt.Println("Content:", e.Text)
		}
	})

	// Error handling for requests
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, err)
	})

	// Start the scraping process on the specified URL
	url := "https://api.hackertarget.com/dnslookup/?q=flipkart.com" // Replace with the target URL
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}
