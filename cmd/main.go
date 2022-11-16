package main

import (
	"balance/internal/db"
	"balance/internal/config"
	"net/http"
	"log"
	"fmt"
)

func main() {
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

    // http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
    //     fmt.Fprintf(w, "Hello!")
    // })

	pool, err := db.CreatePool(&c)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}