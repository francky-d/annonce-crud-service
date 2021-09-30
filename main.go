package main

import (
	. "github.com/franck-djacoto/announce-service/conf"
	dbCon "github.com/franck-djacoto/announce-service/db-connection"
	"github.com/franck-djacoto/announce-service/routes"
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error while loading .env file : %v", err)
	}
}

func init() {

}

func main() {
	LoadEnv()
	var App Application
	App.Db = dbCon.NewConnection()
	App.Db.MakeMigration()
	routes.RouterInit(&App)

}
