package repository

import (
	dbCon "github.com/franck-djacoto/announce-service/db-connection"
	. "github.com/franck-djacoto/announce-service/models"

)

type AnnonceRepository struct {
	DbConnect *dbCon.DbConnection
}

func (annonceRepo *AnnonceRepository) Add (annonce AnnonceModel) (annonceCreatedId int64, err error){
	db := annonceRepo.DbConnect.Db

	query := `INSERT INTO annonce('id', 'titre', 'contenu', 'categorie_id', 'marque_id', 'model_id', 'created_at'  ) VALUES( ?, ?, ? , ?, ?, ?, ?)`
	result, err := db.Exec(query, annonce.Id, annonce.Titre, annonce.Contenu, annonce.CategorieId, annonce.MarqueId, annonce.ModeleId, annonce.CreatedAt)

	if err != nil {
		return 0, err
	}

	annonceCreatedId, err = result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return annonceCreatedId, nil
}

func (annonceRepo *AnnonceRepository) getById(annonceId int) Annonce {
	db := annonceRepo.DbConnect.Db
	var annonce Annonce

	query := `SELECT an.id, an.titre, an.contenu, cat.libelle, mq.libelle, md.libelle, an.created_at 
			  FROM annonce as an INNER JOIN categories as cat ON an.categorie_id = cat.id
			  INNER JOIN modele as md ON an.model_id = md.id 
			  INNER JOIN marque as mq ON an.marque_id = mq.id
			  WHERE an.id = ?`

	row := db.QueryRow(query, annonceId)
	row.Scan(&annonce.Id, &annonce.Titre, &annonce.Contenu, &annonce.Categorie, &annonce.Marque, &annonce.Modele, &annonce.CreatedAt, &annonce.UpdatedAt )

	if annonce.Id > 0 {
		return annonce
	}

	return Annonce{}
}


func (annonceRepo *AnnonceRepository) getAll() ( []Annonce, error) {
	db := annonceRepo.DbConnect.Db
	var allAnnonce []Annonce
	var annonce Annonce
	query := `SELECT an.id, an.titre, an.contenu, cat.libelle, mq.libelle, md.libelle, an.created_at 
			  FROM annonce as an INNER JOIN categories as cat ON an.categorie_id = cat.id
			  INNER JOIN modele as md ON an.model_id = md.id 
			  INNER JOIN marque as mq ON an.marque_id = mq.id`

	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next(){
		rows.Scan(&annonce.Id, &annonce.Titre, &annonce.Contenu, &annonce.Categorie, &annonce.Marque, &annonce.Modele, &annonce.CreatedAt, &annonce.UpdatedAt )
		allAnnonce = append(allAnnonce, annonce)
	}

	return allAnnonce, nil
}

func (annonceRepo *AnnonceRepository) Delete( annonceId int ) (bool, error){
	db := annonceRepo.DbConnect.Db
	query := `DELETE FROM annonce WHERE id = ?`
	result, err := db.Exec( query )

	if err != nil {
		return false, err
	}

	rowAffected, err := result.RowsAffected()

	if rowAffected > 0 {
		return true, nil
	}

	return false, nil
}