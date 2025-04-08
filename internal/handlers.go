package internal

import (
	"awesomeProject2/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ItemHandler struct {
	service *ItemService
}

func NewItemHandler(service *ItemService) *ItemHandler {
	return &ItemHandler{service: service}
}

func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateItem(r.Context(), &item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	shopID, err := strconv.Atoi(chi.URLParam(r, "shopId"))
	if err != nil {
		http.Error(w, "Invalid shop ID", http.StatusBadRequest)
		return
	}

	sortBy := r.URL.Query().Get("sort")
	var items []models.Item
	var getErr error

	switch sortBy {
	case "category":
		items, getErr = h.service.GetItemsSortedByCategory(r.Context(), shopID)
	case "size":
		items, getErr = h.service.GetItemsSortedBySize(r.Context(), shopID)
	case "brand":
		items, getErr = h.service.GetItemsSortedByBrand(r.Context(), shopID)
	case "price_asc":
		items, getErr = h.service.GetItemsSortedByPrice(r.Context(), shopID, true)
	case "price_desc":
		items, getErr = h.service.GetItemsSortedByPrice(r.Context(), shopID, false)
	default:
		items, getErr = h.service.GetItemsForShop(r.Context(), shopID)
	}

	if getErr != nil {
		http.Error(w, getErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *ItemHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	item, err := h.service.GetItemByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateItem(r.Context(), item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteItem(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
