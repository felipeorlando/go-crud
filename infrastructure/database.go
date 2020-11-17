package database

import (
	"com.go-crud/config"
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

var session mgo.Session

// ConnectDB creates a session and connect to the database
func ConnectDB() *mgo.Database {
	session, err := mgo.Dial(config.DbURI)
	if err != nil {
		log.Fatalln("Error on dial DB:", err)
	}

	err = session.Ping()
	if err != nil {
		log.Fatalln("Error on Ping DB:", err)
	}

	log.Println("MongoDB session started successfully")

	return session.DB("users-crud")
}

// CloseDB close a connection with MongoDB
func CloseDB() {
	log.Println("Stop MongoDB session")
	session.Close()
}
