package main

import (
	"fmt"
	"net/http"
)

type WebServer struct {
	Port      string
	BasePath  string
	Endpoints map[string]RequestHandler
}

type WebServerBuilder struct {
	port      string
	basePath  string
	endpoints map[string]RequestHandler
}

func (ws *WebServerBuilder) Port(port string) *WebServerBuilder {
	ws.port = port
	return ws
}

func (ws *WebServerBuilder) BasePath(basePath string) *WebServerBuilder {
	ws.basePath = basePath
	return ws
}

func (ws *WebServerBuilder) EndPoints(endpoints map[string]RequestHandler) *WebServerBuilder {
	ws.endpoints = endpoints
	return ws
}

func (ws *WebServerBuilder) Build() *WebServer {
	return &WebServer{
		Port:      ws.port,
		BasePath:  ws.basePath,
		Endpoints: ws.endpoints,
	}
}

func (ws *WebServer) Run() {
	var router = http.NewServeMux()

	for k, _ := range ws.Endpoints {
		var endpoint = ws.BasePath + k
		router.HandleFunc(endpoint, genericErrorHandler(ws.genericHandler))
		fmt.Println(endpoint)
	}

	var port = ":" + ws.Port

	http.ListenAndServe(port, router)
}

func (wsb *WebServerBuilder) Run() {
	var ws = wsb.Build()
	ws.Run()
}

type ServerError struct {
	Error string
}

type RequestHandler func(http.ResponseWriter, *http.Request) error

func genericErrorHandler(handler RequestHandler) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if error := handler(response, request); error != nil {
			fmt.Println("Error!!!")
		}
	}
}

func (ws *WebServer) genericHandler(response http.ResponseWriter, request *http.Request) error {
	switch request.Method {
	case http.MethodGet:
		fmt.Println("GET called!")
	case "POST":
		fmt.Println("POST called!")
	case "PUT":
		fmt.Println("PUT called!")
	case "DELETE":
		fmt.Println("DELETE called!")
	}

	return nil
}
