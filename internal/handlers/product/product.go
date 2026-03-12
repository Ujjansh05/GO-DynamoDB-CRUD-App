package product

import (
	"errors"
	"net/http"
	"time"

	ControllerProduct "github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/controllers/product"
	EntityProduct "github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/entities/product"
	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/handlers"
	"github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/repository/adapter"
	Rules "github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/rules"
	RulesProduct "github.com/Ujjansh05/GO_Dynamo_CRUD_App/internal/rules/product"
	HttpStatus "github.com/Ujjansh05/GO_Dynamo_CRUD_App/utils/http"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type Handler struct {
	handlers.Interface

	Controller ControllerProduct.Interface
	Rules      Rules.Interface
}

func NewHandler(respository adapter.Interface) handlers.Interface {
	return &Handler{
		Controller: ControllerProduct.NewController(respository),
		Rules:      RulesProduct.NewRules(),
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if chi.URLParam(r, "ID") != "" {
		h.getOne(w, r)
		return
	}

	h.getAll(w, r)
}

func (h *Handler) getOne(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("Id is not valid uuid"))
		return
	}

	response, err := h.Controller.ListOne(id)
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, response)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	response, err := h.Controller.ListAll()
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, response)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	productBody, err := h.getBodyAndValidate(r, uuid.Nil)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}

	id, err := h.Controller.Create(productBody)
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.Status(w, r, map[string]interface{}{"id": id.String()})
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}

	productBody, err := h.getBodyAndValidate(r, id)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}

	if err := h.Controller.Update(id, productBody); err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusNoContent(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("Id is not uuid valid"))
		return
	}

	if err := h.Controller.Remove(id); err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusNoContent(w, r)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}

func (h *Handler) getBodyAndValidate(r *http.Request, id uuid.UUID) (*EntityProduct.Product, error) {
	productBody := &EntityProduct.Product{}
	body, err := h.Rules.ConvertIoReaderToStruct(r.Body, productBody)
	if err != nil {
		return &EntityProduct.Product{}, errors.New("body is required")
	}

	productParsed, err := EntityProduct.InterfaceToModel(body)
	if err != nil {
		return &EntityProduct.Product{}, errors.New("error on converting body")
	}

	setDefaultValues(productParsed, id)
	return productParsed, h.Rules.Validate(productParsed)
}

func setDefaultValues(product *EntityProduct.Product, id uuid.UUID) {
	product.UpdatedAt = time.Now()
	if id == uuid.Nil {
		product.ID = uuid.New()
		product.CreatedAt = time.Now()
		return
	}

	product.ID = id
}
