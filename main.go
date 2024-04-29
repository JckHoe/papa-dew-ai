package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Application has started")

	if len(os.Args) < 2 {
		log.Fatalln("Invalid usage of application")
	}
	log.Printf("You have entered:  %s \n", os.Args[1])

	log.Println("Application has terminated")
}
