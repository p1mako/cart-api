package database

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
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

func init() {
	connectionInfo, err := getConfig()
	if err != nil {
		panic("Unable to open connection with db: error while loading config")
	}
	db, err = sqlx.Open("postgres", parseToConStr(connectionInfo))

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