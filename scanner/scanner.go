package scanner

import (
	"net"
	"strconv"
	"time"

	"github.com/tommyforlini/go-portscan/model"
)

// ScanPortSync sync specific
func ScanPortSync(protocol, hostname string, port int) model.State {
	result := model.State{Port: strconv.Itoa(port), Protocol: protocol, Opened: false}
	err := scan(protocol, hostname, port)
	if err != nil {
		return result
	}

	result.Opened = true
	return result
}

// ScanPortAsync async specific
func ScanPortAsync(protocol, hostname string, port int) *model.State {
	result := &model.State{Port: strconv.Itoa(port), Protocol: protocol, Opened: false}
	err := scan(protocol, hostname, port)
	if err != nil {
		return result
	}

	result.Opened = true
	return result
}

// scan scan for the port on host with protocol
// error implies that is closed and unused
func scan(protocol, hostname string, port int) error {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 100*time.Millisecond)

	if err != nil {
		return err
	}
	conn.Close()
	return nil
}
