package main

import server "github.com/Ashfakh/gogonet"

func request(req server.Request) server.Response {
	if req.Params()["name"] != "" {
		return server.NewResponse(200, "Hello "+req.Params()["name"])
	}
	return server.NewResponse(200, "Hello World")
}

func main() {
	var serv = server.Server("0.0.0.0", 8080, 1, 10) // host, port, numHandlers, maxQueued
	serv.Route(server.GET, "/", request)
	serv.Route(server.GET, "/hello/:name", request)
	serv.Route(server.GET, "/hello", request)
	serv.Route(server.POST, "/hell", request)
	serv.Run()
}
