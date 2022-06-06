package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"merchant-api/utils/types"
	"os"
	"path/filepath"
)

var dbConnections map[string]types.DBConnection

// LoadConfigSettings -
func LoadConfigSettings() {
	localPath := "config"
	fileName := "development.json"

	absPath, absErr := filepath.Abs(localPath)
	if absErr != nil {
		log.Panic("error occurred while getting abs file", absErr)
	}
	path := absPath + "/" + fileName
	LoadJSON(path)
}

// LoadJSON -
func LoadJSON(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("error occurred while loading conf file", err)
		return
	}
	rawBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("error occurred while reading conf file", err)
	}
	var config types.Config
	err = json.Unmarshal(rawBytes, &config)
	if err != nil {
		log.Println("error occurred while reading conf file", err)
	}
	dbConnections = make(map[string]types.DBConnection, 0)
	for _, conn := range config.Connections {
		dbConnections[conn.Database] = conn
	}
}

func GetDBCredentials(key string) types.DBConnection {
	if len(dbConnections) == 0 {
		LoadConfigSettings()
	}

	return dbConnections[key]
}
