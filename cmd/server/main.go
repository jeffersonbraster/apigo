package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jeffersonbraster/apigo/configs"
	"github.com/jeffersonbraster/apigo/internal/entity"
	"github.com/jeffersonbraster/apigo/internal/infra/database"
	"github.com/jeffersonbraster/apigo/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
		_, err := configs.LoadConfig(".")
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

		r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Post("/products", productHandle.CreateProduct)
		r.Get("/products", productHandle.GetProducts)
		r.Get("/products/{id}", productHandle.GetProduct)		
		r.Put("/products/{id}", productHandle.UpdateProduct)
		r.Delete("/products/{id}", productHandle.DeleteProduct)

		http.ListenAndServe(":8000", r)
}

