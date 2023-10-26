package router

import (
	"github.com/Hycient195/Mongo_Exercise/controller"
	"github.com/gorilla/mux"
)


func Router() *mux.Router {
	var r = mux.NewRouter();

	r.HandleFunc("/post", controller.HandleInsertOnePost).Methods("POST");
	// r.HandleFunc("/post/{postId}", controller.HandleGetOnePost).Methods("GET");
	r.HandleFunc("/post/{postId}", controller.HandleUpdateOnePost).Methods("PATCH");
	r.HandleFunc("/post/{postId}", controller.HandleDeleteOnePost).Methods("DELETE");
	r.HandleFunc("/posts", controller.HandleDeleteManyPosts).Methods("DELETE");
	r.HandleFunc("/posts", controller.HandleGetAllPosts).Methods("GET");

	return r;
};