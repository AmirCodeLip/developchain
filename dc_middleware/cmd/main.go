package main

import (
	"dc_middleware/http_handlers"
	"dc_middleware/ipfs_connector"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func checkFileUpload(apiData *http_handlers.HandlerData) {
	//contractLogic := x7opsilon_client.NewContractLogic("http://168.119.58.0:8545")
	for {
		time.Sleep(10000 * time.Millisecond)
		//fmt.Println(apiData.FileHashList)
	}
}

func main() {
	r := mux.NewRouter()
	connector := ipfs_connector.NewRpc("http://168.119.58.0:5001")
	api, apiData := http_handlers.GetHandlers(connector)
	go checkFileUpload(apiData)
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	r.HandleFunc("/upload", api.Upload())
	r.HandleFunc("/pin", api.Pin())
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8090", handlers.CORS(header, methods, origins)(r)))
}
