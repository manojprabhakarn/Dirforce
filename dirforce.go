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

	wordlist(Wordlist, RHOST)
	// checkpath(RHOST)

	// // scanning
	// for i := 0; i < 4; i++{
	// 	fmt.Printf("GeeksforGeeks\n")
	// 	}
}

// Retreving Wordlist
func wordlist(Wordlist, RHOST string) {
	fmt.Println("[*] Parshing Wordlist...")

	file, err := os.Open(Wordlist)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		url := RHOST + "/" + scanner.Text()
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			fmt.Println("[*] Error: An UnExpected Error Occured")
		}
		status_code := resp.StatusCode
		fmt.Println("[*]", url, ":", status_code)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// func checkpath(RHOST, path string) {
// 	url := RHOST + "/" + path
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 		fmt.Println("[*] Error: An UnExpected Error Occured")
// 	}
// 	if resp.StatusCode == 200{
// 		fmt.Println("[*] Valid Path Found:",path)
// 	}

// }
