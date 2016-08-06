package main

import (
    "fmt"
    "net/http"
    "net/url"
    "io/ioutil"
    "time"
)

// Parses values in POST request body into url.Values structure.
func getFormValues(r *http.Request) (url.Values, error) {
    // Read function body from HTTP Request.
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        if options.debug {
            fmt.Printf("Could not read request body.\n")
            fmt.Printf("Error: %s\n", err)
        }
        return url.Values{}, err
    }
    if options.debug {
        fmt.Printf("Body: %s\n", body)
    }

    // Parse body into form values.
    // The Form/PostForm fields in the HTTP Request do not appear
    // to be valid. This may be due to a limitation on our server.
    values, err := url.ParseQuery(string(body[:]))
    if err != nil {
        if options.debug {
            fmt.Printf("Could not parse request body to url.Values.\n")
            fmt.Printf("Error: %s\n", err)
        }
    }
    return values, err
}

// Return Client structure form heartbeat POST data.
func getClientFromRequest(r *http.Request) (Client, error) {
    values, err := getFormValues(r)
    name := ""
    ip := ""
    if err != nil {
        if options.debug {
            fmt.Printf("Form values were nil. Client invalid.\n")
            fmt.Printf("Error: %s\n", err)
        }
    } else {
        name = values.Get("name")
        ip = values.Get("ip")
    }
    return Client{name, ip, time.Now()}, err
}
