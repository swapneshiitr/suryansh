package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.gom/swapneshiitr/suryansh/api/models"
	"github.gom/swapneshiitr/suryansh/internal/database"
	"github.gom/swapneshiitr/suryansh/utils"
)

func GetUserDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["flat"] // probably string
	flatNum, err := strconv.ParseUint(ID, 0, 0)
	if err != nil {
		panic(err)
	}
	userDetails := database.GetUserDetail(uint16(flatNum))
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateOrEditUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utils.ParseBody(r, user)
	f, u := database.CreateOrEditUser(user)
	if f {
		fmt.Println("User registered successfully!")
	} else {
		fmt.Println("User edited successfully!")
	}
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
