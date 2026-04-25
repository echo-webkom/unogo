package unogo

import "net/http"

var globalHttpClient = &http.Client{}

type httpClient struct {
	*http.Client
}

func newHttpClient() *httpClient {
	return &httpClient{globalHttpClient}
}
