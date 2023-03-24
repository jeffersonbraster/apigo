package main

import (
	"encoding/json"
	"net/http"

	"github.com/jeffersonbraster/apigo/configs"
	"github.com/jeffersonbraster/apigo/internal/dto"
	"github.com/jeffersonbraster/apigo/internal/entity"
	"github.com/jeffersonbraster/apigo/internal/infra/database"
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
		productHandle := NewProductHandle(productDB)

		http.HandleFunc("/products", productHandle.CreateProduct)
		http.ListenAndServe(":8000", nil)
}

type ProductHandle struct {
	ProductDB database.ProductInterface
}

func NewProductHandle(db database.ProductInterface) *ProductHandle {
	return &ProductHandle{
		ProductDB: db,
	}
}

func (h *ProductHandle) CreateProduct(w http.ResponseWriter, r *http.Request) {
		var product dto.CreateProductInput

		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
		}

		p, err := entity.NewProduct(product.Name, product.Price)
		if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
		}

		err = h.ProductDB.Create(p)
		if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
		}

		w.WriteHeader(http.StatusCreated)
}