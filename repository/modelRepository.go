package repository

import dbCon "github.com/franck-djacoto/announce-service/db-connection"

type ModelRepository struct {
	DbConnect *dbCon.DbConnection
}

func (modelRepo *ModelRepository) getModelIdByLibelle ( modelLibelle string ) (int, error) {
	db := modelRepo.DbConnect.Db
	var modelId int
	query := `select id from models where libelle = ?`
	row := db.QueryRow(query, modelLibelle)
	err := row.Scan(&modelId)
	if err != nil {
		return 0 , err
	}
	return modelId, nil
}


func (modelRepo *ModelRepository) getModelLibelleById ( modelId int ) (string, error) {
	db := modelRepo.DbConnect.Db
	var modelLibelle string
	query := `select id from models where id = ?`
	row := db.QueryRow(query, modelId)
	err := row.Scan(&modelLibelle)
	if err != nil {
		return "" , err
	}
	return modelLibelle, nil
}