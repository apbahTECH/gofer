package main

import (
    "fmt"
    "html/template"
    "net/http"
)

/* Notes
 *  - for handlers you can look at http.Request.Method to determine Get or Put etc.
 */

var clients []Client

// Register clients that send heartbeats with the server.
func heartbeatHandler(w http.ResponseWriter, r *http.Request) {
    client := Client{get_ip(r)}
    fmt.Println("Got client: %s\n", client.ip)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("dashboard.html")
    if err == nil {
        if t == nil {
            fmt.Println("t is nil!")
        } else {
            t.Execute(w, r)
        }
    } else {
        fmt.Printf("Error parsing template: %s", err)
    }
}

func main() {
    var port = 8080

    http.HandleFunc("/", dashboardHandler)
    http.HandleFunc("/dashboard", dashboardHandler)
    http.HandleFunc("/heartbeat", heartbeatHandler)

    fmt.Printf("Web service started on port %d...", port)
    http.ListenAndServe(":8080", nil)
}
