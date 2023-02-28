package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CheckSimpleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckError(err error, message string, status int, w http.ResponseWriter){
	if err != nil{
		msg := fmt.Sprint(message)
		//w.Write([]byte(msg))
		json.NewEncoder(w).Encode(msg)
		log.Fatal(err)
		//fmt.Println(err)
	}
}