package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/andreicalinciuc/mock-api/service/router"
	"github.com/andreicalinciuc/mock-api/transport/http/handler"
	"github.com/gorilla/mux"
	"github.com/oklog/run"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	var httpServerAddress string

	flag.StringVar(&httpServerAddress, "addr", ":3380", "The http listen address")
	flag.Parse()

	r := mux.NewRouter()
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	// Internal API
	api := r.PathPrefix("/data").Subrouter()
	// Middleware goes here

	muxRouter := router.NewMuxRouter(api, log)
	handler.NewUser(muxRouter, log)

	httpd := &http.Server{
		Addr:           httpServerAddress,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        api,
	}

	var g run.Group

	ctx, _ := context.WithCancel(context.Background())
	g.Add(
		func() error {
			return httpd.ListenAndServe()
		},
		func(error) {
			httpd.Shutdown(ctx)
		},
	)

	fmt.Println("RUN")
	g.Run()

}
