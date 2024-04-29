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

func ExportToFile() {
	logFile, err := os.Create("Output.txt")
	if err != nil {
		log.Fatalln("Failed to create")
	}
	defer logFile.Close()
}
