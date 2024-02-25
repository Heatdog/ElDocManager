package transport

import (
	"ElDocManager/internal/user"
	"ElDocManager/pkg/logging"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type handlerAuth struct {
	logger  *logging.Logger
	service user.UserService
}

func NewHandlerAuth(logger *logging.Logger, service user.UserService) Handler {
	return &handlerAuth{
		logger:  logger,
		service: service,
	}
}

const (
	signInURL  = "/login"
	signUpURL  = "/register"
	signOutURL = "/quite"
)

func (h *handlerAuth) Register(router *mux.Router) {
	router.HandleFunc(signInURL, h.SignIn).Methods(http.MethodPost)
	router.HandleFunc(signUpURL, h.SignUp).Methods(http.MethodPost)
	router.HandleFunc(signOutURL, h.SugnOut).Methods(http.MethodPost)
}

func (h *handlerAuth) SignIn(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("sign in post action")
	var body []byte
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Infof("bad request: %s", err.Error())
		return
	}
	defer r.Body.Close()

	var userSignIn user.UserSignIn
	if err := json.Unmarshal(body, &userSignIn); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Infof("bad request: %s", err.Error())
		return
	}

	token := h.service.SignIn(context.Background(), &userSignIn)

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
	h.logger.Infof("Succesfull authorization user: %s", userSignIn.Login)
}

func (h *handlerAuth) SignUp(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("sign up post action")
	var body []byte
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Infof("bad request: %s", err.Error())
		return
	}
	defer r.Body.Close()

	var userSignUp user.UserSignUp
	if err := json.Unmarshal(body, &userSignUp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Infof("bad request: %s", err.Error())
		return
	}

	if err := h.service.SignUp(context.Background(), &userSignUp); err != nil {
		//  посмтреть разные коды ошибок, т.к. може быть неуникальные значения
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Infof("error happened: %s", err.Error())
		return
	}

	isSuccessful, err := json.Marshal(false)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Infof("parsing error: %s", err.Error())
		return
	}
	w.Write(isSuccessful)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "aplication/json")
	h.logger.Infof("successful authorize")
}

func (h *handlerAuth) SugnOut(w http.ResponseWriter, r *http.Request) {

}
