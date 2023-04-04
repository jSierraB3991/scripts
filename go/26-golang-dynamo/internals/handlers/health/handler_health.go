package health

import (
	"errors"
	"net/http"

	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/handlers"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/repositories/adapter"
	response "github.com/jsierrab3991/scripts/26-golang-dynamo/utils/http"
)

type Handler struct {
	handlers.Interface
	repository adapter.Interface
}

func NewHandler(adapterInterface adapter.Interface) *Handler {
	return &Handler{
		repository: adapterInterface,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if !h.repository.Health() {
		response.StatusInternalServerError(w, r, errors.New("Relational database not alive"))
		return
	}

	response.StatusOk(w, r, "Service Ok")
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	response.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	response.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	response.StatusMethodNotAllowed(w, r)

}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	response.StatusMethodNotAllowed(w, r)

}
