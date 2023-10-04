package server

import (
	"fmt"
	"net"
)

type ServerObj struct {
	host                   string
	port                   int
	numHandlers, maxQueued int
	router                 router
}

func Server(host string, port int, numHandlers int, maxQueued int) ServerObj {
	fmt.Println("Initialising Server...")
	server := ServerObj{
		host:        host,
		port:        port,
		numHandlers: numHandlers,
		maxQueued:   maxQueued,
		router:      make([]route, 0),
	}
	return server
}

func (s *ServerObj) Run() error {
	addr := s.host + ":" + fmt.Sprintf("%v", s.port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	fmt.Printf("Server listening on %v\n", addr)
	var requestQueue = make(chan net.Conn, s.maxQueued)
	for i := 0; i < s.numHandlers; i++ {
		go httpHandler(i, s, requestQueue)
	}

	for {
		cl, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		requestQueue <- cl
	}
}
