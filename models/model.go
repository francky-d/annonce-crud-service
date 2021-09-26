package models

import "time"

type Model struct {
	Id int64 `json:"id"`
	Libelle string `json:"libelle"`
	Marque Marque  `json:"marque"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}