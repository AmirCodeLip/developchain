package main

import (
	"dc_middleware/http_handlers"
	"dc_middleware/ipfs_connector"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	connector := ipfs_connector.NewRpc("http://168.119.58.0:5001")
	api := http_handlers.GetHandlers(connector)
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	r.HandleFunc("/upload", api.Upload())
	r.HandleFunc("/pin", api.Pin())
	http.Handle("/", r)

	http.ListenAndServe(":8090", r)

	log.Fatal(http.ListenAndServe(":8090", handlers.CORS(header, methods, origins)(r)))
	//if err != nil {
	//	fmt.Println(err)
	//}
}
