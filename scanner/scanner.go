package scanner

import (
	"net"
	"strconv"
	"time"

	"github.com/tommyforlini/go-portscan/model"
)

// ScanPortSync sync specific
func ScanPortSync(protocol, hostname string, port int) model.State {
	return scanPort(protocol, hostname, port)
}

// ScanPortAsync async specific
func ScanPortAsync(protocol, hostname string, port int) model.State {
	return scanPort(protocol, hostname, port)
}

// Scan function
func scanPort(protocol, hostname string, port int) model.State {
	result := model.State{Port: strconv.Itoa(port), Protocol: protocol, Opened: false}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return result
	}
	defer conn.Close()
	result.Opened = true
	return result
}
