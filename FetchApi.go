package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Serve int

func (m Serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := ioutil.ReadFile("data.json")
	if err != nil {

	}
	fmt.Fprintln(w, string(res))
}

func main() {
	data := httpGet("http://140.124.39.202:30001/api/onsale")
	err := ioutil.WriteFile("data.json", []byte(data), 0644)

	if err != nil {
		// handle error
	}

	var s Serve
	http.ListenAndServe(":8080", s)
}

func httpGet(url string) (data string) {
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

	return string(body)
}
