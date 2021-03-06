package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

const (
	adminPrefix = "/admin"
	publicPrefix = "/api"
)

type Server struct {
	HTTPServer *http.Server
}

var server Server

func SetupServer(addr string) error {
	r := mux.NewRouter()
	pub := r.PathPrefix(publicPrefix).Subrouter()
	pub.HandleFunc("/test", TestHandler)
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	server.HTTPServer = srv
	return srv.ListenAndServe()
}