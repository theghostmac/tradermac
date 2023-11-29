package alphavantage

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Client interface {
	MakeRequest(path, method, reqBody string, headers map[string]string) (*http.Request, error)
	DoRequest(r *http.Request) (*http.Response, error)
}

var _ Client = &RequestClient{}

type RequestClient struct {
	dataHost   string
	httpClient *http.Client
}

func NewRequestClient(dataHost string, httpClient *http.Client) *RequestClient {
	return &RequestClient{
		dataHost:   dataHost,
		httpClient: httpClient,
	}
}

func (client *RequestClient) MakeRequest(path, method, reqBody string, headers map[string]string) (*http.Request, error) {
	var (
		url  = fmt.Sprintf("%s%s", client.dataHost, path)
		req  *http.Request
		err  error
		body io.Reader
	)

	if reqBody != "" {
		body = strings.NewReader(reqBody)
	}

	req, err = http.NewRequest(method, url, body)
	if err != nil {
		log.Printf("error processing request: %v", err)
		firstErr := errors.New(err.Error())
		secondErr := errors.New("base request error on new request")
		// Join the errors for better tracking on large request data.
		fullErrorMsg := errors.Join(firstErr, secondErr)
		return nil, fullErrorMsg
	}

	if body == nil && method == http.MethodPut {
		req.Header.Add("Content-Length", "0")
	}

	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	log.Printf("headers: '%s\n' url:'%s\n' body:'%s\n'", headers, url, body)
	return req, nil
}

func (client *RequestClient) DoRequest(r *http.Request) (*http.Response, error) {
	start := time.Now()
	res, err := client.httpClient.Do(r)
	if err != nil {
		log.Printf("cannot do request: '%v' from client", err)
		return nil, err
	}

	duration := time.Since(start)
	log.Printf("request took '%v ' to process", duration)
	return res, nil
}
