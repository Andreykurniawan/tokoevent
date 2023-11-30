package main

func main(server *server) initializeRoutes() {

	server.Router.HandleFunc( path "/", controllers.Home).Methods(methods...: "GET")
	
}