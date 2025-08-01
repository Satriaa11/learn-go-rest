package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// FactResponse Struct with json struct tags
type FactResponse struct {
	Text   string `json:"fact"` // Ubah dari "text" ke "fact"
	Length int    `json:"length"`
}

func main() {
	// create request with empty method and URL
	req, err := http.NewRequest("GET", "https://catfact.ninja/fact", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// create a client and send the request
	client := http.Client{}

	// call request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Tambahkan pengecekan status code
	fmt.Printf("Status Code: %d\n", res.StatusCode)
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status: %s\n", res.Status)
	}

	// close the response body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
			os.Exit(0)
		}
	}(res.Body)

	// read the response body
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Debug: print raw response body
	// fmt.Printf("Raw response body: %s\n", string(resBody))
	// fmt.Printf("Response body length: %d\n", len(resBody))

	// convert the response body to a string and print it
	var factResponse FactResponse
	err = json.Unmarshal(resBody, &factResponse)
	if err != nil {
		fmt.Println("JSON unmarshal error:", err)
		os.Exit(0)
	}

	// fmt.Printf("Full struct: %+v\n", factResponse)
	fmt.Println("fact:", factResponse.Text)
	fmt.Println("length:", factResponse.Length)

}
