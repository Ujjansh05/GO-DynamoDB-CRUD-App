package product

import (
	"_/C_/Users/ujjan/Music/Go/GO_DynamoDB_CRUD_App/internal/handlers/product"
	"net/http"
)

type Handler struct{
	handler.Interface
	Controller product.Interface
	Rules Rules.Interface
}

func NewHandler(respository adapter.Interface) handlers.Interface {
	return &Handler{
		Controller: product.NewController(respository),
		Rules : RulesProduct.NewRules(),
	}
}


func Get(){

}

func getOne() {

}

func getAll() {

}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	productBody, err := h.getBodyAndValidate(r, uuid.Nil)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return 
	}

	ID, err := h.Controller.Create(productBody)

	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}
	HttpStatus.Status(w, r,map[string]interface{}{"id" : ID.String()})
}

func Put(){

}

func Delete(){

}

func Options(){

}

func