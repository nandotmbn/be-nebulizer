package models

import (
	"time"
)

type Nebulizer struct {
	NebulizerName string    `json:"nebulizer_name,omitempty" bson:"nebulizer_name,omitempty" validate:"required,min=0"`
	Password      string    `json:"password,omitempty" validate:"required,min=3,max=255"`
	State         bool      `json:"state,omitempty"`
	Battery       float32   `json:"battery,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
