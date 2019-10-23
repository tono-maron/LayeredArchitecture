package interfaces

import (
	"LayeredArchitecture/interfaces/handler"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/julienschmidt/httprouter"
)

// IsLetter function to check string is aplhanumeric only
var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

// Run start server
func Run(port int) error {
	log.Printf("Server running at http://localhost:%d/", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), Routes())
}

// Routes returns the initialized router
func Routes() *httprouter.Router {
	router := httprouter.New()

	// Index Route
	router.GET("/", handler.Index)

	// User Route
	router.GET("/user/get", handler.HandleUserGet)

	return router
}
