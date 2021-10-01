package repository

import (
	dbCon "github.com/franck-djacoto/announce-service/db-connection"
	"github.com/franck-djacoto/announce-service/models"
)

type CategoryRepository struct {
	DbConnect *dbCon.DbConnection
}

func (catRepo *CategoryRepository) GetAll() ([]models.Categorie, error) {
	db := catRepo.DbConnect.Db
	var allCat []models.Categorie
	var singleCat models.Categorie
	query := `select id, libelle from categories`
	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next(){
		err = rows.Scan(&singleCat.Id, singleCat.Libelle)
		if err != nil {
			return nil, err
		}

		allCat =  append(allCat, singleCat)
	}
	return allCat, nil
}


func (catRepo *CategoryRepository) GetCatIdByLibelle(catLibelle string) (int, error) {
	db := catRepo.DbConnect.Db
	var categorieId int
	query := `select id from categories where libelle = ?`
	row := db.QueryRow(query, catLibelle)
	err := row.Scan(&categorieId)
	if err != nil {
		return 0, err
	}
	return categorieId, nil
}

func (catRepo *CategoryRepository) GetCatLibelleById(catId int) (string, error) {
	db := catRepo.DbConnect.Db
	var catLibelle string
	query := `select id from categories where id = ?`
	row := db.QueryRow(query, catId)
	err := row.Scan(&catLibelle)
	if err != nil {
		return "", err
	}
	return catLibelle, nil
}


