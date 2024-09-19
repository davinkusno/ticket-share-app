package models

import "gorm.io/gorm"

type Cart struct {
    gorm.Model
    UserID   uint    `json:"user_id"`
    EventID  uint    `json:"event_id"`
    Event    Event   `json:"event"`
    Quantity int     `json:"quantity"`
}
