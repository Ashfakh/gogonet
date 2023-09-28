package server

import "strings"

type Request struct {
	method      string
	path        string
	httpVersion string
	headers     map[string]string
	params      map[string]string
}

func (req *Request) addHeader(key, val string) {
	req.headers[key] = val
}

func (rq *Request) Path() string {
	return rq.path
}

func (rq *Request) Headers() map[string]string {
	return rq.headers
}

func (rq *Request) Method() string {
	return rq.method
}

func (rq *Request) HTTPVersion() string {
	return rq.httpVersion
}

func (rq *Request) Params() map[string]string {
	return rq.params
}

func parseRequest(requestInput []string) (*Request, *RequestError) {
	var request Request
	request.headers = make(map[string]string)
	if len(requestInput) == 0 {
		return nil, &RequestError{"No Headers present in the Request", 400}
	}
	httpFields := strings.Fields(requestInput[0])
	if len(httpFields) < 3 {
		return nil, &RequestError{"HTTP Standards missing in the Request", 400}
	}
	request.method = httpFields[0]
	request.path = httpFields[1]

	request.httpVersion = httpFields[2]

	for i := 1; i < len(requestInput); i++ {
		tmp := strings.Split(requestInput[i], ":")
		if len(tmp) < 2 {
			return nil, &RequestError{"Error Parsing the Headers", 400}
		}
		request.addHeader(strings.TrimSpace(tmp[0]), strings.TrimSpace(tmp[1]))
	}
	return &request, nil
}
