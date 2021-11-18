package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main()  {
	urlString := "https://www.google.com/search/newpath?q=golang&oq=golang&aqs=chrome..69i57j69i60l4j69i65l3.1353j0j1&sourceid=chrome&ie=UTF-8"
	url, _ := url.Parse(urlString)

	fmt.Println(url.Scheme)
	fmt.Println(url.Host)
	fmt.Println(strings.Split(url.Path, "/"))

	fmt.Println(url.Query()["sourceid"][0])
	fmt.Println(url.Query()["q"][0])
}
