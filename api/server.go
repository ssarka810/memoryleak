package api

import (
	"net/http"

	_ "net/http/pprof"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)
type Server struct{
	router *mux.Router
}

func NewServer() *Server{
	return &Server{}
}

func (server *Server)Init(){
server.router=mux.NewRouter().StrictSlash(true)
server.RouteApi()
}

func (server *Server)StartServer(port string){
	// logrus.Info("starting the server with the port : ",port)
	// logrus.Fatal(http.ListenAndServe(":"+port,server.router))

	errCh := make(chan error)
	go func (ch chan error){
		logrus.Info("starting the server with the port : ",port)
		errCh <- http.ListenAndServe(":"+port,server.router)
	}(errCh)

	go func (ch chan error){
    logrus.Info("starting profiling server with port 8080")
    errCh <- http.ListenAndServe(":8080",nil)
	}(errCh)

	for {
		logrus.Fatal(<- errCh)
	}
}