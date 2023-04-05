package main

import "fmt"

type HostNotFoundError struct {
	Err       string
	Website   host
	DnsServer host
}

func (e *HostNotFoundError) Error() string {
	return fmt.Sprintf("Host %v not found, DNS server: %v", e.Website.Host, e.DnsServer.Host)
}
