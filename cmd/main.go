package main

import (
	"balance/internal/db"
	"balance/internal/config"
	"balance/internal/handler"
	//"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"log"
	"fmt"
)

func main() {
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello!")
    })

	db, err := db.ConnectDB(&c)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Printf("Starting server at port 8080")

    http.HandleFunc("/balance", handler.GetBalance)

	// router := mux.NewRouter()
	// //router.HandleFunc("/add", handler.AddMoney).Methods("POST")
	// router.HandleFunc("/balance", handler.GetBalance).Methods("GET")
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}