package models

import (
	"gorm.io/gorm"
)

type Weather struct {
	gorm.Model
	Water       float64
	Wind        float64 
	WaterStatus string  	
	WindStatus  string 
}