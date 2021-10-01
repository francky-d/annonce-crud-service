package repository

import (
	"errors"
	"fmt"
	dbCon "github.com/franck-djacoto/announce-service/db-connection"
	. "github.com/franck-djacoto/announce-service/models"
	"time"
)

type AnnonceRepository struct {
	DbConnect *dbCon.DbConnection
}

func (annonceRepo *AnnonceRepository) Save(annonce AnnonceModel) (annonceCreatedId int64, err error) {
	db := annonceRepo.DbConnect.Db

	query := `INSERT INTO annonce(titre, contenu, categorie_id, marque_id, model_id, created_at ) VALUES( ?, ?, ? , ?, ?, ?)`
	result, err := db.Exec(query, annonce.Titre, annonce.Contenu, annonce.CategorieId, annonce.MarqueId, annonce.ModeleId, time.Now())

	if err != nil {
		return 0, err
	}

	annonceCreatedId, err = result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return annonceCreatedId, nil
}

func (annonceRepo *AnnonceRepository) GetById(annonceId int) (Annonce, error) {
	db := annonceRepo.DbConnect.Db
	var annonce Annonce

	query := `SELECT an.id, an.titre, an.contenu, cat.libelle, mq.libelle, md.libelle
			  FROM annonce as an INNER JOIN categories as cat ON an.categorie_id = cat.id
			  INNER JOIN models as md ON an.model_id = md.id 
			  INNER JOIN marques as mq ON an.marque_id = mq.id
			  WHERE an.id = ?`

	row := db.QueryRow(query, annonceId)
	err := row.Scan(&annonce.Id, &annonce.Titre, &annonce.Contenu, &annonce.Categorie, &annonce.Marque, &annonce.Modele)
	if err != nil {
		return Annonce{}, err
	}

	if annonce.Id > 0 {
		return annonce, nil
	}

	return Annonce{}, nil
}

func (annonceRepo *AnnonceRepository) GetAll() ([]Annonce, error) {
	db := annonceRepo.DbConnect.Db
	var allAnnonce []Annonce
	var annonce Annonce
	query := `SELECT an.id, an.titre, an.contenu, cat.libelle, mq.libelle, md.libelle
			  FROM annonce as an INNER JOIN categories as cat ON an.categorie_id = cat.id
			  INNER JOIN models as md ON an.model_id = md.id 
			  INNER JOIN marques as mq ON an.marque_id = mq.id`

	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&annonce.Id, &annonce.Titre, &annonce.Contenu, &annonce.Categorie, &annonce.Marque, &annonce.Modele)

		if err != nil {
			return nil, err
		}

		allAnnonce = append(allAnnonce, annonce)

	}

	return allAnnonce, nil
}

func (annonceRepo *AnnonceRepository) Delete(annonceId int) (bool, error) {
	db := annonceRepo.DbConnect.Db
	query := `DELETE FROM annonce WHERE id = ?`
	result, err := db.Exec(query, annonceId)

	if err != nil {
		return false, err
	}

	rowAffected, err := result.RowsAffected()

	if rowAffected > 0 {
		return true, nil
	}

	return false, nil
}

func (annonceRepo *AnnonceRepository) Update(annonce AnnonceModel) (bool, error) {
	db := annonceRepo.DbConnect.Db
	query := `UPDATE annonce SET titre=?, contenu=?, categorie_id=?, marque_id=?, model_id=?, updated_at=?`
	result, err := db.Exec(query, annonce.Titre, annonce.Contenu, annonce.CategorieId, annonce.MarqueId, annonce.ModeleId, time.Now())

	if err != nil {
		return false, err
	}

	rowAffected, err := result.RowsAffected()

	if rowAffected > 0 {
		return true, nil
	}

	return false, nil
}

func (annonceRepo *AnnonceRepository) GetByModelName(model string) ([]Annonce, error) {
	db := annonceRepo.DbConnect.Db
	modelRepo := ModelRepository{DbConnect: annonceRepo.DbConnect}
	modelId, err := modelRepo.GetModelIdByLibelle(model)
	if err != nil {
		return nil, err
	}

	if modelId > 0 {
		var allAnnonces []Annonce

		query := `SELECT an.id, an.titre, an.contenu, cat.libelle, mq.libelle, md.libelle
				  FROM annonce as an INNER JOIN categories as cat ON an.categorie_id = cat.id
				  INNER JOIN models as md ON an.model_id = md.id 
				  INNER JOIN marques as mq ON an.marque_id = mq.id
				  WHERE md.id = ?`

		rows, err := db.Query(query, modelId)
		if err != nil {
			return nil, err
		}


		for rows.Next(){
			var annonce Annonce
			rows.Scan(&annonce.Id, &annonce.Titre, &annonce.Contenu, &annonce.Categorie, &annonce.Marque, &annonce.Modele)
			allAnnonces =  append(allAnnonces, annonce)
		}

		if len(allAnnonces) > 0 {
			return allAnnonces, nil
		}
	}

	return nil, errors.New( fmt.Sprintf("Announce not found for model %s", model) )
}
