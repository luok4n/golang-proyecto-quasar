package commons

import (
	"net/http"
)

func SendResponse(writer http.ResponseWriter, status int, data []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(data)
}

func SendError(writer http.ResponseWriter, status int) {
	data := []byte(`{}`)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(data)
}
