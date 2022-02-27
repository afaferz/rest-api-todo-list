package routes

import (
	"log"
	"net/http"

	"github.com/afaferz/rest-api/controllers"
	"github.com/afaferz/rest-api/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/tasks/", controllers.AllTasks).Methods("GET")
	r.HandleFunc("/api/tasks/", controllers.CreateNewTask).Methods("POST")
	r.HandleFunc("/api/tasks/{id}/", controllers.TaskById).Methods("GET")
	r.HandleFunc("/api/tasks/{id}/", controllers.DeleteTask).Methods("DELETE")
	r.HandleFunc("/api/tasks/{id}/", controllers.EditTask).Methods("PATCH")

	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8080", r), handler)
}
