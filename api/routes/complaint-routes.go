package routes

import (
	"github.com/gorilla/mux"
	"github.gom/swapneshiitr/suryansh/api/controllers"
)

var ComplaintRoutes = func(router *mux.Router) {
	router.HandleFunc("/raisecomplaint/", controllers.RaiseComplaint).Methods("POST")
	router.HandleFunc("/complaints/", controllers.GetComplaints).Methods("GET")
	// router.HandleFunc("/complaint/{id}", controllers.GetComplaintDetail).Methods("GET")
	// router.HandleFunc("/addremark/{id}", controllers.AddRemark).Methods("POST")
	// router.HandleFunc("/resolve/{id}", controllers.ResolveComplaint).Methods("POST")
	// router.HandleFunc("deletecomplaint/{id}", controllers.DeleteComplaint).Methods("POST")
}
