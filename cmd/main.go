package main

//Каркас микросервиса
//Реализовать "каркас" микросервиса, считывающий конфиг из файла,
//создающий логгер/логгеры с указанными уровнями детализации.

import (
	serviceConfig "github.com/iae94/Microservice/internal/config"
	"github.com/iae94/Microservice/internal/routes"
	"github.com/iae94/Microservice/internal/service"
	"log"
	"net/http"
)

func main() {
	//Service struct
	webService := service.WebService{}

	//Read config
	config, err := serviceConfig.ReadConfig()
	if err != nil {
		log.Fatalf("Reading config error: %v \n", err)
	}

	//Init webService with config
	err = webService.Init(config)
	if err != nil {
		log.Fatalf("WebService initialization has failed with error: %v\n", err)
	}

	//Register handlers
	webService.RegisterHandlers(&map[string]func(http.ResponseWriter, *http.Request){"/": routes.HelloHandler})

	//Start
	webService.Logger.Sugar().Infof("Start service at port: %v", webService.Config.Port)
	webService.Logger.Error("Error")
	err = webService.Start()
	if err != nil {
		log.Fatalf("WebService Fatal error: %v\n", err)
	}
}
