package main

import (
	"log"
	"net/http"
	"io"
	"encoding/json"
	"strconv"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func signup(w http.ResponseWriter, r *http.Request) {
	log.Println("exec signup function")

	length, err := strconv.Atoi(r.Header.Get("Content-Length"))

	body := make([]byte, length)
  length, err = r.Body.Read(body)
  if err != nil && err != io.EOF {
    w.WriteHeader(http.StatusInternalServerError)
    return
	}

	var jsonBody map[string]interface{}

  err = json.Unmarshal(body[:length], &jsonBody)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
  log.Printf("%v\n", jsonBody)
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("exec login function")
	log.Println(r)
	w.Write([]byte("successfully called login"))
}

func QueryStringHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	log.Println(w, "%s Loves Gorilla\n", q.Get("name"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/signup", signup).Methods("POST")
	r.HandleFunc("/signup", QueryStringHandler)
	r.HandleFunc("/login", login).Methods("POST")

	log.Println("start server : 3002 port")
	log.Fatal(http.ListenAndServe(":3002",
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
									handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
									handlers.AllowedOrigins([]string{"*"}))(r)))
	http.Handle("/", r)
}
