package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Singtodev/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/", r)
	fmt.Println("Server is starting... port 9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}
