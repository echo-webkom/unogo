package unogo

import "time"

type HealthResponse struct {
	Status string `json:"status"`
}

type ProfileImageResponse struct {
	Source string
}

type UserResponse struct {
	ID                         string              `json:"id"`
	Name                       *string             `json:"name"`
	Email                      string              `json:"email"`
	HasImage                   bool                `json:"hasImage"`
	AlternativeEmail           *string             `json:"alternativeEmail"`
	AlternativeEmailVerifiedAt *time.Time          `json:"alternativeEmailVerifiedAt"`
	Degree                     *DegreeResponse     `json:"degree"`
	Year                       *int                `json:"year"`
	Type                       string              `json:"type"`
	LastSignInAt               *time.Time          `json:"lastSignInAt"`
	UpdatedAt                  *time.Time          `json:"updatedAt"`
	CreatedAt                  *time.Time          `json:"createdAt"`
	HasReadTerms               bool                `json:"hasReadTerms"`
	Birthday                   *time.Time          `json:"birthday"`
	IsPublic                   bool                `json:"isPublic"`
	Groups                     []UserGroupResponse `json:"groups"`
}

type DegreeResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserGroupResponse struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	IsLeader bool   `json:"isLeader"`
}
