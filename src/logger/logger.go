package logger

import (
	"log"
	"os"
)

func InitLogger() *os.File {

	LOG_FILE := "./LibreRemotePlay.log"

	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal("logger file can not be opened")
	}

	log.SetOutput(logFile)

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	return logFile

}
