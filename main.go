package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	var url string
	fmt.Scanln(&url)
	getBody(url)
	//fmt.Printf(body)
}

func getBody(url string) string {
	var resp *http.Response = makeRequest(url)
	var body string
	var err error
	body, err = readResponse(resp)
	if err != nil {
		return body
	} else {
		//fmt.Printf("there is an error: %d", err)
		return "error"
	}
}

func makeRequest(url string) *http.Response {
	var resp *http.Response
	var err error
	resp, err = http.Get(url)
	if err != nil {
		//fmt.Printf("error:%d", err)
		return nil
	}
	return resp
}

func readResponse(resp *http.Response) (string, error) {
	var body []byte
	var err error

	body, err = io.ReadAll(resp.Body)
	fmt.Printf(string(body))
	return string(body), err
}
