package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/samoei/coffeeapp/db"
)

type Config struct {
	Port string
}
type Application struct {
	Config Config
}

func (app Application) serve() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading the .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println("API server is listening on port:", port)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		//TODO: add router
	}

	return srv.ListenAndServe()

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading the .env file")
	}

	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DSN")
	dbConn, err := db.ConnectPostgres(dsn)

	if err != nil {
		log.Fatal("Can not connect to the database")
	}

	defer dbConn.DB.Close()

	app := &Application{
		Config: cfg,
		//TODO: add models later
	}

	err = app.serve()

	if err != nil {
		log.Fatal(err)
	}
}
