package providers

import (
	"bytes"
	"log"
	"net/http"
)

// Basic external provider abstraction. This could grow into a full module, defining Connectors and endpoints separately from a shared interface
type ExternalProviderRouting struct {
	BaseUrl     string
	ContentType string
	// Headers
	// Private keys
	// Public keys
}

// Post request creator
func (r *ExternalProviderRouting) Post(endpoint string, b []byte) (*http.Response, error) {
	var fullUrl = r.BaseUrl + endpoint
	var jsonBody = bytes.NewBuffer(b)
	log.Printf("| Outgoing Request POST | URL: %s | json: %s", fullUrl, jsonBody)
	return http.Post(fullUrl, r.ContentType, jsonBody)
}
