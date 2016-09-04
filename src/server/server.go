package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
)


func Log(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func Run(port int) {

	rand.Seed(time.Now().UTC().UnixNano())
	r := mux.NewRouter()

	r.HandleFunc("/phrase", getPhraseFromNum).Methods("GET") 
	r.HandleFunc("/number", getNumFromPhrase).Methods("GET") 

	for {
		log.Printf("Running at 0.0.0.0:%d\n", port)
		log.Println(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), r))
		time.Sleep(1 * time.Second)
	}
}
