package server

import (
	"bufio"
	"net"
)

func httpHandler(id int, server *ServerObj, requestQueue <-chan net.Conn) {
	for {
		cl := <-requestQueue

		reader := bufio.NewReader(cl)
		requestLines := make([]string, 0)
		for {
			buffer, _, _ := reader.ReadLine()
			if string(buffer) == "" {
				break
			}
			requestLines = append(requestLines, string(buffer))
		}
		request, err := parseRequest(requestLines)
		response := NewResponse(200, "")
		if err != nil {
			response = NewResponse(err.ErrorCode(), err.Error())
			response.addServerHeaders("HTTP/1.1") // Defaulting to HTTP/1.1
		} else {
			r, params, err := server.get(request.path, request.method)
			request.params = params
			if err != nil {
				response = NewResponse(err.ErrorCode(), err.Error())
			} else {
				response = r.handler(*request)
			}
			response.addServerHeaders(request.httpVersion)
		}
		cl.Write(response.toByteSlice())
		cl.Close()
	}
}
