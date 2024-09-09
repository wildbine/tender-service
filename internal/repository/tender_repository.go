package repository

import (
    "database/sql"
    "tender-service/internal/models"
)

type TenderRepository interface {
    CreateTender(tender *model.Tender) error
    ListTenders() ([]model.Tender, error)
}

type tenderRepository struct {
    db *sql.DB
}

func NewTenderRepository(db *sql.DB) TenderRepository {
    return &tenderRepository{db: db}
}

func (r *tenderRepository) CreateTender(tender *model.Tender) error {
    query := `
    INSERT INTO tender (name, description, service_type, status, organization_id, creator_username)
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id;
    `
    return r.db.QueryRow(query, tender.Name, tender.Description, tender.ServiceType, tender.Status, tender.OrganizationID, tender.CreatorUsername).Scan(&tender.ID)
}

func (r *tenderRepository) ListTenders() ([]model.Tender, error) {
    rows, err := r.db.Query("SELECT id, name, description, service_type, status, organization_id, creator_username FROM tender")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tenders []model.Tender
    for rows.Next() {
        var tender model.Tender
        err := rows.Scan(&tender.ID, &tender.Name, &tender.Description, &tender.ServiceType, &tender.Status, &tender.OrganizationID, &tender.CreatorUsername)
        if err != nil {
            return nil, err
        }
        tenders = append(tenders, tender)
    }
    return tenders, nil
}