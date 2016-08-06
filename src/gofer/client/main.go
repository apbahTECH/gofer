package main

import (
    "fmt"
    "flag"
)

var server_ip string
var debug bool

func parseCmdline() {
    sip := flag.String("s", "", "IP address of gofer server. Specify port if server is not running on port 80.")
    dbg := flag.Bool("d", false, "Enable debug messages.")
    flag.Parse()

    server_ip = *sip
    debug = *dbg

    if(debug) {
        fmt.Printf("Options:\n")
        fmt.Printf("    Debug: Enabled\n")
        fmt.Printf("    Server IP: %s\n", server_ip)
    }
}

func main() {
    fmt.Printf("gofer client started.\n")

    parseCmdline()

    ip := getIp()
    name := getHostname()
    if(debug) {
        fmt.Printf("Client IP: %s\n", ip)
        fmt.Printf("Client Hostname: %s\n", name)
    }

    heartbeat(ip, name)
}
