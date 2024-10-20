package model

import (
	"gorm.io/gorm"
)

type Repository struct {
	Tenant Tenant
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Tenant: NewTransaction(db),
	}
}
