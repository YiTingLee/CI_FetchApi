package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	httpGet("http://140.124.39.202:30001/api/onsale")
}

func httpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
