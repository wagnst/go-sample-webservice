package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Configuration struct {
	Port string
}

const ConfigFile = "config.json"

func main() {
	/* initialize config file */
	file, _ := os.Open(ConfigFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatalf("Could not decode config file: %s", err)
	}
	/* initialize router */
	router := mux.NewRouter()

	router.HandleFunc("/platform", GetPlatformsJson).Methods("GET")
	router.HandleFunc("/platform/{id}", GetPlatformJson).Methods("GET")

	log.Printf("server listening on tcp/%s initialized", configuration.Port)
	log.Fatal(http.ListenAndServe(":"+configuration.Port, router))
}
