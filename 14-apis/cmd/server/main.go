package main

import (
	"net/http"

	"github.com/Lucioschenkel/goexpert/14-apis/configs"
	_ "github.com/Lucioschenkel/goexpert/14-apis/docs"
	"github.com/Lucioschenkel/goexpert/14-apis/internal/entity"
	"github.com/Lucioschenkel/goexpert/14-apis/internal/infra/database"
	"github.com/Lucioschenkel/goexpert/14-apis/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title	GoExpert API
// @version 1.0
// @description Product API with authentication.
// @termsOfService	http://swagger.io/terms/

// @contact.name		Lucio Schenkel
// @contact.url	https://lucioschenkel.com
// @contact.email contato@lucioschenkel.com

// @license.name Lucio Schenkel License
// @license.url https://lucioschenkel.com/license

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", config.TokenAuth))
	r.Use(middleware.WithValue("JwtExpiresIn", config.JWTExpiresIn))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
		r.Get("/", productHandler.GetProducts)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/login", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))

	http.ListenAndServe(":8080", r)
}
