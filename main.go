package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/arschles/kubehttpbin/handlers"
	"github.com/gorilla/mux"
)

const (
	defaultPort = 8080
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = defaultPort
	}

	router := mux.NewRouter()
	router.HandleFunc("/ip", handlers.IP).Methods("GET")
	router.HandleFunc("/get", handlers.Get).Methods("GET")
	router.HandleFunc("/post", handlers.Post).Methods("POST")
	router.HandleFunc("/put", handlers.Put).Methods("PUT")
	router.HandleFunc("/delete", handlers.Delete).Methods("DELETE")
	router.HandleFunc("/head", handlers.Head).Methods("HEAD")
	router.HandleFunc("/patch", handlers.Patch).Methods("PATCH")
	router.HandleFunc("/headers", handlers.Headers).Methods("GET")

	hostStr := fmt.Sprintf(":%d", port)
	log.Printf("kubehttpbin listening on %s", hostStr)
	http.ListenAndServe(hostStr, router)

}
