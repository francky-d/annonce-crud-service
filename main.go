package main

import (
	dbCon "github.com/franck-djacoto/announce-service/db-connection"
	"github.com/franck-djacoto/announce-service/routes"
	"github.com/joho/godotenv"
	"log"
)

type App struct {
	db dbCon.DbConnection
}
func LoadEnv(){
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error while loading .env file : %v", err)
	}
}

func init(){

}

func main(){
	LoadEnv()
	var app App
	app.db = dbCon.NewConnection()
	app.db.MakeMigration()
	routes.RouterInit()

}