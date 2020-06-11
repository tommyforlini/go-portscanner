package main

import (
	"log"
	"sync"

	"github.com/tommyforlini/go-portscan/model"
	"github.com/tommyforlini/go-portscan/scanner"
	"github.com/tommyforlini/go-portscan/utils"
)

// PORTS max ports to scan
const PORTS = 1024

// PROTOCOL what type to scan
const PROTOCOL = "tcp"

// HOST what to scan
const HOST = "localhost"

func main() {

	log.Printf("Scanning Host: %s | Protocol: %s\n\n\n", HOST, PROTOCOL)

	scanSync(HOST)

	scanASync(HOST)

	scanASyncNoChannel(HOST)
}

func scanASync(hostname string) {
	defer utils.Track(utils.Runningtime("execute scanASync"))

	var results []*model.State
	ch := make(chan *model.State)

	var wg sync.WaitGroup

	go func() {
		for i := 1; i <= PORTS; i++ {
			wg.Add(1)
			go func(port int) {
				defer wg.Done()
				ch <- scanner.ScanPortAsync(PROTOCOL, hostname, port)
			}(i)
		}
		wg.Wait()
		close(ch)
	}()

	for elem := range ch {
		if elem.Opened {
			results = append(results, elem)
		}
	}

	log.Printf("async slice size %v", len(results))
	log.Printf("async results %v\n", results)
}

func scanASyncNoChannel(hostname string) {
	defer utils.Track(utils.Runningtime("execute scanASyncNoChannel"))

	var results []*model.State

	var wg sync.WaitGroup

	for i := 1; i <= PORTS; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			result := scanner.ScanPortAsync(PROTOCOL, hostname, port)
			if result.Opened {
				results = append(results, result)
			}

		}(i)
	}
	wg.Wait()

	log.Printf("scanASyncNoChannel slice size %v", len(results))
	log.Printf("scanASyncNoChannel results %v\n", results)
}

func scanSync(hostname string) {
	defer utils.Track(utils.Runningtime("execute scanSync"))

	var results []model.State

	for i := 1; i <= PORTS; i++ {
		result := scanner.ScanPortSync(PROTOCOL, hostname, i)
		if result.Opened {
			results = append(results, result)
		}
	}

	log.Printf("sync slice size %v\n", len(results))
	log.Printf("sync results %v\n", results)
}
