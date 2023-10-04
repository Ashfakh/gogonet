# GoGonet: A Lightweight HTTP Server in Go 🚀

Welcome to GoGonet! A simple, yet powerful lightweight HTTP server written in Go. Whether you're prototyping or developing a minimal application, GoGonet is here to make your development smoother.

## 🌟 Features
- 🚀 Fast server initialization
- 🛣️ Dynamic URL routing with path and query parameters
- 🧵 Multithreaded request handling
- 🛡️ Buffer overflow protection with max queued requests

## 🚀 Getting Started

### 🔧 Prerequisites
- [Go](https://golang.org/doc/install) (>= 1.15)

### 📦 Installation
1. Fetch the Go package:
```bash
$ go get github.com/Ashfakh/gogonet
```

2. Import the package in your Go code:
```bash
$ import server "github.com/Ashfakh/gogonet"
```

### 💼 Basic Usage

1. Initialize a server instance:
```go
serv := server.Server("0.0.0.0", 8080, 1, 10) // host, port, numHandlers, maxQueued
```

2. Define routes and associated request handlers:
```go
serv.Route(server.GET, "/", requestHandlerFunc)
```

3. Fire up the server! 🚀
```go
serv.Run()
```

### 🖋️ Sample Request Handler
```go
func request(req server.Request) server.Response {
    if req.Params()["name"] != "" {
        return server.NewResponse(200, "Hello "+req.Params()["name"]+" 🖐️")
    }
    return server.NewResponse(200, "Hello World 🌍")
}
```

## 📑 Files Overview

- `server.go` 🔧 - Contains the primary server code, including structures and methods required for setting up and running the HTTP server.

- `httpHandler.go` 💌 - Takes charge of parsing the incoming requests and routing them to the appropriate request handlers.

- The `Implementation` block is a practical example of how to wield this server library in a Go application.

## 👤 Author
[Ashfakh](https://github.com/Ashfakh) ✨