package routers

import (
	"github.com/gorilla/mux"
	"net/http"  // This import might be unnecessary if you're not using it directly
	"github.com/filmons/go-ingnerd-backend/interfaces/controllers"
)

// SetupRouter configures and returns a new router.
func SetupRouter(todoController *controllers.TodoController) *mux.Router {
	router := mux.NewRouter()

	// Using the router to handle a path. If you are using handlers from net/http, then keep the import
	router.HandleFunc("/todos", todoController.GetAllTodos).Methods("GET")

	// Example of directly using http.HandlerFunc if needed
	router.HandleFunc("/example", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Example endpoint!"))
	})).Methods("GET")

	return router
}
