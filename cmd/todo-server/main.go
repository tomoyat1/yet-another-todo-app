package main

import (
	"github.com/tomoyat1/yet-another-todo-app/server"
	"github.com/tomoyat1/yet-another-todo-app/postgres"
	"github.com/tomoyat1/yet-another-todo-app/config"
)

const (
	serverPort = 8080
)

var (
	dbConf *config.PgConfig
)

func init() {
	var err error
	dbConf, err = config.DBConfigFromEnv()
	if err!= nil {
		panic("Couldn't read DB config from environment: " + err.Error())
	}

	if err!= nil {
		panic("Couldn't read DB config from environment: " + err.Error())
	}
}

func main() {
	repo, err := postgres.NewPgItemRepositoryImpl(dbConf.GenerateConnectionString())
	if err != nil {
		panic("failed to create item repository: " + err.Error())
	}
	s, err := server.NewServer(repo)
	if err != nil {
		panic("failed to create Server: " + err.Error())
	}
	if err := s.Start(serverPort); err != nil {
		panic("server failed: " + err.Error())
	}
}
