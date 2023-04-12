package request

import (
	"testing"
)

func TestRequest(t *testing.T) {
	statusCode, response := Request("GET", "http://httpbin.org/ip", nil, "")
	t.Log(statusCode, string(response))
}
