package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	log.Println("Application has started")

	if len(os.Args) < 2 {
		log.Fatalln("Invalid usage of application")
	}
	log.Printf("You have entered:  %s \n", os.Args[1])

	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	var htmlContent []string

	for scanner.Scan() {
		var line = scanner.Text()
		if strings.Contains(line, "magnet") {
			htmlContent = append(htmlContent, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	ExportToFile(&htmlContent)

	log.Println("Application has terminated")
}

func ExportToFile(content *[]string) {
	logFile, err := os.OpenFile("Output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Failed to open or create file:", err)
	}
	defer logFile.Close()

	for _, line := range *content {
		_, err := logFile.WriteString(line + "\n")
		if err != nil {
			log.Println("Error writing to file:", err)
		}
	}
}
