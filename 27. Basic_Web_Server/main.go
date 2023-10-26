package main;

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("Introduction to Web servers in Golang");


	CreateWebServerWithRouter();
}

func CreateWebServerWithoutRouter() {
	http.HandleFunc("/", HandleBaseUrl);

	log.Fatal(http.ListenAndServe(":3001", nil));
}

func CreateWebServerWithRouter() {
	routerMultiplexer := http.NewServeMux();

	routerMultiplexer.HandleFunc("/", HandleBaseUrl)
	routerMultiplexer.HandleFunc("/send-json", SendJSON);
	routerMultiplexer.HandleFunc("/receive-json", ReceiveJSON)

	log.Fatal(http.ListenAndServe(":3001", routerMultiplexer));
}

func HandleBaseUrl(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>First Golang Web Server</>`));
}

/* =============================================== */
/* Sending Back JSON Response To Request Initiator */
/* =============================================== */
type Job struct{
	Desc string `json:"description"`
	Pay int			`json:"remuneration"`
	IsRemote bool	`json:"workStyle"`
}

func SendJSON(w http.ResponseWriter, r *http.Request) {
	var Analyst Job = Job{"Data Analyst", 880000, true};
	w.Header().Set("Content-Type", "app76lication/json");
	json.NewEncoder(w).Encode(Analyst);
}

func ReceiveJSON(w http.ResponseWriter, r *http.Request) {
	var jobInstance Job;
	json.NewDecoder(r.Body).Decode(&jobInstance);
	json.NewEncoder(w).Encode(jobInstance);
}