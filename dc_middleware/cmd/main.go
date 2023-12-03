package main

import (
	"dc_middleware/http_handlers"
	"dc_middleware/ipfs_connector"
	"log"
	"net/http"
)

func main() {
	connector := ipfs_connector.NewRpc("http://168.119.58.0:5001")
	handlers := http_handlers.GetHandlers(connector)
	http.HandleFunc("/upload", handlers.Upload())
	http.HandleFunc("/pin", handlers.Pin())
	log.Fatal(http.ListenAndServe(":8090", nil))
}
