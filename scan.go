package main

import (
	"fmt"
	"log"
	"net"
	"sort"
	"time"
)

// Heartbeat dials a (network) "tcp" (addr) "localhost:9001". If the dial returns
// a response (err == nil), return true; else return false (default).
func Heartbeat(addr string) bool {
	// TODO: Improve ...
	_, err := net.DialTimeout("tcp", addr, time.Microsecond*400)
	if err == nil {
		log.Printf("Service running on port %s ...", addr)
		return true
	}
	return false
}

// ScanLocalhost ports
func ScanLocalhost() []string {
	var active []string
	maxPort := 65535

	// No point in checking < 22 really, but leaving
	// at 2 until absolutely necessary
	for i := 2; i < maxPort; i++ {
		// TODO: run Heartbeat concurretly, *and consistently*.
		localAddr := fmt.Sprintf("localhost:%d", i)
		if Heartbeat(localAddr) == true {
			active = append(active, localAddr)
		}
	}

	sort.Strings(active)

	return active
}
