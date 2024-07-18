package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.gom/swapneshiitr/suryansh/api/models"
	"github.gom/swapneshiitr/suryansh/internal/database"
	"github.gom/swapneshiitr/suryansh/utils"
)

func GetComplaints(w http.ResponseWriter, r *http.Request) {
	complaint := database.GetComplaints()
	res, _ := json.Marshal(complaint)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func RaiseComplaint(w http.ResponseWriter, r *http.Request) {
	temp := models.Complaint{}
	utils.ParseBody(r, temp)
	c := database.RaiseComplaint(&temp)
	if c != nil {
		fmt.Println("User registered successfully!")
		res, _ := json.Marshal(c)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		fmt.Println("Complaint Already Exists!")
	}
}
