package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	port string
}

type ServerError struct {
	Error string
}

type errorHandler func(http.ResponseWriter, *http.Request) error

func httpHandler(handler errorHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, reader *http.Request) {
		if error := handler(writer, reader); error != nil {
			fmt.Println("Error!!!")
		}
	}
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (server *Server) Run() {
	var router = mux.NewRouter()

	router.HandleFunc("/", httpHandler(server.handleLogin))

	http.ListenAndServe(server.port, router)
}

func (server *Server) handleLogin(writer http.ResponseWriter, reader *http.Request) error {
	switch reader.Method {
	case "GET":
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
