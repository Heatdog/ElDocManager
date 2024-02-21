package app

import (
	"ElDocManager/internal/config"
	"ElDocManager/internal/transport"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()

	routerSub := router.PathPrefix("/api").Subrouter()
	routerSub.HandleFunc("/login", transport.SignInHandler).Methods(http.MethodPost)
	routerSub.HandleFunc("/register", transport.SignUpHandler).Methods(http.MethodPost)
	routerSub.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get action")
		w.Write([]byte("Тест"))
	}).Methods(http.MethodGet)

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
