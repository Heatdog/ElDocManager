package transport

import "net/http"

func RootHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("root url"))
}
