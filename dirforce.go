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

	Wordlist()
}

// wordlist
func Wordlist() {
	fmt.Println("[*] Parshing Wordlist...")

	file, err := os.Open("common.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Scanner := bufio.NewScanner(file)
	// Scanner.Split(bufio.ScanWords)

	// for Scanner.Scan() {
	// 	fmt.Println(Scanner.Text())
	// }
	// if err := Scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

}
