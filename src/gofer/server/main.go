package main

import (
    "fmt"
    "html/template"
    "net/http"
    "flag"
    "io/ioutil"
    "encoding/json"
    "time"
    "sync"
)

/* Notes
 *  - for handlers you can look at http.Request.Method to determine Get or Put etc.
 */

var clientsMutex sync.Mutex
var clients []Client

type Options struct {
    debug bool
}
var options Options

func parseCmdline() {
    dbg := flag.Bool("d", false, "Enable debug messages.")
    flag.Parse()

    options.debug = *dbg

    if(options.debug) {
        fmt.Printf("Options:\n")
        fmt.Printf("    Debug: Enabled\n")
    }
}

// Register clients that send heartbeats with the server.
func heartbeatHandler(w http.ResponseWriter, r *http.Request) {
    client, err := getClientFromRequest(r)
    if err != nil {
        fmt.Printf("Failed to create Client from heartbeat: %s.", err)
    }
    if options.debug {
        fmt.Printf("Client Heartbeat Received:\n")
        fmt.Printf("    IP: %s\n", client.Ip)
        fmt.Printf("    Name: %s\n", client.Name)
    }
    clientsMutex.Lock()
    found := false
    for i, c := range clients {
        if c.Ip == client.Ip {
            found = true
            // Note that range creates a copy into c, cannot modify directly.
            clients[i].Timestamp = time.Now()
            if options.debug {
                fmt.Println("Client already registered.\n")
            }
        }
    }
    if !found {
        if options.debug {
            fmt.Printf("Adding new client!\n")
        }
        clients = append(clients, client)
    }
    clientsMutex.Unlock()
}

// Displays the skeleton of the dashboard page.
// This is what is included in dashboard.html.
// The data displayed on the dashboard page is actually retrieved
// via AJAX from other handlers.
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("dashboard.html")
    if err != nil {
        if options.debug {
            fmt.Printf("Error parsing dashboard.html template: %s\n", err)
        }
    } else {
        t.Execute(w, r)
    }
}

// Grabs data describing what clients are currently registered with this
// server. This is the data that the AJAX calls will consume.
func getClientsData() []byte {
    // Output format is JSON.
    // [{"name" : "NAME", "ip" : "10.0.0.1"}, ...]
    output := []byte{'['}
    for _, client := range clients {
        b, err := json.Marshal(client)
        if err != nil {
            if options.debug {
                fmt.Println("Error marshalling client to JSON %s\n", err)
            }
            continue
        }
        output = append(output, b...)
        output = append(output, ',')
    }
    if output[len(output)-1] == ',' {
        output[len(output)-1] = ']'
    } else {
        output = append(output, ']')
    }
    if options.debug {
        fmt.Printf("Fully marshalled clients data:  %s\n", output)
    }
    return output
}

// Handles AJAX data requests. These requests fill data for the
// included web pages.
func ajaxDataHandler(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        if options.debug {
            fmt.Printf("Error in AJAX data handler: %s\n", err)
        }
        return
    }
    if options.debug {
        fmt.Printf("HTTP Request: %s\n", body)
    }

    // Structure for our JSON to get unmarshalled into.
    // If the JSON does not have a field in the DataRequest, or
    // the JSON has an extra field not in the DataRequest, it is
    // simply ignored. Nice and simple!
    type DataRequest struct {
        Operation string
    }
    var dr DataRequest
    err = json.Unmarshal(body, &dr)
    if err != nil {
        if options.debug {
            fmt.Printf("Failed to unmarshal JSON\n")
            fmt.Printf("Error: %s\n", err)
        }
    }

    if options.debug {
        fmt.Printf("Unmarshalled JSON:\n")
        fmt.Printf("    Operation: %s\n", dr.Operation)
    }

    switch dr.Operation {
    case "clients":
        w.Write(getClientsData())
    default:
        if options.debug {
            fmt.Printf("Unexpected Data Request operation %s.\n", dr.Operation)
        }
    }
}

func main() {
    var port = 8080

    parseCmdline()

    // Launch client monitoring thread.
    go monitorClients()

    // Register handlers and start web service.
    http.HandleFunc("/", dashboardHandler)
    http.HandleFunc("/dashboard", dashboardHandler)
    http.HandleFunc("/heartbeat", heartbeatHandler)
    http.HandleFunc("/data", ajaxDataHandler)

    fmt.Printf("Web service started on port %d...", port)
    http.ListenAndServe(":8080", nil)
}
