package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// This URL gets the current active roster for the Seattle Mariners (team 136).
	url := "https://statsapi.mlb.com/api/v1/teams/136/roster?rosterType=active"

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request failed: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Check that the server said "OK"
	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "bad status: %s\n", resp.Status)
		os.Exit(1)
	}

	// Read the entire response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read body: %v\n", err)
		os.Exit(1)
	}

	// Print the raw JSON to the terminal
	fmt.Println(string(body))
}
