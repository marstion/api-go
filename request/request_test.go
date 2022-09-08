package request

import (
	"log"
	"testing"
)

func TestRequest(t *testing.T) {
	response := Request("GET", "http://httpbin.org/ip", nil, "")
	log.Printf("response: %#v\n", string(response))
}
