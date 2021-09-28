package repository

import dbCon "github.com/franck-djacoto/announce-service/db-connection"

type MarqueRepository struct {
	DbConnect *dbCon.DbConnection
}

func (marqueRepo *MarqueRepository) getMarqueIdByLibelle (marqueLibelle string) (int, error) {
	db := marqueRepo.DbConnect.Db
	var marqueId int
	query := `select id from marques where libelle = ?`
	row := db.QueryRow(query, marqueLibelle)
	err := row.Scan(&marqueId)
	if err != nil {
		return 0 , err
	}
	return marqueId, nil
}


func (marqueRepo *MarqueRepository) getMarqueLibelleById (marqueId int) (string, error) {
	db := marqueRepo.DbConnect.Db
	var marqueLibelle string
	query := `select id from models where id = ?`
	row := db.QueryRow(query, marqueId)
	err := row.Scan(&marqueLibelle)
	if err != nil {
		return "" , err
	}
	return marqueLibelle, nil
}