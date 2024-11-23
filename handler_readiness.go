package main

import "net/http"

func handerlReadiness(w http.ResponseWriter, r *http.Request){
	respondWithJSON(w, 200, struct{}{})
}