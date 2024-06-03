package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type dbInfo struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Dbname   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var pathToConfig = "./internal/config/db-conf.json"
var db *sqlx.DB

func InitDb() {
	connectionInfo, err := getConfig()
	if err != nil {
		log.Fatalf("unable to open connection with db: error while loading config")
	}
	log.Printf("Got database config for db %v\n", connectionInfo.Dbname)
	db, err = sqlx.Open("postgres", parseToConStr(connectionInfo))
	if err != nil {
		log.Fatalf("unable to open connection with db %v\n", err.Error())
	}
	log.Printf("Opened connection with db %v\n", connectionInfo.Dbname)

}

func ConnectDB() *sqlx.DB {
	return db
}

func parseToConStr(connectionInfo *dbInfo) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", connectionInfo.Host, connectionInfo.Port, connectionInfo.User, connectionInfo.Password, connectionInfo.Dbname)
}

func getConfig() (*dbInfo, error) {
	var connectionInfo dbInfo
	confFile, err := os.Open(pathToConfig)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(confFile).Decode(&connectionInfo)
	return &connectionInfo, err
}

func Close() error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
