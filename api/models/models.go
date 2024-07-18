package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Flat            int16       `gorm:"primaryKey;not null;autoIncrement:false" json:"flat"`
	Resident        string      `gorm:"not null" json:"resident"`
	OwnerName       string      `gorm:"not null" json:"owner_name"`
	OwnerPhone      int64       `gorm:"not null" json:"owner_phone"`
	OwnerAlternate  int64       `json:"owner_alternate"`
	RenterName      string      `json:"renter_name"`
	RenterPhone     int64       `json:"renter_phone"`
	RenterAlternate int64       `json:"renter_alternate"`
	Complaints      []Complaint `gorm:"foreignKey:RaisedBy"`
}

type Complaint struct {
	ID        int16     `gorm:"primaryKey;not null;autoIncrement:true" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	RaisedBy  int16     `json:"raised_by"`
	Subject   string    `gorm:"not null" json:"subject"`
	Body      string    `gorm:"not null" json:"body"`
	Status    string    `gorm:"default:'Pending';not_null" json:"status"`
	Remarks   []Remark  `gorm:"foreignKey:ComplaintID"`
	// Redundant - The default foreign key’s name is the owner’s type name(i.e. Complaint) plus the name of its primary key(i.e. ID) field
}

type Remark struct {
	ID          int16 `gorm:"primaryKey;not null;autoIncrement:true" json:"id"`
	ComplaintID int16 `json:"complaint_id"`
	// GORM usually uses the owner’s primary key as the foreign key’s value, for the above example, it is the Complaint‘s ID
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
}
