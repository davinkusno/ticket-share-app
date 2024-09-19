package models

import "gorm.io/gorm"

type Event struct {
    gorm.Model
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Date        string  `json:"date"`
    Price       float64 `json:"price"`
}
