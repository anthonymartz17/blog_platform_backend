package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

func Error(w http.ResponseWriter, statusCode int, msg string) {
	JSON(w, statusCode, map[string]string{"error": msg})
}