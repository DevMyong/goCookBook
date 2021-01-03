package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type serverResponse struct {
	Message string `json:"message"`
}

type serverRequest struct{
	Name string `json:"name"`
}


func serverHandler(w http.ResponseWriter, r *http.Request){

	var request serverRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := serverResponse{Message: "Hello " +request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

func main(){
	port := 8080

	http.HandleFunc("/server", serverHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))


}