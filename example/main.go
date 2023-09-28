package main

import (
	"github.com/Ashfakh/gogonet/server"
)

//	func main() {
//		ln, _ := net.Listen("tcp", "0.0.0.0:8080")
//		for {
//			c, _ := ln.Accept()
//			buf := make([]byte, 1024)
//			for {
//				n, _ := c.Read(buf)
//				if n == 0 {
//					break
//				}
//				c.Write(buf[:n])
//			}
//		}
//	}

func request(req server.Request) server.Response {
	if req.Params()["name"] != "" {
		return server.NewResponse(200, "Hello "+req.Params()["name"])
	}
	return server.NewResponse(200, "Hello World")
}

func main() {

	var serv = server.Server("0.0.0.0", 8080, 1, 10)
	serv.Route(server.GET, "/", request)
	serv.Route(server.GET, "/hello/:name", request)
	serv.Route(server.GET, "/hello", request)
	serv.Route(server.POST, "/hell", request)
	serv.Run()
}
