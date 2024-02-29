package transport

import (
	"ElDocManager/internal/user"
	"ElDocManager/pkg/logging"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type userHandler struct {
	logger  *logging.Logger
	service user.UserService
}

func NewUserHandler(logger *logging.Logger, service user.UserService) Handler {
	return &userHandler{
		logger:  logger,
		service: service,
	}
}

const (
	signInURL  = "/login"
	signUpURL  = "/register"
	signOutURL = "/quite"
)

func (h *userHandler) Register(router *mux.Router) {
	router.HandleFunc(signInURL, h.signIn).Methods(http.MethodPost)
	router.HandleFunc(signUpURL, h.signUp).Methods(http.MethodPost)
	router.HandleFunc(signOutURL, h.signOut).Methods(http.MethodPost)
}

func (h *userHandler) signIn(w http.ResponseWriter, r *http.Request) {
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

	token := h.service.SignIn(r.Context(), &userSignIn)

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

func (h *userHandler) signUp(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("sign up post action")
	var body []byte
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Infof("bad request: %s", err.Error())
		return
	}
	defer r.Body.Close()

	userSignUp := &user.UserSignUp{}
	if err := json.Unmarshal(body, userSignUp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.logger.Infof("bad request: %s", err.Error())
		return
	}

	if err := h.service.SignUp(r.Context(), userSignUp); err != nil {
		//  посмтреть разные коды ошибок, т.к. може быть неуникальные значения
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Infof("error happened: %s", err.Error())
		return
	}
	// сервис работы с запросами на подключение к системе

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "aplication/json")
	h.logger.Infof("successful authorize")
}

func (h *userHandler) signOut(w http.ResponseWriter, r *http.Request) {

}
