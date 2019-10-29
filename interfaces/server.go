package interfaces

import (
	"LayeredArchitecture/interfaces/handler"
	"LayeredArchitecture/interfaces/middleware"
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

	//Automatic OPTIONS response and CORS
	router.GlobalOPTIONS = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := writer.Header()
			header.Set("Access-Control-Allow-Methods", request.Header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		writer.WriteHeader(http.StatusNoContent)
	})

	// Index Route
	router.GET("/", handler.Index)

	// User Route
	router.GET("/user/get", middleware.Authenticate(handler.HandleUserGet))
	router.POST("/user/signup", handler.HandleUserSignup)
	router.POST("/user/signin", handler.HandleUserSignin)

	return router
}
