package router

import (
	"github.com/ankit85/movieapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkedWatchedMovie).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovieById).Methods("DELETE")
	router.HandleFunc("/api/deleteAll", controller.DeleteAllMovie).Methods("DELETE")

	return router
}
