package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
	"strconv"
)

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"` // in cents
}

type ProductHandler struct {
	nextId   int
	products map[int]Product
	router   *mux.Router
}

func NewProductHandler(router *mux.Router) *ProductHandler {
	return &ProductHandler{router: router}
}

func (h *ProductHandler) Init() {
	h.initValues()
	h.router.HandleFunc("/products", h.getAllProducts).Methods("GET")
	h.router.HandleFunc("/products/{id:[0-9]+}", h.getProduct).Methods("GET")
	h.router.HandleFunc("/products", h.createProduct).Methods("POST")
	h.router.HandleFunc("/products/{id:[0-9]+}", h.getProduct).Methods("PUT")
}

func (h *ProductHandler) initValues() {
	h.nextId = 3
	h.products = map[int]Product{
		1: {Id: 1, Name: "Apples", Category: "Food", Price: 149},
		2: {Id: 2, Name: "Macaroni and Cheese", Category: "Food", Price: 769},
		3: {Id: 3, Name: "ABC Smart TV", Category: "Electronics", Price: 39999},
		4: {Id: 4, Name: "Motor Oil", Category: "Automobile", Price: 2288},
		5: {Id: 5, Name: "Floral Sleeveless Blouse", Category: "Clothing", Price: 2150},
	}
}

func (h *ProductHandler) getAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []Product
	for _, p := range h.products {
		products = append(products, p)
	}
	sort.Slice(products, func(i, j int) bool {
		return products[i].Id < products[j].Id
	})
	respondWithJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if p, ok := h.products[id]; ok {
		respondWithJSON(w, http.StatusOK, p)
	} else {
		respondWithError(w, http.StatusNotFound, "Product not found")
	}
}

func (h *ProductHandler) createProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	p.Id = h.nextId
	h.products[h.nextId] = p
	h.nextId = h.nextId + 1
	respondWithJSON(w, http.StatusCreated, p)
}

func (h *ProductHandler) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	var p Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if _, ok := h.products[id]; ok {
		h.products[id] = p
		respondWithJSON(w, http.StatusOK, p)
	} else {
		respondWithError(w, http.StatusNotFound, "Product not found")
	}
}

func (h *ProductHandler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if _, ok := h.products[id]; ok {
		delete(h.products, id)
		respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		respondWithError(w, http.StatusNotFound, "Product not found")
	}
}
