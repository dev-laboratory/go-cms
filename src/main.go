package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response struct {
	Headers []string
	Content string
}

// This is the entry point function for the application
// It is called (from index.php) with command line arguments representing
// the GET and POST data that is received from the client by the web server
func main() {
	// Get the command line arguments
	args := os.Args[1:]

	// Extract the GET map
	getArg := ""
	if len(args) > 0 {
		getArg = args[0]
	}

	// Extract the POST map
	postArg := ""
	if len(args) > 1 {
		postArg = args[1]
	}

	headers, content := processInput(getArg, postArg)

	// Construct the response containing headers and content
	resA := &response{
		Headers: headers,
		Content: content,
	}
	// Encode the response in text format
	resB, _ := json.Marshal(resA)

	// Output the response
	fmt.Print(string(resB))
}

func processInput(getArg string, postArg string) ([]string, string) {
	return headers(), content(getArg, postArg)
}

func headers() []string {
	return []string{"apple", "peach", "pear"}
}

func content(getArg string, postArg string) string {
	return fmt.Sprintf("The content. Get: %s Post: %s", getArg, postArg)
}
