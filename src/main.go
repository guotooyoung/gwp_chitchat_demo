package main

import (
	"net/http"
	"time"
	"webexpand"
)

func main() {
	expand.p("ChitChat", expand.version(), "started at", expand.config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(expand.config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", expand.index)
	// error
	mux.HandleFunc("/err", expand.err)

	// defined in route_auth.go
	mux.HandleFunc("/login", expand.login)
	mux.HandleFunc("/logout", expand.logout)
	mux.HandleFunc("/signup", expand.signup)
	mux.HandleFunc("/signup_account", expand.signupAccount)
	mux.HandleFunc("/authenticate", expand.authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", expand.newThread)
	mux.HandleFunc("/thread/create", expand.createThread)
	mux.HandleFunc("/thread/post", expand.postThread)
	mux.HandleFunc("/thread/read", expand.readThread)

	// starting up the server
	server := &http.Server{
		Addr:           expand.config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(expand.config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(expand.config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
