package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/Rayzon3/go_project/internal/app"
	"github.com/Rayzon3/go_project/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Default server port")

	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("Application Running on port %d \n", port)

	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatal(err)
	}
}
