package main;

import (
	"github.com/gorilla/mux"
	"github.com/Hycient195/Mongo_Exercise/router"
	"net/http"
	"log"
)

func main() {
	var router *mux.Router = router.Router();
	log.Fatal(http.ListenAndServe(":3001", router));
}