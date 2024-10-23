package repository

import (
	"github.com/retere/IOTSmartMeasureKWH/models"
)

type Tenants struct {
	TenantID   string `gorm:"primaryKey" json:"tenant_id"`
	TenantName string `json:"tenant_name"`
	Address    string `json:"address"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func (tenant *Tenants) Save() (*Tenants, error) {
	err := models.Database.Model(&tenant).Create(&tenant).Error
	if err != nil {
		return &Tenants{}, err
	}
	return tenant, nil
}

func SaveExistingTenant(tenant *Tenants, tenantID string) error {
	result := models.Database.Model(&Tenants{}).Where("tenant_id = ?", tenantID).Updates(tenant)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func FindTenantByID(tenantID string, tenant *Tenants) error {
	result := models.Database.Where("tenant_id = ?", tenantID).First(tenant)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
