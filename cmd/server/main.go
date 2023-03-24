package main

import (
	"net/http"

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

		http.HandleFunc("/products", productHandle.CreateProduct)
		http.ListenAndServe(":8000", nil)
}

