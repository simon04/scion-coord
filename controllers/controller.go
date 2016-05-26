// This is the main interface file which defines the methods of the controller
package controllers

import (
	"encoding/json"
	"net/http"
)

type output interface {
	// returns JSON data
	JSON(data interface{}, w http.ResponseWriter, r *http.Request)
	// Response with a 500 and an error
	Error500(err error, w http.ResponseWriter, r *http.Request)

	BadRequest(err error, w http.ResponseWriter, r *http.Request)

	NotFound(err error, w http.ResponseWriter, r *http.Request)

	Forbidden(err error, w http.ResponseWriter, r *http.Request)
}

type HTTPController struct {
	output
}

func (c HTTPController) JSON(data interface{}, w http.ResponseWriter, r *http.Request) {

	var content []byte
	var err error

	// set the JSON header
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	//marshall the data into json bytes
	content, err = json.Marshal(data)

	// in case of marshalling error
	// return 500
	if err != nil {
		c.Error500(err, w, r)
		return
	}

	// write the content to socket
	if _, err := w.Write(content); err != nil {
		c.Error500(err, w, r)
	}

}

func (c HTTPController) Error500(err error, w http.ResponseWriter, r *http.Request) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func (c HTTPController) BadRequest(err error, w http.ResponseWriter, r *http.Request) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func (c HTTPController) NotFound(err string, w http.ResponseWriter, r *http.Request) {
	http.Error(w, err, http.StatusNotFound)
}

func (c HTTPController) Forbidden(err error, w http.ResponseWriter, r *http.Request) {
	http.Error(w, err.Error(), http.StatusNotFound)
}
