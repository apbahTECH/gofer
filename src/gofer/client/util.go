package main

import (
    "net"
    "os"
)

// Gets the system hostname.
func getHostname() string {
    hostname, err := os.Hostname()
    if err != nil {
        return ""
    }
    return hostname
}

// Gets a valid system IP address.
func getIp() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, addr := range addrs {
        if ip, _, err := net.ParseCIDR(addr.String()); err == nil && !ip.IsLoopback() && ip.To4() != nil {
            if ip != nil {
                return ip.String()
            }
        }
    }
    return ""
}
