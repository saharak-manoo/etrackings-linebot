package app

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/saharak/etrackings-linebot/app/controllers"
)

var server = controllers.Server{}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return fmt.Sprintf(":%s", port)
}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Println(fmt.Sprintf("Error getting env, %v", err))
	} else {
		log.Println("We are getting values")
	}

	server.Initialize()

	apiPort := getPort()
	log.Print(fmt.Sprintf("Running on ENV -> %s", strings.ToUpper(os.Getenv("FIBER_MODE"))))
	log.Print(fmt.Sprintf("Listening to port -> %s", apiPort))

	server.Run(apiPort)
}
