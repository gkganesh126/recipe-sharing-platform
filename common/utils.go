package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
)

type (
	/*
		appError struct {
			Error   string `json:"error"`
			Message string `json:"message"`
			HttpStatus int    `json:"status"`
		}
			errorResource struct {
			Data errorResponse `json:"data"`
		}*/
	configuration struct {
		Server                        string `json:"server"`
		MongoDBHost                   string `json:"mongoDBHost"`
		MongoDBUser                   string `json:"mongoDBUser"`
		MongoDBPwd                    string `json:"mongoDBPwd"`
		RecipeSharingPlatformDatabase string `json:"recipeSharingPlatformDatabase"`
	}
)

/*
{
  "server"      : "0.0.0.0:8081",
  "mongoDBHost" : "localhost",
  "mongoDBUser"	: "",
  "mongoDBPwd"	: "",
  "ohnoDatabase" : "admin"
}
*/

func DisplayAppError(w http.ResponseWriter, username string, handlerError error, message string, code int) {
	zap.S().Errorf("Username: %s, GroupID: %s, message: %s, error: %s", username, message, handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(message); err == nil {
		w.Write(j)
	}
}

// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// Initialize AppConfig
func initConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		zap.S().Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[logAppConfig]: %s\n", err)
	}
}

// Session holds the mongodb session for database access
var session *mgo.Session

// Get database session
func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{AppConfig.MongoDBHost},
			Username: AppConfig.MongoDBUser,
			Password: AppConfig.MongoDBPwd,
			Database: AppConfig.RecipeSharingPlatformDatabase,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			zap.S().Fatalf("[GetSession]: %s\n", err)
		}
	}
	return session
}

// Create database session
func createDbSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConfig.MongoDBHost},
		Username: AppConfig.MongoDBUser,
		Password: AppConfig.MongoDBPwd,
		Database: AppConfig.RecipeSharingPlatformDatabase,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		zap.S().Fatalf("[createDbSession]: %s\n", err)
	}
	/*
		//fmt.Println("AppConfig: ", AppConfig)
		copySession := session.Copy()
		if err := copySession.DB(AppConfig.OhnoDatabase).Login(AppConfig.MongoDBUser, AppConfig.MongoDBPwd); err != nil {
			zap.S().Fatalf("[createDbSession] Login: %s\n", err)
		}
	*/
}
