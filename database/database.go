package database

import (
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

var session mgo.Session

// ConnectDB creates a session and connect to the database
func ConnectDB(dbName string, dbURI string) *mgo.Database {
	session, err := mgo.Dial(dbURI)
	if err != nil {
		log.Fatalln("Error on dial DB:", err)
	}

	err = session.Ping()
	if err != nil {
		log.Fatalln("Error on Ping DB:", err)
	}

	return session.DB(dbName)
}

// CloseDB close a connection with MongoDB
func CloseDB() {
	session.Close()
}
