package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/jeffersonbraster/apigo/configs"
	"github.com/jeffersonbraster/apigo/internal/entity"
	"github.com/jeffersonbraster/apigo/internal/infra/database"
	"github.com/jeffersonbraster/apigo/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
		configs, err := configs.LoadConfig(".")
		if err != nil {
				panic(err)
		}

		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
				panic(err)
		}

		db.AutoMigrate(&entity.Product{}, &entity.User{})
		productDB := database.NewProduct(db)
		productHandle := handlers.NewProductHandle(productDB)

		userDB := database.NewUser(db)
		userHandle := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExpiresIn)

		r := chi.NewRouter()
		r.Use(middleware.Logger)

		r.Route("/products", func(r chi.Router) {
			r.Use(jwtauth.Verifier(configs.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Post("/", productHandle.CreateProduct)
			r.Get("/", productHandle.GetProducts)
			r.Get("/{id}", productHandle.GetProduct)		
			r.Put("/{id}", productHandle.UpdateProduct)
			r.Delete("/{id}", productHandle.DeleteProduct)
		})
		

		r.Post("/users", userHandle.CreateUser)
		r.Post("/users/login", userHandle.GetJwt)

		http.ListenAndServe(":8000", r)
}

