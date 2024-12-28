package utils

import (
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

type Response struct {
	StatusCode int
	Header     http.Header
	Body       string
}

func HttpRequest(url string, method string, headers http.Header, body string) (*Response, error) {
	client := resty.New()
	r := client.R().EnableTrace()

	for k, v := range headers {
		r.Header.Add(k, strings.Join(v, ","))
	}

	var (
		res *resty.Response
		err error
	)

	switch strings.ToUpper(method) {
	case "GET":
		res, err = r.Get(url)
	case "POST":
		res, err = r.SetBody(body).Post(url)
	case "PUT":
		res, err = r.SetBody(body).Put(url)
	case "DELETE":
		res, err = r.Delete(url)
	}

	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: res.StatusCode(),
		Header:     res.RawResponse.Header,
		Body:       string(res.Body()),
	}, nil
}
