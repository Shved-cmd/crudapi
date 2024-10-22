package main

import (
	"go.mod/pkg/router"
	"log"
	"net/http"
)

// CrudAPI

func main() {
	r := router.SetupeRouter()
	log.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
