package domain

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street     string `json:"street"`
	Number     int    `json:"number"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}
