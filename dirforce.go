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
	Wordlist := os.Args[2]

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

	wordlist(Wordlist)
	checkpath(RHOST)

	// scanning
	fmt.Println("\n [*] Beginning scan... \n")
	for i in 
}
}

// Retreving Wordlist
func wordlist(Wordlist string) {
	fmt.Println("[*] Parshing Wordlist...")

	file, err := os.Open(Wordlist)
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
	// fmt.Printf("[*] Total Paths to Check: '%d'", count)

	// check if there was an error while reading words from file
	if err := fileScanner.Err(); err != nil {
		panic(err)
	}
	return count
}

func checkpath(RHOST, path string) {
	url := RHOST + "/" + path
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		fmt.Println("[*] Error: An UnExpected Error Occured")
	}
	if resp.StatusCode == 200{
		fmt.Println("[*] Valid Path Found:",path)
	}

}

