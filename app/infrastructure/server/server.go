package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"number-server/app"
	"number-server/app/interface/handler"

	"github.com/gorilla/mux"
)

type server struct {
	http.Server
	shutdownReq   chan bool
	reqCount      uint32
	configuration app.Server
}

func NewServer(h handler.NumberHandler, config app.Server) *server {
	//create server
	s := &server{
		Server: http.Server{
			Addr:         ":" + config.Port,
			ReadTimeout:  time.Duration(config.Timout.Read) * time.Second,
			WriteTimeout: time.Duration(config.Timout.Write) * time.Second,
		},
		shutdownReq:   make(chan bool),
		configuration: config,
	}

	router := mux.NewRouter()
	handler := http.HandlerFunc(h.Add)
	//register handlers
	router.Handle(config.Ws.Endpoint, maxClients(handler, config.Ws.Connections))
	router.HandleFunc(config.Shutdown.Endpoint, s.ShutdownHandler)

	//set http server handler
	s.Handler = router

	return s
}

func (s *server) WaitShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)

	//Wait interrupt or shutdown request through /shutdown
	select {
	case sig := <-irqSig:
		log.Printf("Shutdown request (signal: %v)", sig)
	case sig := <-s.shutdownReq:
		log.Printf("Shutdown request (/shutdown %v)", sig)
	}

	log.Printf("Stoping http server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.configuration.Shutdown.Timeout)*time.Second)
	defer cancel()

	//shutdown the server
	err := s.Shutdown(ctx)
	if err != nil {
		log.Printf("Shutdown request error: %v", err)
	}
}

func (s *server) ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shutdown server"))

	//Do nothing if shutdown request already issued
	//if s.reqCount == 0 then set to 1, return true otherwise false
	if !atomic.CompareAndSwapUint32(&s.reqCount, 0, 1) {
		log.Printf("Shutdown through API call in progress...")
		return
	}

	go func() {
		s.shutdownReq <- true
	}()
}

func maxClients(h http.Handler, n int) http.Handler {
	var concurrentConn = make(chan struct{}, n)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		concurrentConn <- struct{}{}
		defer func() {
			<-concurrentConn
		}()

		h.ServeHTTP(w, r)
	})
}
