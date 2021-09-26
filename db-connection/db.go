package db_connection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type DbConnection struct {
	Db *sql.DB
}

func  NewConnection() DbConnection{
	userName := os.Getenv("DB_USER")
	userPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	if userName == "" || userPassword == "" || dbHost == "" || dbName == "" {
		log.Fatal("You must set all env variable for db connection : DB_USER, DB_PASSWORD, DB_HOST, DB_NAME")
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8",
									userName, userPassword,dbHost,dbName ) // "username:password@(127.0.0.1:3306)/dbname?parseTime=true"
	newDbConnection, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Panicf("Could not open connection %v", err)
	}

	if err = newDbConnection.Ping(); err != nil {
		log.Panicf("Error while pinging.. %v", dataSourceName)
	}

	return DbConnection{Db : newDbConnection}
}

func (dbConnect *DbConnection) MakeMigration(){
	content , err := ioutil.ReadFile("./db-connection/migration.sql")

	if err != nil {
		log.Panicf("Error while retreiving query for migration : %v", err )
	}

	if string(content) == "" {
		log.Panicf("Migration call without queries")
	}

	allQueries := strings.Split( string(content), ";")

	for _,query := range allQueries {
		if strings.TrimSpace(query) != "" {
			_, err := dbConnect.Db.Exec( query )
			if  err != nil {
				log.Panicf("Error while migrating : %v", err )
			}
		}

	}

	fmt.Println("Migration was successfull!")

}