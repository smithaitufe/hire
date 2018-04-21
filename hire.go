package main

import (
	"fmt"
	"log"
	"net/http"

	hcontext "github.com/smithaitufe/hire/hire/context"
)

func main() {
	config := hcontext.LoadConfig()
	db := hcontext.OpenDB(config)
	defer db.Close()
	hcontext.Migrate(db)
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello World")
	}))
	port := fmt.Sprintf(":%s", config.Server.Port)
	fmt.Println("Server is listening at port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
