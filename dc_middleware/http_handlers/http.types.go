package http_handlers

import "net/http"

type Handler interface {
	Upload() func(w http.ResponseWriter, r *http.Request)
}
