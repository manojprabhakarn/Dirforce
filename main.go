package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	RHOST := os.Args[1]
	Wordlist := os.Args[2]

	fmt.Println(RHOST)
	fmt.Println(Wordlist)

	fmt.Println("[*] Checking RHOST....")

	// Establish Connection
	connection, err := net.Dial("tcp", RHOST+":"+"8989")

	if err != nil {
		panic(err)
	}
	        ///send some data
			_, err = connection.Write([]byte("Hello Server! Greetings."))
			buffer := make([]byte, 1024)
			mLen, err := connection.Read(buffer)
			if err != nil {
					fmt.Println("Error reading:", err.Error())
			}
			fmt.Println("Received: ", string(buffer[:mLen]))
			defer connection.Close()
}
