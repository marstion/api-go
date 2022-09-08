package request

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Request(method string, url string, header map[string][]string, data string) (bodyText []byte) {
	client := &http.Client{}
	var body = strings.NewReader(data)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return
}
