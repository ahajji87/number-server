package main

import (
	"flag"
	"log"
	"number-server/app/configuration"
	"number-server/app/domain/model"
	"number-server/app/domain/service"
	"number-server/app/infrastructure/server"
	"number-server/app/infrastructure/storage"
	"number-server/app/interface/handler"
	"number-server/app/interface/repository/file"
	"number-server/app/usecase"
	"time"
)

func main() {
	configFilename := flag.String("config", "config.yml", "Path to configuration file")
	flag.Parse()

	config := configuration.Load(*configFilename)

	storage := storage.NewFileStorage(config.Storage.Path)

	if err := storage.Init(); err != nil {
		return
	}

	var pipeline = make(chan *model.Number)

	numberRepository := file.NewFileNumberRepository(storage)
	numberService := service.NewNumberService(numberRepository)
	numberUseCase := usecase.NewNumberUseCase(numberService, pipeline)

	numberHandler := handler.NewNumberHandler(numberUseCase)

	//Start the server
	server := server.NewServer(numberHandler, config.Server)

	done := make(chan bool)

	go func() {
		for {
			number := <-pipeline
			if err := numberUseCase.Store(number); err != nil {
				break
			}
		}
	}()

	go func() {
		for t := range time.Tick(time.Duration(config.App.Report.Time) * time.Second) {
			report := numberUseCase.GetReport()

			log.Printf("T%0.2d:%0.2d:%0.2d >> %s", t.Hour(), t.Minute(), t.Second(), report)
		}
	}()

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Listen and serve: %v", err)
		}
		done <- true
	}()

	//wait shutdown
	server.WaitShutdown()

	<-done
	log.Printf("DONE!")
}
