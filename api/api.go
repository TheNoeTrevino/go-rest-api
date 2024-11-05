package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	// success code
	Code int16
	// account balance
	Balance int64
}

type Error struct {
	Message string
	Code    int
}

func writeError(writer http.ResponseWriter, message string, code int) {
	response := Error{
		Code:    code,
		Message: message,
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)

	json.NewEncoder(writer).Encode(response)
}

// this is just an anon function!
var (
	RequestErrorHandler = func(writer http.ResponseWriter, err error) {
		writeError(writer, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(writer http.ResponseWriter) {
		writeError(writer, "An unexpected error ocurred", http.StatusInternalServerError)
	}
)
