package server

import (
	"fmt"
	"net/url"
	"strings"
)

type router []route

type route struct {
	path    string
	method  Method
	handler func(Request) Response
}

func (s *ServerObj) Route(method Method, path string, f func(Request) Response) {
	s.router = append(s.router, route{path, method, f})
}

func (s *ServerObj) get(path string, requestMethod string) (*route, map[string]string, *RequestError) {
	for _, r := range s.router {
		if match, params := pathMatch(r.path, r.method.String(), path, requestMethod); match {
			// Extract query params from the path
			queryParams, _ := extractQueryParams(path)
			// Merge the path params and query params
			for k, v := range queryParams {
				params[k] = v
			}
			return &r, params, nil
		}
	}
	return nil, nil, &RequestError{fmt.Sprintf("no router found for path: %s and method: %s", path, requestMethod), 400}
}

func pathMatch(pattern string, method string, path string, requestMethod string) (bool, map[string]string) {
	if method != requestMethod {
		return false, nil
	}

	params := make(map[string]string)

	// Split the path from any query strings
	pathWithoutQuery := strings.Split(path, "?")[0]

	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(pathWithoutQuery, "/")

	if len(patternParts) != len(pathParts) {
		return false, nil
	}

	for i, patternPart := range patternParts {
		if strings.HasPrefix(patternPart, ":") {
			params[patternPart[1:]] = pathParts[i]
		} else if patternPart != pathParts[i] {
			return false, nil
		}
	}
	return true, params
}

func extractQueryParams(path string) (map[string]string, error) {
	params := make(map[string]string)
	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	values := u.Query()
	for k, v := range values {
		params[k] = v[0] // Only considering the first value for each query parameter
	}
	return params, nil
}
