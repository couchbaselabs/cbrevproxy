package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type CBRevProxy struct {

	// target url of reverse proxy
	target *url.URL

	// instance of Go ReverseProxy thatwill do the job for us
	proxy *httputil.ReverseProxy

}


func New(target string) (*CBRevProxy, error) {

	url, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	cbRevProxy := &CBRevProxy{
		target: url,
		proxy: httputil.NewSingleHostReverseProxy(url),
	}

	return cbRevProxy, nil
}

func (p *CBRevProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-GoProxy", "GoProxy")

	p.proxy.ServeHTTP(w, r)

}
