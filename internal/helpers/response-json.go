package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)

}

func ResponseError(w http.ResponseWriter, statusCode int, message string) {
	fmt.Println(message)
	ResponseJson(w, statusCode, map[string]string{"error": message})
}
