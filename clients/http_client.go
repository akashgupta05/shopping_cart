package clients

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var client *http.Client
var once sync.Once

type HTTPClient struct {
	Host   string
	client *http.Client
}

type HTTPClientInterface interface {
	Send(method, path string) ([]byte, error)
}

func NewHTTPClient(host string) *HTTPClient {
	return &HTTPClient{
		Host:   host,
		client: getHTTPClient(),
	}
}

func getHTTPClient() *http.Client {
	once.Do(func() {
		client = &http.Client{
			Timeout: time.Second,
		}
	})

	return client
}

func (hc *HTTPClient) Send(method, path string) ([]byte, error) {
	return []byte(``), nil
}

func (hc *HTTPClient) buildRequest(method, path string) (*http.Request, error) {
	http.NewRequest(method, fmt.Sprintf("%s/%s", hc.Host, path), nil)

	return nil, nil
}
