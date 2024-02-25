package auth

import (
	"ElDocManager/internal/transport"
	"ElDocManager/pkg/logging"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) transport.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/login", h.SignIn).Methods("POST")
	router.HandleFunc("/register", h.SignUp).Methods("POST")
	router.HandleFunc("/quite", h.SugnOut).Methods("POST")
}

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("login post action")
	fmt.Println("POST action")
	var body []byte
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Infof("bad request: %s", err.Error())
		return
	}
	defer r.Body.Close()

	var user UserLogin
	json.Unmarshal(body, &user)
	fmt.Println(user)
	token := "1"
	resp, err := json.Marshal(map[string]interface{}{
		"token": token,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Errorf("token generation error: %s", err.Error())
		return
	}
	w.Write(resp)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+token)
	h.logger.Infof("Succesfull authorization user: %s", user.Login)
}

func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) SugnOut(w http.ResponseWriter, r *http.Request) {

}
