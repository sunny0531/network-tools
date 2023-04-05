package main

import (
	"fmt"
)

func main() {
	timeout := 10
	c := readConfig("example.json")
	for _, dh := range c.DnsServers {
		fmt.Printf("%v, %v\n", dh.Name, dh.Host)
		if ping(dh, timeout) {
			fmt.Println("Ping successful\n")
		} else {
			fmt.Println("Ping failed\n")
		}
		for _, wh := range c.Website {
			fmt.Printf("%v\n Adresses:%v\n ", wh.Name, wh.Host)
			result, err := lookup(wh, dh)
			if err != nil {
				fmt.Println("DNS lookup failed")
			} else {
				fmt.Printf("ip: %v\n", result.Host)
				if ping(result, timeout) {
					fmt.Println("Ping successful\n")
				} else {
					fmt.Println("Ping failed\n")
				}
			}

		}
		fmt.Println("-----------------------------")
	}

}
