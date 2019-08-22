package main

import (
	_"fmt"
	"net/http"
	_"os"
	"time"
	"github.com/guotooyoung/gwp_chitchat_demo/src/webexpand"
)

func main() {
	//fmt.Println(os.Getwd())
	webexpand.P("ChitChat", webexpand.Version(), "started at", webexpand.Config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(webexpand.Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", webexpand.Index)
	// error
	mux.HandleFunc("/err", webexpand.Err)

	// defined in route_auth.go
	mux.HandleFunc("/login", webexpand.Login)
	mux.HandleFunc("/logout", webexpand.Logout)
	mux.HandleFunc("/signup", webexpand.Signup)
	mux.HandleFunc("/signup_account", webexpand.SignupAccount)
	mux.HandleFunc("/authenticate", webexpand.Authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", webexpand.NewThread)
	mux.HandleFunc("/thread/create", webexpand.CreateThread)
	mux.HandleFunc("/thread/post", webexpand.PostThread)
	mux.HandleFunc("/thread/read", webexpand.ReadThread)

	// starting up the server
	server := &http.Server{
		Addr:           webexpand.Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(webexpand.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(webexpand.Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
