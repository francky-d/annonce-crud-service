package repository

import dbCon "github.com/franck-djacoto/announce-service/db-connection"

type CategoryRepository struct {
	DbConnect *dbCon.DbConnection
}

func (catRepo *CategoryRepository) getCatIdByLibelle (catLibelle string) (int, error) {
	db := catRepo.DbConnect.Db
	var categorieId int
	query := `select id from categories where libelle = ?`
	row := db.QueryRow(query, catLibelle)
	err := row.Scan(&categorieId)
	if err != nil {
		return 0 , err
	}
	return categorieId, nil
}


func (catRepo *CategoryRepository) getCatLibelleById ( catId int) (string, error) {
	db := catRepo.DbConnect.Db
	var catLibelle string
	query := `select id from categories where id = ?`
	row := db.QueryRow(query, catId)
	err := row.Scan(&catLibelle)
	if err != nil {
		return "" , err
	}
	return catLibelle, nil
}