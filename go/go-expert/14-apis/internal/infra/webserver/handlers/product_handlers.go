package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Lucioschenkel/goexpert/14-apis/internal/dto"
	"github.com/Lucioschenkel/goexpert/14-apis/internal/entity"
	"github.com/Lucioschenkel/goexpert/14-apis/internal/infra/database"
	pkgEntity "github.com/Lucioschenkel/goexpert/14-apis/pkg/entity"
	"github.com/go-chi/chi"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// Create Product godoc
// @Summary 		Create a product
// @Description Create a product
// @Tags 				products
// @Accept 			json
// @Produce 		json
// @Param 			product 					body dto.CreateProductInput true "product request"
// @Success 		201
// @Failure 		500
// @Router 			/products 				[post]
// @Security 		ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetProduct godoc
// @Summary 		Get a product
// @Description Get a product
// @Tags 				products
// @Accept 			json
// @Produce 		json
// @Param 			id 							path 			string 					true "product id" Format(uuid)
// @Success 		200 						{object} 	entity.Product
// @Failure 		400
// @Failure 		404
// @Router 			/products/{id} 	[get]
// @Security 		ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct Product godoc
// @Summary 		Product a product
// @Description Product a product
// @Tags 				products
// @Accept 			json
// @Produce 		json
// @Param 			product 					body dto.CreateProductInput true "product request"
// @Param 			id 								path 			string 					true "product id" Format(uuid)
// @Success 		201
// @Failure 		500
// @Router 			/products 				[put]
// @Security 		ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
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

	product.ID, err = pkgEntity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindById(id)
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

// DeleteProduct godoc
// @Summary 		Delete a product
// @Description Delete a product
// @Tags 				products
// @Accept 			json
// @Produce 		json
// @Param 			id 							path 			string 					true "product id" Format(uuid)
// @Success 		200 						{object} 	entity.Product
// @Failure 		400
// @Failure 		404
// @Failure 		500
// @Router 			/products/{id} 	[delete]
// @Security 		ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// ListProducts godoc
// @Summary 		List products
// @Description List products
// @Tags 				products
// @Accept 			json
// @Produce 		json
// @Param 			page 					query 	string 					false "page number"
// @Param 			limit 				query 	string 					false "limit"
// @Success 		200 					{array} entity.Product
// @Failure 		500
// @Router 			/products 		[get]
// @Security 		ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
