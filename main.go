package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Grab struct {
	IP             string    `json:"ip"`
	Domain         string    `json:"domain,omitempty"`
	Time           string    `json:"timestamp"`
	Data           *GrabData `json:"data,omitempty"`
	Error          *string   `json:"error,omitempty"`
	ErrorComponent string    `json:"error_component,omitempty"`
}

type GrabData struct {
	Banner string `json:"banner,omitempty"`
}

func main() {
	file, err := os.Open("./example.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	store := NewStore("my.db")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		var jsontype Grab
		json.Unmarshal([]byte(line), &jsontype)
		fmt.Printf("Results: %v\n", jsontype)

		if jsontype.Error == nil {
			store.AddScanTarget([]byte(line))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
