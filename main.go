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

func main() {

	scanSync("localhost")

	scanASync("localhost")

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
				ch <- scanner.ScanPortAsync("tcp", hostname, port)
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

func scanSync(hostname string) {
	defer utils.Track(utils.Runningtime("execute scanSync"))

	var results []model.State

	for i := 1; i <= PORTS; i++ {
		result := scanner.ScanPortSync("tcp", hostname, i)
		if result.Opened {
			results = append(results, result)
		}
	}

	log.Printf("sync slice size %v\n", len(results))
	log.Printf("sync results %v\n", results)
}
