package handler

import (
    "encoding/json"
    "net/http"
    "strconv"
    "tender-service/internal/models"
    "tender-service/internal/service"
)

type TenderHandler struct {
    service service.TenderService
}

func NewTenderHandler(service service.TenderService) *TenderHandler {
    return &TenderHandler{service: service}
}

func (h *TenderHandler) CreateTenderHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var tender models.Tender
    if err := json.NewDecoder(r.Body).Decode(&tender); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := h.service.CreateTender(&tender); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(tender)
}

func (h *TenderHandler) GetTenderHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    tender, err := h.service.GetTenderByID(id)
    if err != nil {
        http.Error(w, "Tender not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(tender)
}