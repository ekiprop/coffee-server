package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ekiprop/coffee-app/coffee-app/coffee-server/db"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

func (app *Application) Serve() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error leading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println("API is lietening on port", port)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}
	return srv.ListenAndServe()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error leading .env file")
	}

	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DSN")
	dbConn, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Error connecting to Database")
	}

	defer dbConn.Close()

	app := &Application{
		Config: cfg,
		//TODO: add models later

	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
