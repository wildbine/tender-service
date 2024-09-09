package repository

import (
    "database/sql"
    "tender-service/internal/models"
)

type TenderRepository interface {
    CreateTender(tender *models.Tender) error
    GetTenderByID(id int) (*models.Tender, error)
    UpdateTender(tender *models.Tender) error
}

type tenderRepository struct {
    db *sql.DB
}

func NewTenderRepository(db *sql.DB) TenderRepository {
    return &tenderRepository{db: db}
}

func (r *tenderRepository) CreateTender(tender *models.Tender) error {
    query := `
    INSERT INTO tenders (name, description, status, version, organization_id, creator_id)
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id, created_at, updated_at;
    `
    return r.db.QueryRow(query, tender.Name, tender.Description, tender.Status, tender.Version, tender.OrganizationID, tender.CreatorID).
        Scan(&tender.ID, &tender.CreatedAt, &tender.UpdatedAt)
}

func (r *tenderRepository) GetTenderByID(id int) (*models.Tender, error) {
    query := `SELECT id, name, description, status, version, organization_id, creator_id, created_at, updated_at FROM tenders WHERE id = $1;`
    var tender models.Tender
    err := r.db.QueryRow(query, id).Scan(&tender.ID, &tender.Name, &tender.Description, &tender.Status, &tender.Version, &tender.OrganizationID, &tender.CreatorID, &tender.CreatedAt, &tender.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return &tender, nil
}

func (r *tenderRepository) UpdateTender(tender *models.Tender) error {
    query := `UPDATE tenders SET name = $1, description = $2, status = $3, version = $4, updated_at = CURRENT_TIMESTAMP WHERE id = $5;`
    _, err := r.db.Exec(query, tender.Name, tender.Description, tender.Status, tender.Version, tender.ID)
    return err
}