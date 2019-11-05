package main

import (
	"fmt"
	"github.com/PharbersDeveloper/es-sql-pods/handler"
	"log"
	"net/http"
)

const EsServer string = "http://192.168.100.174:9200"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/sql", handler.SqlHandler)

	port := "3000"

	log.Println("Listening...", port)
	http.ListenAndServe(fmt.Sprint(":", port), mux)
}

