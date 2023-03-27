package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jeffersonbraster/apigo/internal/dto"
	"github.com/jeffersonbraster/apigo/internal/entity"
	"github.com/jeffersonbraster/apigo/internal/infra/database"
	entityPkg "github.com/jeffersonbraster/apigo/pkg/entity"
)

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

func (h *ProductHandle) GetProduct(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
		}

		product, err := h.ProductDB.FindByID(id)
		if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
}

func (h *ProductHandle) UpdateProduct(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
		}

		var product entity.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
		}

		product.ID, err = entityPkg.PaserID(id)
		if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
		}

		_, err = h.ProductDB.FindByID(id)
		if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
		}

		err = h.ProductDB.Update(&product)
		if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
		}

		w.WriteHeader(http.StatusOK)
}