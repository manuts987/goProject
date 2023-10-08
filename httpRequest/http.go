package httprequest

import (
	"fmt"
	"io"
	"net/http"
)

func GetBodyRequest(url string) string {
	var resp *http.Response = makeRequest(url)
	var body string
	var err error
	body, err = readbody(resp)
	if err == nil {
		return body
	} else {
		fmt.Print(err)
		return "error"
	}
}

func GetheaderRequest(url string) string {
	var resp *http.Response = makeRequest(url)
	var header string
	var err error
	header, err = readheader(resp)
	if err == nil {
		return header
	} else {
		return "error"
	}
}

func readheader(resp *http.Response) (string, error) {
	var header []byte
	var err error
	fmt.Print(resp.Header)
	//header, err = io.ReadAll(resp.Header)
	return string(header), err
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

func readbody(resp *http.Response) (string, error) {
	var body []byte
	var err error

	body, err = io.ReadAll(resp.Body)
	return string(body), err
}
