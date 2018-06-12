package main

import (
	"github.com/tomoyat1/yet-another-todo-app/server"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		panic("failed to create Server: " + err.Error())
	}
	if err := s.Start(); err != nil {
		panic("server failed: " + err.Error())
	}
}
