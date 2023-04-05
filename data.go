package main

type host struct {
	Name string `json:"name"`
	Host string `json:"host"`
}
type Config struct {
	DnsServers []host `json:"dnsServer"`
	Website    []host `json:"websites"`
}
