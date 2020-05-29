package api

import (
	"fmt"
	"notification-app/api/database"
	"notification-app/api/routers"
	s "notification-app/api/settings"

	"log"
	"net/http"
)

// Run is a main entry point
func Run() {

	s.DB = database.GetDB()
	defer s.DB.Close()

	s.Router = routers.SetupRouter()

	apiPort := fmt.Sprintf(":%s", s.Env["API_PORT"])

	fmt.Printf("Starting API server on port %s\n", apiPort)
	log.Fatal(http.ListenAndServe(apiPort, s.Router))
}
