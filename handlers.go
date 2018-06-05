package main

import (
	"net/http"
)

/*
PingHandler Pong!
curl -X GET http://localhost:3000/ping
*/
func PingHandler(writer http.ResponseWriter, request *http.Request) {
	//Pong
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Pong!"))
}
