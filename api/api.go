package api

import (
	"encoding/json"
	"net/http"
)

//coin balance params 
type CoinBalanceParams struct {
	Username string
}

//coin balance response
type CoinBalanceResponse struct {
	Code int
	Balance int64
}

type Error struct{
	Code int
	Message string
} 

func writeError(w http.ResponseWriter, meassage string, code int){
	resp := Error{
		Code: code,
		Message: meassage,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)


}


var (
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)