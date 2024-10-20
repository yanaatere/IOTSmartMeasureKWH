package tenantmodel

import (
	"github.com/retere/IOTSmartMeasureKWH/helpers/errorcodehandling"
	"gorm.io/gorm"
)

type TenantRepository struct {
	db        *gorm.DB
	codeError *errorcodehandling.CodeError
}

func NewTenantRepository(db *gorm.DB) *TenantRepository {
	return &TenantRepository{db: db}
}
