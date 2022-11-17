package main

import (
	"fmt"
	"log"
	"net/http"

	"balance/internal/config"
	"balance/internal/controller"
	"balance/internal/controller/transport"
	"balance/internal/db"
)

func main() {
	c, err := config.GetConfig("/config")
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.New(&c)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctrl := controller.New(db)
	transport := transport.NewHTTP(ctrl)

	http.HandleFunc("/get-balance", transport.GetBalance)
	http.HandleFunc("/add-balance", transport.AddBalance)
	http.HandleFunc("/create-order", transport.CreateOrder)
	http.HandleFunc("/finish-order", transport.FinishOrder)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
