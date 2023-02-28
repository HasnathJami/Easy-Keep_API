package utils

import (
	"encoding/json"
	"net/http"
)

func ProcessSuccessful(status int, message string, w http.ResponseWriter, data any) {
    w.WriteHeader(status)
   if data != nil{
       json.NewEncoder(w).Encode(data)
   }else{
       json.NewEncoder(w).Encode(message)
   }
  
}