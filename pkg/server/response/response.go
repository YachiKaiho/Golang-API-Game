package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Success Process (http code:200)
func Success(writer http.ResponseWriter, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		InternalServerError(writer, "marshal error")
		return
	}
	writer.Write(data)
}

//BadRequest process(http code:400)
func BadRequest(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusInternalServerError, message)
}

//InternalServerError process (http code:500)
func InternalServerError(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusInternalServerError, message)
}

//response output for error
func httpError(writer http.ResponseWriter, code int, message string) {
	data, _ := json.Marshal(errorResponse{
		Code:    code,
		Message: message,
	})
	writer.WriteHeader(code)
	if data != nil {
		writer.Write(data)
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
