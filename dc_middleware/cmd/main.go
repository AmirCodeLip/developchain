package main

import (
	"dc_middleware/http_handlers"
	"dc_middleware/ipfs_connector"
	"dc_middleware/x7opsilon_client"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"math/big"
	"net/http"
	"time"
)

func checkFileUpload(apiData *http_handlers.HandlerData) {
	contractLogic := x7opsilon_client.NewContractLogic("http://168.119.58.0:8545")
	var lastBlock *big.Int = nil
	for {
		time.Sleep(10000 * time.Millisecond)
		blockNumber, eventResults := contractLogic.GetUploadedFiles(lastBlock)
		if blockNumber != 0 && eventResults != nil {
			lastBlock = new(big.Int).SetUint64(blockNumber)
			for _, eventResult := range eventResults {
				for hashIndex, fileHash := range apiData.FileHashList {
					if fileHash == eventResult.FileHash {
						apiData.Connector.Pin(eventResult.FileHash)
						apiData.FileHashList[hashIndex] = ""
						break
					}
				}

			}
			fmt.Println("last block is %s", lastBlock)
		}
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
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8090", handlers.CORS(header, methods, origins)(r)))
}
