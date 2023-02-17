package ports

import "net/http"

type BeefService interface {
	Count() map[string]int
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
