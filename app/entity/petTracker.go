package entity

import "time"

type Pet struct {
	ID      int        `json:"id"`
	Name    string     `json:"name" validate:"required"`
	Owner   string     `json:"owner" validate:"required"`
	Species string     `json:"species" validate:"required"`
	Birth   *time.Time `json:"birth,omitempty"`
	Death   *time.Time `json:"death,omitempty"`
}

type Event struct {
	PetID  int       `json:"pet_id" validate:"required"`
	Date   time.Time `json:"date" validate:"required"`
	Type   string    `json:"type" validate:"required"`
	Remark string    `json:"remark,omitempty"`
}
