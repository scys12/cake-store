package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/scys12/cake-store/config"
	"github.com/scys12/cake-store/database"
	"github.com/scys12/cake-store/internal/api/handler"
	"github.com/scys12/cake-store/internal/repositories"
	"github.com/scys12/cake-store/internal/server"
	"github.com/scys12/cake-store/internal/service"
	"github.com/scys12/cake-store/pkg/monitoring"
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

	if err := monitoring.Init(); err != nil {
		log.Fatalln("[Monitoring] Unable to initialize monitoring")
	}

	cakeRepo := repositories.NewCakeRepo(db)
	cakeService := service.NewCakeService(cakeRepo)
	cakeHandler := handler.NewCakeHandler(cakeService)

	router := mux.NewRouter()
	router.HandleFunc("/cakes", cakeHandler.InsertCake).Methods(http.MethodPost).Name("insertCake")
	router.HandleFunc("/cakes/{id}", cakeHandler.UpdateCake).Methods(http.MethodPatch).Name("updateCake")
	router.HandleFunc("/cakes/{id}", cakeHandler.DeleteCake).Methods(http.MethodDelete).Name("deleteCake")
	router.HandleFunc("/cakes/{id}", cakeHandler.GetCakeDetail).Methods(http.MethodGet).Name("getCakeDetail")
	router.Handle("/prometheus", promhttp.Handler())
	wrappedMux := monitoring.Middleware(router)

	serverConfig := server.Config{
		WriteTimeout: time.Duration(servConf.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(servConf.ReadTimeout) * time.Second,
		Port:         servConf.Port,
	}
	server.Serve(serverConfig, wrappedMux)
}
