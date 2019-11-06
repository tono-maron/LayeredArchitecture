package main

import (
	"LayeredArchitecture/infrastructure"
	"LayeredArchitecture/infrastructure/persistence"
	"LayeredArchitecture/interfaces/handler"
	"LayeredArchitecture/interfaces/middleware"
	"LayeredArchitecture/usecase"
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
	//NewDBConnection create db connection info.
	err := infrastructure.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server running at http://localhost:%d/", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), Routes())
	if err != nil {
		log.Fatal(err)
	}
}

// Routes returns the initialized router
func Routes() *httprouter.Router {
	userPersistence := persistence.NewUserPersistence(infrastructure.DB)
	userUsecase := usecase.NewUserUsecase(userPersistence)
	userHandler := handler.NewUserHandler(userUsecase)
	postPersistence := persistence.NewPostPersistence(infrastructure.DB)
	postUsecase := usecase.NewPostUsecase(postPersistence)
	postHandler := handler.NewPostHandler(postUsecase)

	router := httprouter.New()

	// User Route
	router.GET("/user/get", middleware.Authenticate(userHandler.HandleUserGet))
	router.POST("/user/signup", userHandler.HandleUserSignup)
	router.POST("/user/signin", userHandler.HandleUserSignin)

	// Post Route
	router.GET("/post/:id", middleware.Authenticate(postHandler.HandlePostGet))
	router.GET("/posts/index", middleware.Authenticate(postHandler.HandlePostsGet))
	router.POST("/post/create", middleware.Authenticate(postHandler.HandlePostCreate))
	router.PUT("/post/:id", middleware.Authenticate(postHandler.HandlePostUpdate))
	router.DELETE("/post/:id", middleware.Authenticate(postHandler.HandlePostDelete))

	return router
}
