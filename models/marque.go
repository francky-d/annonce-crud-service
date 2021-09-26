package models

import "time"

type Marque struct {
	Id int64 `json:"id"`
	Libelle string `json:"libelle"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}