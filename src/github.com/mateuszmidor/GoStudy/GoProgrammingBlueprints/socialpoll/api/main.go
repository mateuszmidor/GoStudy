package main

import (
	"context"
	"log"
	"net/http"
	"spconfig"

	mgo "gopkg.in/mgo.v2"
)

type Server struct {
	db *mgo.Session
}
type contextKey struct {
	name string
}

var contextKeyAPIKey = &contextKey{"klucz-api"}

func APIKey(ctx context.Context) (string, bool) {
	key, ok := ctx.Value(contextKeyAPIKey).(string)
	return key, ok
}

func main() {
	config := spconfig.GetConfig()
	var (
		addr  = ":8080"
		mongo = config.MongoDbAddress
	)
	log.Println("Dialing mongo", mongo)
	db, err := mgo.Dial(mongo)
	if err != nil {
		log.Fatalln("Couldnt connect to MongoDB: ", err)
	}
	defer db.Close()
	s := &Server{
		db: db,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/polls/", withCORS(withAPIKey(s.handlePolls)))
	log.Println("Starting http api at", addr)
	http.ListenAndServe(addr, mux)
	log.Println("Stopping http api server...")
}

func withAPIKey(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if !isValidAPIKey(key) {
			respondErr(w, r, http.StatusUnauthorized, "Invalid API key")
			return
		}
		ctx := context.WithValue(r.Context(), contextKeyAPIKey, key)
		fn(w, r.WithContext(ctx))
	}
}

func isValidAPIKey(key string) bool {
	return key == "abc123"
}

func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Location")
		fn(w, r)
	}
}
