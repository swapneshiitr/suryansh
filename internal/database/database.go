package database

import (
	"github.com/jinzhu/gorm"
	"github.gom/swapneshiitr/suryansh/api/models"
	"github.gom/swapneshiitr/suryansh/internal/config"
)

var db *gorm.DB

// initialise DB
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&models.User{}, &models.Complaint{}, &models.Remark{})
}

// User Service
func CreateOrEditUser(u *models.User) (bool, *models.User) {
	var create bool = false
	if db.NewRecord(u) {
		db.Create(u) // You cannot pass a struct to ‘create’, so you should pass a pointer to the data.
		create = true
	} else {
		db.Save(u)
	}
	return create, u
}

func GetUserDetail(f uint16) *models.User {
	var details models.User
	db.Where("Flat=?", f).Find(&details)
	return &details
}

// Complaint Service
func RaiseComplaint(c *models.Complaint) *models.Complaint {
	if db.NewRecord(c) {
		db.Create(&c)
		return c
	}
	return nil
}

func GetComplaints() []models.Complaint {
	var complaints []models.Complaint
	db.Find(&complaints)
	return complaints
}
