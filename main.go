package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	proxyTargetUrlStr := "http://localhost:4984"
	proxyTargetUrl, err := url.Parse(proxyTargetUrlStr)
	if err != nil {
		panic(fmt.Sprintf("Error parsing url: %v", err))
	}

	revProxy := httputil.NewSingleHostReverseProxy(proxyTargetUrl)

	http.Handle("/", revProxy)

	log.Fatal(http.ListenAndServe(":49840", nil))

}
