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
	//db接続
	infrastructure.NewDBConnection()
	log.Printf("Server running at http://localhost:%d/", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), Routes())
	if err != nil {
		log.Fatal(err)
	}
}

// Routes returns the initialized router
func Routes() *httprouter.Router {
	userPersistence := persistence.NewUserPersistence()
	userUsecase := usecase.NewUserUsecase(userPersistence)
	userHandler := handler.NewUserHandler(userUsecase)
	postPersistence := persistence.NewPostPersistence()
	postUsecase := usecase.NewPostUsecase(postPersistence)
	postHandler := handler.NewPostHandler(postUsecase)

	router := httprouter.New()

	// User Route
	router.GET("/user/get", middleware.Authenticate(userHandler.HandleUserGet))
	router.POST("/user/signup", userHandler.HandleUserSignup)
	router.POST("/user/signin", userHandler.HandleUserSignin)

	// Post Route
	//router.GET("/post/:id", middleware.Authenticate(handler.HandlePostGet))
	router.GET("/post/index", middleware.Authenticate(postHandler.HandlePostsGet))
	router.POST("/post/create", middleware.Authenticate(postHandler.HandlePostCreate))
	router.PUT("/post/update", middleware.Authenticate(postHandler.HandlePostUpdate))
	router.DELETE("/post/delete", middleware.Authenticate(postHandler.HandlePostDelete))

	return router
}