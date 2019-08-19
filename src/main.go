package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/guotooyoung/gwp_chitchat_demo/src/webexpand"
)

func main() {
	fmt.Println("chitchat begin")
	webexpand.p("ChitChat", webexpand.version(), "started at", webexpand.config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(webexpand.config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", webexpand.index)
	// error
	mux.HandleFunc("/err", webexpand.err)

	// defined in route_auth.go
	mux.HandleFunc("/login", webexpand.login)
	mux.HandleFunc("/logout", webexpand.logout)
	mux.HandleFunc("/signup", webexpand.signup)
	mux.HandleFunc("/signup_account", webexpand.signupAccount)
	mux.HandleFunc("/authenticate", webexpand.authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", webexpand.newThread)
	mux.HandleFunc("/thread/create", webexpand.createThread)
	mux.HandleFunc("/thread/post", webexpand.postThread)
	mux.HandleFunc("/thread/read", webexpand.readThread)

	// starting up the server
	server := &http.Server{
		Addr:           webexpand.config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(webexpand.config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(webexpand.config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
