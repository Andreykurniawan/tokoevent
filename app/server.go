package app

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	Router *mux.Router
	server.initializeRoutes()
}
func (server *Server) Initialize() {
	fmt.Println("Selamat Datang di Toko Event")

	server.Router = mux.NewRouter()
}
	
func (server *Server) Run(addr string) {
	fmt.Printf ("Server berjalan pada %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	var server = Server{}

	server.Initialize()
	server.Run( ":9000")
}