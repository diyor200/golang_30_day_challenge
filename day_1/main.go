package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var logs = []string{
		"[INFO] Running checks...",
		"[INFO] Starting linting the current package at /Users/macbook/Desktop",
		"[ERROR] DS",
		"[WARN] gopls: failed to install gopls",
		"this is invalid log",
	}

	start := time.Now()
	result := ProcessLogs(logs)
	fmt.Println("Time elapsed:", time.Since(start))

	fmt.Println(result)
}

func ProcessLogs(logs []string) map[string]int {
	var (
		info = "INFO"
		warn = "WARN"
		err  = "ERROR"
	)

	var (
		summary = make(map[string]int, 3)
		jobs    = make(chan string)
		results = make(chan int, len(logs))
		workers = 2
	)

	for i := 0; i < workers; i++ {
		go worker(jobs, results)
	}

	for i := range logs {
		jobs <- logs[i]
	}
	close(jobs)

	for i := 0; i < len(logs); i++ {
		r := <-results
		switch r {
		case 1:
			summary[info]++
		case 2:
			summary[warn]++
		case 3:
			summary[err]++
		}
	}

	return summary
}

func worker(jobs <-chan string, results chan<- int) {
	var (
		info = "INFO"
		warn = "WARN"
		err  = "ERROR"
	)

	for j := range jobs {
		if ok := strings.Contains(j, info); ok {
			results <- 1
		} else if ok := strings.Contains(j, warn); ok {
			results <- 2
		} else if ok := strings.Contains(j, err); ok {
			results <- 3
		} else {
			results <- 0
		}
	}
}
