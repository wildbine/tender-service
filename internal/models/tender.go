package models

import "time"

type Tender struct {
    ID              int       `json:"id"`
    Name            string    `json:"name"`
    Description     string    `json:"description"`
    Status          string    `json:"status"`
    Version         int       `json:"version"`
    OrganizationID  int       `json:"organization_id"`
    CreatorID       int       `json:"creator_id"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}
type Offer struct {
    ID         int       `json:"id"`
    Name       string    `json:"name"`
    Content    string    `json:"content"`
    Status     string    `json:"status"`
    Version    int       `json:"version"`
    TenderID   int       `json:"tender_id"`
    CreatorID  int       `json:"creator_id"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}