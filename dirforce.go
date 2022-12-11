package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	RHOST := os.Args[1]
	// Wordlist := os.Args[2]

	fmt.Println("[*] Checking RHOST....")

	// Evaluate the URL
	resp, err := http.Get(RHOST)
	if err != nil {
		fmt.Println("[FAIL]")
		fmt.Println("[!] Error: Cannot Reach RHOST:", RHOST)
		log.Fatal(err)
	}

	fmt.Println("[DONE]")
	fmt.Println("[*] RHOST is Reachable")

	status_code := resp.StatusCode
	fmt.Println("[*]", RHOST, ":", status_code)


}

// Retreving Wordlist 
func Wordlist() {
	fmt.Println("[*] Parshing Wordlist...")

	file, err := os.Open("wordlist/all.txt")
	if err != nil {
		panic(err)
	}

	// Word Counter
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)

	// initiate counter
	count := 0
	// for looping through results
	for fileScanner.Scan() {
		count++
	}
	// print total word count
	fmt.Printf("[*] Total Paths to Check: '%d'", count)

	// check if there was an error while reading words from file
	if err := fileScanner.Err(); err != nil {
		panic(err)
	}

}
