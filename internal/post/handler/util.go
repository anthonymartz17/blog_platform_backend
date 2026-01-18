package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter,statusCode int, data any) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(statusCode)
	if err:= json.NewEncoder(w).Encode(data); err != nil{
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
	}

}

func ResponseError(w http.ResponseWriter,statusCode int,msg string, err error){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(statusCode)
	_= json.NewEncoder(w).Encode(map[string]string{
		"error":fmt.Sprintf("%s %v: ",msg,err),
	})
}