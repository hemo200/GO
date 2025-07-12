package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Creating application logger struct

type application struct {
	logger *slog.Logger
}

func main() {
	//This variable is to define the default port of the application
	addr := flag.String("addr", ":4000", "HTTP network address")

	//Define a newe connection to DB
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=True", "MySQL data source name")

	//To read the flag value from the command line
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	//Creation of a logger

	//Declare an application logger method
	app := &application{logger: logger}
	defer db.Close()
	//To display all application info logs
	logger.Info("Starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())

	//to create error logs when something goes wrong
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
