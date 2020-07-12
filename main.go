package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup関数実行")
	w.Write([]byte("successfully called signup"))
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login関数実行")
	w.Write([]byte("successfully called login"))
}


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/singup", signup).Methods("POST")
	r.HandleFunc("/login", login).Methods("POST")

	log.Println("サーバー起動 : 8000 port で受信")
	log.Fatal(http.ListenAndServe(":8000",
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(r)))
	http.Handle("/", r)
}
