package http_handlers

import (
	"dc_middleware/ipfs_connector"
	"encoding/json"
	"fmt"
	"io"
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

func GetHandlers(connector ipfs_connector.RpcConnector) Handler {
	data := HandlerData{connector}
	var handler Handler
	handler = &data
	return handler
}

func (hd *HandlerData) Upload() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if isMultipartFormData(r.Header) {
			r.ParseMultipartForm(0)
			file := r.MultipartForm.File["file"][0]
			data := hd.connector.Add(file, false)
			b, _ := json.Marshal(data)
			fmt.Fprintf(w, "%s", string(b))
		}
	}
}

func (hd *HandlerData) Pin() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if isMultipartJson(r.Header) {
			strData, _ := io.ReadAll(r.Body)
			fmt.Println(string(strData))
		}
	}
}
