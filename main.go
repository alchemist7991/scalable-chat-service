package main

import (
	"fmt"
	"log"
	"os"
	wsServer "github.com/alchemist7991/scalable-chat-service/server"
	"github.com/alchemist7991/scalable-chat-service/constant"
)

func ConfigureLogger() {
	logsPath := constant.LOGS_PATH
	logFile, err := os.OpenFile(logsPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Unable to create log file")
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
}

func main() {
	ConfigureLogger()
	wsServer.StartServer()
}
