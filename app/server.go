package main

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	Router *mux.Router
}
func (server *server) Initialize() {
	fmt.Println(a...: "Selamat Datang di Toko Event")

	server.Router = mux.NewRouter()
}
	
func (server *server) Run(addr string) {
	fmt.Printf ("Server berjalan pada %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	var server server{}

	server.Initialize()
	server.Run(addr ":9000")
}