package app

import "github.com/Andreykurniawan/tokoevent/app/controllers"

func main(server *Server) initializeRoutes() {

	server.Router.HandleFunc( path: "/", controllers.Home).Methods( methods...: "GET")
	
}