package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.gom/swapneshiitr/suryansh/api/routes"
)

const LOCALHOST = "127.0.0.1"
const SERVER_PORT = "9010"

func main() {
	fileServer := http.FileServer(http.Dir("./web"))
	http.Handle("/", fileServer)

	r := mux.NewRouter()
	routes.UserDetailRoutes(r)
	routes.ComplaintRoutes(r)

	server := LOCALHOST + ":" + SERVER_PORT
	fmt.Println("Starting server at " + server)
	log.Fatal(http.ListenAndServe(server, r))
}
