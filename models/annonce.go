package models

import "time"

type AnnonceModel struct {
	Id          int64     `json:"id"`
	Titre       string    `json:"titre"`
	Contenu     string    `json:"contenu"`
	CategorieId int       `json:"categorie"`
	MarqueId    int       `json:"marque"`
	ModeleId    int       `json:"modele"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Annonce struct {
	Id        int64     `json:"id"`
	Titre     string    `json:"titre" binding:"required"`
	Contenu   string    `json:"contenu" binding:"required"`
	Categorie string    `json:"categorie" binding:"required"`
	Marque    string    `json:"marque"`
	Modele    string    `json:"modele"`
}
