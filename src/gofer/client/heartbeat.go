package main

import (
    "fmt"
    "net/http"
    "net/url"
    "time"
)

// Sends heartbeat to the server. Runs forever.
func heartbeat(ip string, name string) {
    for {
        if(debug) {
            fmt.Printf("Sending heartbeat to %s.\n", server_ip)
        }

        resp, err := http.PostForm("http://" + server_ip + "/heartbeat",
            url.Values{"ip" : {ip}, "name" : {name}})
        if err != nil {
            if(debug) {
                fmt.Printf("Error on Post: %s\n", err)
            }
        } else {
            resp.Body.Close() // Must close response.
        }
        // Send a heartbeat every 2 seconds.
        time.Sleep(time.Duration(2) * time.Second)
    }
}
