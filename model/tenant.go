package model

import (
	"context"
	"github.com/retere/IOTSmartMeasureKWH/entity/tenantentity"
	"github.com/retere/IOTSmartMeasureKWH/model/tenantmodel"
	"gorm.io/gorm"
)

type Tenant interface {
	saveNewTenant(ctx context.Context, payload *tenantentity.Tenant)
}

func NewTransaction(db *gorm.DB) Tenant {
	return tenantmodel.NewTenantRepository(db)
}
