package main

import (
	"Project-Workey/config"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/subosito/gotenv"
)

// to call functions before main functions
func init() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {

	//Loading value from env file
	port := os.Getenv("PORT")
	// port := os.Getenv("PORT")
	fmt.Println("Workey job-portal projects beggined")

	// creating an instance of chi r
	router := chi.NewRouter()
	config.Init()

	var (
		db *sql.DB = config.ConnectDB()
	)
	fmt.Println(db)

	log.Println("Api is listening on port:", port)
	// Starting server
	http.ListenAndServe(":"+port, router)
}
