package repository

import (
	"github.com/retere/IOTSmartMeasureKWH/entity/tenantentity"
	"github.com/retere/IOTSmartMeasureKWH/models"
)

type Tenants struct {
	TenantID   string `json:"tenant_id"`
	TenantName string `json:"tenant_name"`
	Address    string `json:"address"`
}

func SaveTenants(tenant *tenantentity.UpdateTenant) error {
	err := models.Database.Table("tenants").Create(&tenant).Error
	if err != nil {
		return err
	}
	return nil
}

func SaveExistingTenant(tenant *tenantentity.Tenants, tenantID string) error {
	result := models.Database.Model(&Tenants{}).Where("tenant_id = ?", tenantID).Updates(tenant)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func FindTenantByID(tenantID string, tenant *tenantentity.Tenants) error {
	result := models.Database.Where("tenant_id = ?", tenantID).First(tenant)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteById(tenantId string, tenant *Tenants) error {
	result := models.Database.Where("tenant_id = ?", tenantId).Delete(tenant)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func FindAllTenants(page int, pageSize int) ([]Tenants, error) {
	offset := (page - 1) * pageSize

	var tenants []Tenants
	result := models.Database.Limit(pageSize).Offset(offset).Find(&tenants)
	if result.Error != nil {
		return nil, result.Error
	}

	return tenants, nil
}
