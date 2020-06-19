package handlers

import "net/http"

// Handlers ...
type Handlers struct {
	Index *HandlerIndex
	Sender  *HandlerSender
}

// NewHandlers ...
func NewHandlers() *Handlers {

	return &Handlers{
		Index: NewHandlerIndex(),
		Sender:  NewHandlerSender(),
	}
}

// enableCors ...
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}