package handlers

import (
	"io/ioutil"
	"net/http"
)

// HandlerIndex ...
type HandlerIndex struct {

}

// NewHandlerIndex ...
func NewHandlerIndex() *HandlerIndex {
	return &HandlerIndex{}
}

// HandleIndex ...
func (i HandlerIndex ) HandleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		data, err := ioutil.ReadFile("templates/index.html")
		if err == nil {
			w.Write(data)
		}

		w.WriteHeader(http.StatusOK)
		//w.Write([]byte("☄ HTTP status code returned!"))
	}
}
