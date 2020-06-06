package main

import (
	"sync"

	"github.com/tommyforlini/go-portscan/model"
	"github.com/tommyforlini/go-portscan/scanner"
	"github.com/tommyforlini/go-portscan/utils"
)

func main() {

	scanSync("localhost")

	scanASync("localhost")

}

func scanASync(hostname string) []model.State {
	defer utils.Track(utils.Runningtime("execute scanASync"))

	var results []model.State

	wg := &sync.WaitGroup{}
	for i := 0; i <= 1024; i++ {
		go func(i int) {
			wg.Add(1)
			scanner.ScanPortAsync("tcp", hostname, i)
			// results = append(results, scanner.ScanPortAsync("tcp", hostname, i))
			wg.Done()
		}(i)
	}
	wg.Wait()

	return results
}

func scanSync(hostname string) []model.State {
	defer utils.Track(utils.Runningtime("execute scanSync"))

	var results []model.State

	for i := 0; i <= 1024; i++ {
		scanner.ScanPortSync("tcp", hostname, i)
		// results = append(results, scanner.ScanPortSync("tcp", hostname, i))
	}

	return results
}
