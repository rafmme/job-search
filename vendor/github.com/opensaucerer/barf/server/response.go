package server

import (
	"encoding/json"
	"net/http"

	"github.com/opensaucerer/barf/typing"
)

// JSON writes a JSON response to the response writer
func JSON(w http.ResponseWriter, status bool, statusCode int, message string, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(typing.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

type response struct {
	code   int
	body   interface{}
	writer http.ResponseWriter
}

// JSON writes a JSON response to the response writer
func (r *response) JSON(data interface{}) {
	r.body = data
	r.writer.Header().Set("Content-Type", "application/json")
	r.writer.WriteHeader(r.code)
	json.NewEncoder(r.writer).Encode(data)
}

// Status loads a barf response with the given status code
func (r *response) Status(code int) *response {
	r.code = code
	return r
}

// Response prepares a barf response with the given writer
func Response(w http.ResponseWriter) *response {
	return &response{
		writer: w,
	}
}
