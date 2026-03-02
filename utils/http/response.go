package http

import (
	"encoding/json"
	"log"
	"net/http"
)


type response struct{
	Status int `json:"result"`
	Result interface{} `json : "result"`
}

func newResponse() *response{

}

func (resp *response) bytes() []byte{

}

func (resp *response) string() [] byte{

}

func (resp *response) sendResponse(w http.ResponseWriter, r *http.Request){

}

//200
func StatusNoContent()

//400
func StatusBadRequest()

//404
func StatusNotFound()

//405
func  StatusMethodNotAllowed(){

}

//409
func StatusConflict(){

}

//500
func StatusInternalServerError(){
	
}