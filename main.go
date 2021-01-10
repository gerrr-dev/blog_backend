package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Doctrina`)
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
//DB connecting
	
	HOST, _        := os.LookupEnv("HOST")
	DB_PORT, _     := os.LookupEnv("DB_PORT")
	USER, _        := os.LookupEnv("PROFILE")
	PASS, _        := os.LookupEnv("PASSWORD")
	DB_NAME, _     := os.LookupEnv("DB_NAME")
	
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=require",
		HOST, DB_PORT, USER, PASS, DB_NAME)
	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	err = db.Ping()
	if err != nil {
		panic(err)
	}  
//---//
	
	
	
//Routes handler
	http.HandleFunc("/", hello)
	
	
	
//Starting server
	port, _ := os.LookupEnv("PORT")
	log.Print("Listening on :" + port) //simple log
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
