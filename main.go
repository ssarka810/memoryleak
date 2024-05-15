package main

import (
	"github.com/sirupsen/logrus"
	api "github.com/ss530n/go-rest-prac/api"
)

func main(){
	logrus.Info(" testing rest related advance concepts")
	server :=api.NewServer()
	server.Init()
	server.StartServer("9999")

}