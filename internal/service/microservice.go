package service

//Каркас микросервиса
//Реализовать "каркас" микросервиса, считывающий конфиг из файла,
//создающий логгер/логгеры с указанными уровнями детализации.

import (
	serviceConfig "github.com/iae94/Microservice/internal/config"
	serviceLogger "github.com/iae94/Microservice/internal/logger"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type Microservice interface {
	Init(config *serviceConfig.Config) error
	RegisterHandlers(handlers *map[string]func(http.ResponseWriter, *http.Request))
	Start() error
}

type WebService struct {
	Config   *serviceConfig.Config
	Logger   *zap.Logger
	Handlers map[string]func(http.ResponseWriter, *http.Request)
}

func (s *WebService) Init(config *serviceConfig.Config) error {

	logger, err := serviceLogger.CreateLogger(config)
	if err != nil {
		log.Printf("Create logger error: %v\n", err)
		return err
	}

	s.Config = config
	s.Logger = logger

	return nil
}

func (s *WebService) RegisterHandlers(handlers *map[string]func(http.ResponseWriter, *http.Request)){
	sugar := s.Logger.Sugar()
	for url, handler := range *handlers {
		http.HandleFunc(url, handler)
		sugar.Infof("Register handler: %v", url)
	}
}

func (s *WebService) Start() error {
	defer s.Logger.Sync()
	sugar := s.Logger.Sugar()
	err := http.ListenAndServe(fmt.Sprintf(":%v", s.Config.Port), nil)
	if err != nil {
		sugar.Errorf("Listener error: %v\n", err)
	}
	return err
}
