package handlers

import (
    "encoding/json"
    "net/http"
)

// Response represents the structure of the response sent back to the client
type Response struct {
    Message string `json:"message"`
}

// HelloHandler handles requests to the /hello endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    response := Response{Message: "Hello from Go microservice!"}
    json.NewEncoder(w).Encode(response)
}