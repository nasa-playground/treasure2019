package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	method := flag.String("X", "GET", "http method option")
	header := flag.String("H", ":", "http header")

	flag.Parse()

	url := flag.Arg(0)

	request, err := http.NewRequest(*method, url, nil)
	headers := strings.Split(*header, ":")

	if len(headers) > 2 {
		request.Header.Add(headers[0], headers[1])
	}

	if err != nil {
		fmt.Println(err)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func ParseUrl(u string) bool {
	_, err := url.Parse(u)
	if err != nil || u == "" {
		return false
	}

	return true
}
