package model

type Tender struct {
    ID              int    `json:"id"`
    Name            string `json:"name"`
    Description     string `json:"description"`
    ServiceType     string `json:"serviceType"`
    Status          string `json:"status"`
    OrganizationID  int    `json:"organizationId"`
    CreatorUsername string `json:"creatorUsername"`
}