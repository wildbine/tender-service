package handler

import (
    "encoding/json"
    "net/http"
    "tender-service/internal/models"
    "tender-service/internal/service"
)

func CreateTenderHandler(svc service.TenderService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
            return
        }

        var tender model.Tender
        if err := json.NewDecoder(r.Body).Decode(&tender); err != nil {
            http.Error(w, "Bad request body", http.StatusBadRequest)
            return
        }

        if err := svc.CreateTender(&tender); err != nil {
            http.Error(w, "Failed to create tender", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(tender)
    }
}

func ListTendersHandler(svc service.TenderService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
            return
        }

        tenders, err := svc.ListTenders()
        if err != nil {
            http.Error(w, "Failed to retrieve tenders", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(tenders)
    }
}