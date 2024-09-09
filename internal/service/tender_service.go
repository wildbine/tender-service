package service

import (
    "errors"
    "tender-service/internal/models"
    "tender-service/internal/repository"
)

type TenderService interface {
    CreateTender(tender *models.Tender) error
    GetTenderByID(id int) (*models.Tender, error)
    UpdateTender(tender *models.Tender) error
}

type tenderService struct {
    repo repository.TenderRepository
}

func NewTenderService(repo repository.TenderRepository) TenderService {
    return &tenderService{repo: repo}
}

func (s *tenderService) CreateTender(tender *models.Tender) error {
    tender.Status = "CREATED"
    tender.Version = 1
    return s.repo.CreateTender(tender)
}

func (s *tenderService) GetTenderByID(id int) (*models.Tender, error) {
    tender, err := s.repo.GetTenderByID(id)
    if err != nil {
        return nil, err
    }
    return tender, nil
}

func (s *tenderService) UpdateTender(tender *models.Tender) error {
    if tender.Status != "CREATED" && tender.Status != "PUBLISHED" {
        return errors.New("cannot update tender with current status")
    }
    tender.Version++
    return s.repo.UpdateTender(tender)
}