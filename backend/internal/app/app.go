package app

import (
	"ElDocManager/internal/config"
	"ElDocManager/internal/transport"
	"ElDocManager/pkg/logging"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	logger := logging.GetLogger()

	logger.Info("create router")
	router := mux.NewRouter()

	routerSub := router.PathPrefix("/api").Subrouter()
	logger.Info("register login endpoint")
	routerSub.HandleFunc("/login", transport.SignInHandler).Methods(http.MethodPost)
	logger.Info("register registration endpoint")
	routerSub.HandleFunc("/register", transport.SignUpHandler).Methods(http.MethodPost)

	cors, err := config.CorsSettings()
	if err != nil {
		log.Fatalf(err.Error())
	}

	handler := cors.Handler(router)

	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}

}
