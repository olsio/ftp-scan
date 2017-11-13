package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
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

func worker(worker int, tasksCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasksCh
		if !ok {
			return
		}
		d := time.Duration(task) * time.Millisecond
		time.Sleep(d)
		fmt.Println("processing task", worker)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	tasksCh := make(chan int)

	for i := 0; i < workers; i++ {
		go worker(i, tasksCh, wg)
	}

	for i := 0; i < tasks; i++ {
		tasksCh <- i
	}

	close(tasksCh)
}

var world = []byte("world")

func main() {
	/*
		var wg sync.WaitGroup
		wg.Add(36)
		go pool(&wg, 36, 50)
		wg.Wait()
	*/

	file, err := os.Open("./slim.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	f, err := os.OpenFile("./slim2.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		if i%10000 == 0 {
			fmt.Println("Inserted:", i)
		}

		var grab Grab

		jsonBlob := scanner.Bytes()
		err := json.Unmarshal(jsonBlob, &grab)
		if err != nil {
			fmt.Println("error:", err)
		}
		if _, err = f.WriteString(grab.IP + "\n"); err != nil {
			panic(err)
		}
		i = i + 1
	}
}
