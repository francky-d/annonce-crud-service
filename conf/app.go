package conf

import dbCon "github.com/franck-djacoto/announce-service/db-connection"

type Application struct {
	Db *dbCon.DbConnection
}
