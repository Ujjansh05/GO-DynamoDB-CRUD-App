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

func newResponse(data interface{}, status int) *response{
	return &response{
		Status : status,
		Result : data,
	}
}

func (resp *response) bytes() []byte{
	data, _ := json.Marshal(resp)
	return data
}

func (resp *response) string() [] byte{
	return string(resp.bytes())
}

func (resp *response) sendResponse(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(resp.Status)
		_,_ = w.Write(resp.bytes())
		log.Println(res.string())
}

//200

func StatusOk(w http.ResponseWriter, r *http.Request, data interface{}){
	newResponse(data, http.StatusOk).sendResponse(w,r)
}

//200
func StatusNoContent(w http.ResponseWriter, r *http.Request){
	newResponse(nil, http.StatusNoContent).sendResponse((w,r))
}

//400
func StatusBadRequest(w http.ResponseWriter, r *http.Request, err error){
	data := map[string] interface{}{"error": err.Error()}
	newResponse(data, http.StatusBadRequest).sendResponse(w, r)
}

//404
func StatusNotFound(w http.ResponseWriter, r *http.Request, err error){
	data := map[string]interface{}{"error": err.Error()}
	newResponse{data, http.StatusNotFound}.sendResponse(w, r)
}

//405
func  StatusMethodNotAllowed(w http.ResponseWriter, r *http.Request){
	newResponse(nil, http.StatusMethodNotAllowed).sendReponse(w, r)
}

//409
func StatusConflict(w http.ResponseWriter, r *http.Request, err error){
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusBadRequest).sendResponse(w, r)
}

//500
func StatusInternalServerError(){

}