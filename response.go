package server

import (
	"strconv"
)

type Response struct {
	httpVersion string
	statusCode  int
	headers     map[string]string
	body        string
}

func NewResponse(statusCode int, body string) Response {
	var res Response
	res.headers = make(map[string]string)
	res.statusCode = statusCode
	res.body = body
	return res
}

func (res *Response) addServerHeaders(httpVersion string) {
	res.headers["Content-Length"] = strconv.Itoa(len([]byte(res.body)))
	res.httpVersion = httpVersion
}

func (res *Response) AddHeader(key, val string) {
	res.headers[key] = val
}

func (res *Response) SetStatusCode(code int) {
	res.statusCode = code
}

func (res *Response) SetBody(body string) {
	res.body = body
}

func (res *Response) toByteSlice() []byte {
	var httpResponseString string
	httpResponseString += res.httpVersion + " " + strconv.Itoa(res.statusCode) + " " + statusCode[res.statusCode] + "\r\n"
	for k, v := range res.headers {
		httpResponseString += k + ": " + v + "\r\n"
	}
	httpResponseString += "\r\n"
	httpResponseString += res.body
	return []byte(httpResponseString)
}
