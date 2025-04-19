package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	endPoint := base.ResolveReference(reference).String()
	fmt.Println(endPoint)

	req, _ := http.NewRequest("GET", endPoint, nil)
	req.Header.Add("Header", "aaaa")
	q := req.URL.Query()
	fmt.Println(q)
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
}
