package types

import "gopkg.in/mgo.v2/bson"

// types.Map
type Map bson.M

//Holds the page information
type Page struct {
	Page  int
	Limit int
	Skip  int
}

// DB Connection
type DBConnection struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	Connections []DBConnection `json:"connections"`
}
