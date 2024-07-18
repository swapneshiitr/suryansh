package routes

import (
	"github.com/gorilla/mux"
	"github.gom/swapneshiitr/suryansh/api/controllers"
)

var UserDetailRoutes = func(router *mux.Router) {
	router.HandleFunc("/userdetail/{flat}", controllers.GetUserDetail).Methods("GET")
	router.HandleFunc("/user/", controllers.CreateOrEditUser).Methods("POST")
}
