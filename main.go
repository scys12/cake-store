package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/scys12/cake-store/config"
	"github.com/scys12/cake-store/database"
	"github.com/scys12/cake-store/internal/api/handler"
	"github.com/scys12/cake-store/internal/repositories"
	"github.com/scys12/cake-store/internal/server"
	"github.com/scys12/cake-store/internal/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	dbConfig, err := config.InitDBConfig()
	if err != nil {
		log.Info("[Config] Failed to load db config")
	}
	servConf, err := config.InitServerConfig()
	if err != nil {
		log.Info("[Config] Failed to load server config")
	}

	db := database.GetDatabaseConnection(dbConfig)

	cakeRepo := repositories.NewCakeRepo(db)
	cakeService := service.NewCakeService(cakeRepo)
	cakeHandler := handler.NewCakeHandler(cakeService)

	router := mux.NewRouter()
	router.HandleFunc("/cakes", cakeHandler.InsertCake).Methods(http.MethodPost).Name("insertCake")
	router.HandleFunc("/cakes/{id}", cakeHandler.UpdateCake).Methods(http.MethodPatch).Name("updateCake")
	router.HandleFunc("/cakes/{id}", cakeHandler.DeleteCake).Methods(http.MethodDelete).Name("deleteCake")

	serverConfig := server.Config{
		WriteTimeout: time.Duration(servConf.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(servConf.ReadTimeout) * time.Second,
		Port:         servConf.Port,
	}
	server.Serve(serverConfig, router)
}
