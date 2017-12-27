package main

import (
	"log"
	"net/http"

	"github.com/CzarSimon/util"
)

// SetupRoutes Sets up routes and handlers
func SetupRoutes(env *Env) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/comment", env.HandleCommentRequests)
	mux.HandleFunc("/api/health", env.HealthCheck)
	return mux
}

// HealthCheck Checks that the server is responsive and
// has connection with the database
func (env *Env) HealthCheck(res http.ResponseWriter, req *http.Request) {
	err := env.DB.Ping()
	if err != nil {
		log.Printf("Liveness probe failed. Error: %s", err.Error())
		util.SendErrStatus(res, err, http.StatusInternalServerError)
		return
	}
	log.Println("Liveness probe ok")
	util.SendOK(res)
}
