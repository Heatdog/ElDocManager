package transport

import (
	"ElDocManager/pkg/logging"
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
	logger := logging.GetLogger()
	logger.Info("login post action")
	fmt.Println("POST action")
	var body []byte
	body, _ = io.ReadAll(r.Body)
	r.Body.Close()

	var user LoginForm
	json.Unmarshal(body, &user)
	fmt.Println(user)
	token := "1"
	resp, _ := json.Marshal(map[string]interface{}{
		"token": token,
	})
	w.Write(resp)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+token)
	logger.Infof("Succesfull authorization user: %s", user.Login)
}
