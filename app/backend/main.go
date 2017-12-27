package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CzarSimon/util"
)

// Env is the struct for environment objects
type Env struct {
	DB     *sql.DB
	Config Config
}

// SetupEnv Sets up environment
func SetupEnv(config Config) *Env {
	db := util.ConnectPG(config.db)
	return &Env{
		DB:     db,
		Config: config,
	}
}

// SetupServer Creates a server with a route handler
func SetupServer(env *Env) *http.Server {
	return &http.Server{
		Addr:    ":" + env.Config.server.Port,
		Handler: SetupRoutes(env),
	}
}

func main() {
	config := getConfig()
	env := SetupEnv(config)

	server := SetupServer(env)

	log.Printf("Starting backend on port: %s\n", config.server.Port)
	err := server.ListenAndServe()
	util.CheckErrFatal(err)
}
