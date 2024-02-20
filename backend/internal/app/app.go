package app

import (
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

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}

}
