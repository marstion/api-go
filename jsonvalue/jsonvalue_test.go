package jsonvalue

import (
	"fmt"
	"testing"
)

var srcByte []byte = []byte(`{"status": 0, "msg": "OK", "data": {"name": "Jack", "age": 18}, "error": ""}`)
var j J

func TestUnmarshal(t *testing.T) {
	j.Unmarshal2Self(srcByte)
	fmt.Printf("obj: %#v\n", j.obj)
}

func TestGet(t *testing.T) {
	j.Unmarshal2Self(srcByte)
	fmt.Printf("%v", j.Get("msg"))
}

func TestUnmarshal1Self(t *testing.T) {
	var v struct {
		Status int
		Msg    string
		Error  string
	}
	j.Unmarshal2Self(srcByte)
	j.Unmarshal1Self(&v)

	fmt.Printf("v: %#v", v)
}
