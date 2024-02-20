package transport

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LoginForm struct {
	Login    string
	Password string
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST action")
	var body []byte
	body, _ = io.ReadAll(r.Body)
	r.Body.Close()

	var user LoginForm
	json.Unmarshal(body, &user)
	fmt.Println(user)
	resp, _ := json.Marshal(map[string]interface{}{
		"token": "1",
	})
	w.Write(resp)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
}
