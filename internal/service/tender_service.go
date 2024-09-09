package service

import (
    "tender-service/internal/models"
    "tender-service/internal/repository"
)

type TenderService interface {
    CreateTender(tender *model.Tender) error
    ListTenders() ([]model.Tender, error)
}

type tenderService struct {
    repo repository.TenderRepository
}

func NewTenderService(repo repository.TenderRepository) TenderService {
    return &tenderService{repo: repo}
}

func (s *tenderService) CreateTender(tender *model.Tender) error {
    return s.repo.CreateTender(tender)
}

func (s *tenderService) ListTenders() ([]model.Tender, error) {
    return s.repo.ListTenders()
}