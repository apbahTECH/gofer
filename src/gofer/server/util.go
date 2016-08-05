package main

import (
    "fmt"
    "net/http"
)

func get_ip(r *http.Request) string {
    fmt.Printf("IP: %s\n", r.RemoteAddr)
    return r.RemoteAddr
}
