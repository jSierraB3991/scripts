package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func newResponse(data interface{}, status int) *Response {
	return &Response{
		Status: status,
		Result: data,
	}
}

func (resp *Response) bytes() []byte {
	data, _ := json.Marshal(resp)
	return data
}

func (resp *Response) string() string {
	return string(resp.bytes())
}

func (resp *Response) sendResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(resp.Status)
	_, _ = w.Write(resp.bytes())
	log.Print(resp.string())
}

func StatusOk(w http.ResponseWriter, r *http.Request, data interface{}) {
	newResponse(data, http.StatusOK).sendResponse(w, r)
}

func StatusNoContent(w http.ResponseWriter, r *http.Request) {
	newResponse("", http.StatusNoContent).sendResponse(w, r)
}

func StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusBadRequest).sendResponse(w, r)
}

func StatusNotFound(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusNotFound).sendResponse(w, r)
}

func StatusMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"error": "Method Not Allowed"}
	newResponse(data, http.StatusMethodNotAllowed).sendResponse(w, r)
}

func StatusConflict(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusConflict).sendResponse(w, r)
}

func StatusInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusInternalServerError).sendResponse(w, r)
}
