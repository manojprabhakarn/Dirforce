package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	var baseURL string
	var wordlist string

	flag.StringVar(&baseURL, "rhost", "default", "Enter the target URL")
	flag.StringVar(&wordlist, "wordlist", "", "the file to include")
	// Parse the command-line arguments
	flag.Parse()

	fmt.Println(baseURL, wordlist)

	// Read the contents of the file-------------------------------------
	contents, err := ioutil.ReadFile(wordlist)
	if err != nil {
		// There was an error reading the file
		fmt.Println(err)
		return
	}
	// Set the list of directories
	directories := []string{}

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(bytes.NewReader(contents))
	for scanner.Scan() {
		// Append the line to the slice
		directories = append(directories, scanner.Text())
	}


	// Try each directory
	for _, dir := range directories {
		// Construct the full URL for the directory
		url := baseURL +"/"+ dir

		// Make an HTTP GET request to the URL
		resp, err := http.Get(url)
		if err != nil {
			// There was an error making the request
			fmt.Println(err)
			continue
		}

		// Check the response status code
		if resp.StatusCode == 200 {
			// The directory was found
			fmt.Println(url + " was found")
		} else {
			// The directory was not foundS
			fmt.Println(url + " was not found")
		}
	}
}
