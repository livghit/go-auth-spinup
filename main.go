package main

import (
	"io"
	"log"
	"net/http"

	handlers "github.com/livghit/go-auth-spinup/internal/handlers/user"
	"github.com/livghit/go-auth-spinup/internal/middleware"
)

func main() {

	// guest subrouter
	guestRouter := http.NewServeMux()
	guestRouter.HandleFunc("GET /", handlers.Index)
	guestRouter.HandleFunc("GET /hi", handlers.Hi)

	baseMiddleware := middleware.CreateMiddlewareStack(
		middleware.Logger,
	)

	// subrouter for mainRouter that strips the admin prefix
	adminRouter := http.NewServeMux()
	adminRouter.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello from admin")
	})
	adminMiddleware := middleware.CreateMiddlewareStack(
		baseMiddleware,
		middleware.Auth,
	)

	// the main router
	router := http.NewServeMux()

	router.Handle("/", baseMiddleware(guestRouter))
	router.Handle("/admin/", adminMiddleware(http.StripPrefix("/admin", adminRouter)))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Default().Printf("Server started")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
