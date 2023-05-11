package main

import (
	app "contentPublicationBACK"
	"contentPublicationBACK/pkg/handler"
	"contentPublicationBACK/pkg/repository"
	"contentPublicationBACK/pkg/service"
	"log"
)

func main() {
	cnf := repository.Config{
		Host:       "localhost",
		DriverName: "mysql",
		Port:       "3306",
		DBName:     "avia",
		Username:   "root",
		Password:   "1111",
		ParseTime:  "true",
	}
	repository.CreateORMModels(cnf)
	db, err := repository.NewDB(cnf)

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(app.Server)
	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while starting server: %s", err.Error())
	}
}
