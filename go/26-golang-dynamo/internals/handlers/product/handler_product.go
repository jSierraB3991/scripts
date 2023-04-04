package product

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	Controller "github.com/jsierrab3991/scripts/26-golang-dynamo/internals/controllers/product"
	EntityProduct "github.com/jsierrab3991/scripts/26-golang-dynamo/internals/entities/product"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/handlers"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/repositories/adapter"
	"github.com/jsierrab3991/scripts/26-golang-dynamo/internals/rules"
	RulesProduct "github.com/jsierrab3991/scripts/26-golang-dynamo/internals/rules/product"
	response "github.com/jsierrab3991/scripts/26-golang-dynamo/utils/http"
)

type Handler struct {
	handlers.Interface
	Controller Controller.Interface
	Rules      rules.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Controller: Controller.NewController(repository),
		Rules:      RulesProduct.NewRules(),
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if chi.URLParam(r, "id") != "" {
		h.GetOne(w, r)
	} else {
		h.GetAll(w, r)
	}
}

func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		response.StatusBadRequest(w, r, errors.New("Id is not uuid valid"))
		return
	}
	data, err := h.Controller.ListOne(id)
	if err != nil {
		response.StatusInternalServerError(w, r, err)
		return
	}
	response.StatusOk(w, r, data)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.Controller.ListAll()
	if err != nil {
		response.StatusInternalServerError(w, r, err)
		return
	}
	response.StatusOk(w, r, data)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	productBody, err := h.getBodyAndValidate(r, uuid.Nil)
	if err != nil {
		response.StatusBadRequest(w, r, err)
		return
	}
	ID, err := h.Controller.Create(productBody)
	if err != nil {
		response.StatusInternalServerError(w, r, err)
		return
	}
	response.StatusOk(w, r, map[string]interface{}{"id": ID.String()})
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		response.StatusBadRequest(w, r, errors.New("Id is not uuid valid"))
		return
	}
	productBody, err := h.getBodyAndValidate(r, id)
	if err != nil {
		response.StatusInternalServerError(w, r, err)
		return
	}
	err = h.Controller.Update(id, productBody)
	if err != nil {
		response.StatusInternalServerError(w, r, err)
		return
	}
	response.StatusNoContent(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		response.StatusBadRequest(w, r, errors.New("Id is not uuid valid"))
		return
	}
	err = h.Controller.Remove(id)
	if err != nil {
		response.StatusInternalServerError(w, r, err)
		return
	}
	response.StatusNoContent(w, r)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	response.StatusNoContent(w, r)
}

func (h *Handler) getBodyAndValidate(r *http.Request, id uuid.UUID) (EntityProduct.Product, error) {
	product := EntityProduct.ProductBody{}
	body, err := h.Rules.ConvertIoReaderToStruct(r.Body, product)
	if err != nil {
		return EntityProduct.Product{}, errors.New("Body is required")
	}
	productParsed, err := EntityProduct.InterfaceToModel(body)
	if err != nil {
		return EntityProduct.Product{}, errors.New("")
	}
	setDefaultValues(productParsed, id)
	return productParsed, h.Rules.Validate(productParsed)
}

func setDefaultValues(product EntityProduct.Product, id uuid.UUID) {
	product.UpdateAt = time.Now()
	if id == uuid.Nil {
		product.Id = uuid.New()
		product.CreateAt = time.Now()
	}
}
