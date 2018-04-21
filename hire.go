package main

import (
	"fmt"
	"log"
	"net/http"

	hcontext "github.com/smithaitufe/hire/hire/context"
)

func main() {
	config := hcontext.LoadConfig()
	db, err := hcontext.OpenDB(config)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer db.Close()
	hcontext.Migrate(db)
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello World")
	}))

	log.Fatal(http.ListenAndServe("5000"))
}
