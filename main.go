package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "time"
)

func ping(w http.ResponseWriter, req *http.Request) {
    defer req.Body.Close()
    sleepDuration := 0.0
    requestBody := map[string]interface{}{}
    if req.Body != nil {
        bodyBytes, _ := ioutil.ReadAll(req.Body)
        err := json.Unmarshal(bodyBytes, &requestBody)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        var ok bool
        sleepDuration, ok = requestBody["sleep_duration"].(float64)
        if !ok {
            w.WriteHeader(http.StatusBadRequest)
            return
        }
    }

    time.Sleep(time.Duration(sleepDuration) * time.Millisecond)
    w.Write([]byte("pong"))
    w.WriteHeader(http.StatusOK)
    return
}

func main() {
    http.HandleFunc("/ping", ping)

    http.ListenAndServe(":8080", nil)
}