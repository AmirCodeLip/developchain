package http_handlers

import (
	"dc_middleware/ipfs_connector"
	"encoding/json"
	"fmt"
	"net/http"
)

func isMultipartFormData(header http.Header) bool {
	contentType := header["Content-Type"]
	return contentType != nil && len(contentType[0]) >= 19 && contentType[0][0:19] == "multipart/form-data"
}

func isMultipartJson(header http.Header) bool {
	contentType := header["Content-Type"]
	return contentType != nil && len(contentType[0]) >= 16 && contentType[0][0:16] == "application/json"
}

func GetHandlers(connector ipfs_connector.RpcConnector) (Handler, *HandlerData) {
	var fileHashList []string
	data := HandlerData{connector, fileHashList}
	var handler Handler
	handler = &data
	return handler, &data
}

func (hd *HandlerData) Upload() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if isMultipartFormData(r.Header) {
			r.ParseMultipartForm(0)
			file := r.MultipartForm.File["file"][0]
			data := hd.connector.Add(file, false)
			hd.FileHashList = append(hd.FileHashList, data.Hash)
			fmt.Println(hd.FileHashList)
			fmt.Println("file with hash %s is add", data.Hash)
			b, _ := json.Marshal(data)
			fmt.Fprintf(w, "%s", string(b))
		}
	}
}
