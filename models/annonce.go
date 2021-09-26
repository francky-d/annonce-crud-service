package models

import "time"

type Annonce struct {
	Id int64 `json:"id"`
	Titre string `json:"titre"`
	Contenu Categorie `json:"contenu"`
	CategorieId int  `json:"categorie"`
	MarqueId int `json:"marque,omitempty"`
	ModeleId int `json:"modele,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}