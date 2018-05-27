//concierge/main.go
package main
import (
    "net/http"
    "concierge/api"
    "concierge/common"
)

func main() {
    common.Log("Adding API handlers...")
    http.HandlerFunc("/api/fedder", api.FeedHandler)
    common.Log("starting feeder...")
    api.StartFeederSystem()
    common.Log("starting DocumentSearch Concierge server on port :8080")
    http.ListenAndServe(":8080", nil)
}
