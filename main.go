package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HasnathJami/Easy-Keep/router"
	"github.com/HasnathJami/Easy-Keep/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	utils.CheckSimpleError(err)

	var r *mux.Router
	r = router.Router()

	fmt.Println("Starting Server at Port: "+os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("HOST")+":"+os.Getenv("PORT"), r))
    
}