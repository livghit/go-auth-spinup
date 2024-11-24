package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	userHandler "github.com/livghit/go-auth-spinup/internal/handlers/user"
	"github.com/livghit/go-auth-spinup/internal/middleware"
	"github.com/livghit/go-auth-spinup/internal/models"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app_key := os.Getenv("APP_KEY")
	if app_key == "" {
		log.Printf("Please generate an app_key and run the programm again")
	}
	log.Printf("APP_KEY IS -> %s", app_key)
	Port := os.Getenv("PORT")
	log.Printf("PORT IS -> %s", Port)

	jwtConfig := models.JWTConfigs{
		SingMethod: jwt.SigningMethodES256,
		Claims: jwt.MapClaims{
			"iss": os.Getenv("SERVER_NAME"),
			"sub": "john doe",
			"exp": "1123",
		},
	}

	t, err := models.CreateToken(jwtConfig)
	if err != nil {
		log.Print(err)
	}

	token := models.JWT{
		Value: t,
	}

	log.Printf("JWT TOKEN -> %s", token.Value)
}

func main() {

	// guest subrouter
	guestRouter := http.NewServeMux()
	guestRouter.HandleFunc("GET /", userHandler.Index)
	guestRouter.HandleFunc("GET /hi", userHandler.Hi)

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
		Addr:    os.Getenv("PORT"),
		Handler: router,
	}

	log.Printf("Server started at port %s", os.Getenv("PORT"))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
