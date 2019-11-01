package main

import (
	"LayeredArchitecture/interfaces/handler"
	"LayeredArchitecture/interfaces/middleware"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/julienschmidt/httprouter"
)

func main() {
	Run(18080)
}

// IsLetter function to check string is aplhanumeric only
var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

// Run start server
func Run(port int) {
	log.Printf("Server running at http://localhost:%d/", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), Routes())
	if err != nil {
		log.Fatal(err)
	}
}

// Routes returns the initialized router
func Routes() *httprouter.Router {
	router := httprouter.New()

	// User Route
	router.GET("/user/get", middleware.Authenticate(handler.HandleUserGet))
	router.POST("/user/signup", handler.HandleUserSignup)
	router.POST("/user/signin", handler.HandleUserSignin)

	// Post Route
	//router.GET("/post/:id", middleware.Authenticate(handler.HandlePostGet))
	router.GET("/post/index", middleware.Authenticate(handler.HandlePostsGet))
	router.POST("/post/create", middleware.Authenticate(handler.HandlePostCreate))
	router.PUT("/post/update", middleware.Authenticate(handler.HandlePostUpdate))
	router.DELETE("/post/delete", middleware.Authenticate(handler.HandlePostDelete))

	return router
}
