package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup関数実行")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login関数実行")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/singup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	log.Println("サーバー起動 : 8000 port で受信")
	log.Fatal(http.ListenAndServe(":8000", router))
}
