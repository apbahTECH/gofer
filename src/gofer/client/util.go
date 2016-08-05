package main

import (
    "fmt"
    "net"
)

// Gets a valid system IP address.
func get_ip() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, addr := range addrs {

        fmt.Printf("ADDR STRING: %s\n", addr.String())
        if _, ipnet, err := net.ParseCIDR(addr.String()); err == nil && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                fmt.Printf("IP RET: %s\n", ipnet.IP.String())
                //return ipnet.IP.String()
            }
        }

    }
    return ""
}
