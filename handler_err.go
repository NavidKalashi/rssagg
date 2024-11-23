package main

import "net/http"

func handerlErr(w http.ResponseWriter, r *http.Request){
	respondWithError(w, 400, "something went wrong")
}