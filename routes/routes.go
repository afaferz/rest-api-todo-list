package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/afaferz/rest-api/controllers"
	"github.com/afaferz/rest-api/middleware"
	"github.com/gorilla/mux"
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

	http.Handle("/", &CORSRouterDecorator{r})
	fmt.Println("Listening")
	log.Panic(
		http.ListenAndServe(":3030", nil),
	)
}

type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, PUT, DELETE, PATCH")
		rw.Header().Add("Access-Control-Allow-Headers",
			"Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	c.R.ServeHTTP(rw, req)
}
