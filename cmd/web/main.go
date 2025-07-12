package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Creating application logger struct

type application struct {
	logger *slog.Logger
}

func main() {
	//This variable is to define the default port of the application
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse() //To read values from runtime

	//Creation of a logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//Declare an application logger method
	app := &application{logger: logger}

	//To display all application info logs
	logger.Info("Starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())

	//to create error logs when something goes wrong
	logger.Error(err.Error())
	os.Exit(1)
}
