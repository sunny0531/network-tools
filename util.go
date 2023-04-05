package main

import (
	"context"
	"encoding/json"
	"fmt"
	probing "github.com/TheRushingWookie/pro-bing"
	"net"
	"os"
	"strings"
	"time"
)

func readConfig(file string) Config {
	body, ferr := os.ReadFile(file)
	if ferr != nil {
		fmt.Printf("unable to read file: %v\n", ferr)
	}
	var c Config
	jerr := json.Unmarshal(body, &c)
	if jerr != nil {
		fmt.Printf("unable to parse json: %v\n", jerr)
	}
	return c
}

func lookup(website host, dnsServer host) (host, error) {
	//https://www.reddit.com/r/golang/comments/esus2j/comment/ffca0jc/?utm_source=share&utm_medium=web2x&context=3
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, "udp", dnsServer.Host+":53")
		},
	}
	ip, err := r.LookupHost(context.Background(), website.Host)
	if err != nil {
		err_ := strings.SplitAfter(err.Error(), ":")
		if strings.TrimSpace(err_[len(err_)-1]) == "no such host" {
			return host{}, &HostNotFoundError{
				Err:       "DNS lookup failed",
				Website:   website,
				DnsServer: dnsServer,
			}
		} else {
			panic(err_)
		}

	} else {
		return host{
			Name: website.Name,
			Host: ip[0],
		}, nil
	}

}
func ping(website host, timeout int) bool {
	pinger, err := probing.NewPinger(website.Host)
	if err != nil {
		return false
	}
	pinger.Count = 3
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()
	err = pinger.RunWithCtx(ctx) // Blocks until finished.
	if err != nil || ctx.Err() != nil {
		return false
	}
	return true
}
