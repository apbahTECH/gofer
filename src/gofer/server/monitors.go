package main

import (
    "fmt"
    "time"
)

// Monitors clients registered with the server. If a client hasn't sent a
// heartbeat within the time limit, it is removed from server tracking.
func monitorClients() {
    for {
        clientsMutex.Lock()
        for index, client := range clients {
            if options.debug {
                fmt.Printf("Client Timestamp: %s\n", client.Timestamp)
                fmt.Printf("Time Now: %s\n", time.Now())
                fmt.Printf("Time Since: %s\n", time.Since(client.Timestamp))
            }
            // The time limit for no heartbeat is 10 seconds.
            if time.Since(client.Timestamp).Seconds() > float64(10) {
                    clients = append(clients[:index], clients[index+1:]...)
                    if options.debug {
                        fmt.Printf("Client was removed!\n")
                        fmt.Printf("    IP: %s\n", client.Ip)
                        fmt.Printf("    Name: %s\n", client.Name)
                    }
            }
            //TODO: what if this client was doing something?
            //      we need to handle that case, clean way to "cancel"
            //      and just call that here
        }
        clientsMutex.Unlock()
        // The monitor kicks into action every 10 seconds.
        time.Sleep(time.Duration(10) * time.Second)
    }
}
