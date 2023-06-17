package route

import (
	userController "api_simple/delivery/controllers/user"
	"os"

	customMiddleware "api_simple/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterPath(r *mux.Router, u *userController.UserController, m *customMiddleware.CustomMiddleware) {

	//LOGGER
	r.StrictSlash(false)
	// auth := r.PathPrefix("/auth").Subrouter()

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}",
	// }))

	//ROUTE AUTH
	// auth.HandleFunc("/login", aa.Login())

	//ROUTE USERS
	user := r.PathPrefix("/users").Subrouter()
	if useMiddleware := os.Getenv("USE_MIDDLEWARE"); useMiddleware == "true" {
		user.Use(m.RateLimit)
	}
	user.HandleFunc("/register", u.Register).Methods(http.MethodPost)

	// r.HandleFunc("/users", uc.GetById(), middlewares.JwtMiddleware()).Methods(http.MethodGet)
	// r.HandleFunc("/users", uc.Update(), middlewares.JwtMiddleware())
	// r.HandleFunc("/users", uc.Delete(), middlewares.JwtMiddleware())

}
