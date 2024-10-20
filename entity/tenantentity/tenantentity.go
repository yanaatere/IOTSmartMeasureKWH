package tenantentity

import "time"

type Tenant struct {
	TenantID   int       `gorm:"primaryKey;autoIncrement" json:"tenant_id"`
	TenantName string    `gorm:"type:varchar(255);not null" json:"tenant_name"`
	Address    string    `gorm:"type:text" json:"address"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Tenants struct {
	TenantID   string `json:"tenant_id"`
	TenantName string `json:"tenant_name"`
	Address    string `json:"address"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
type TenantData struct {
	TenantID   string `json:"tenant_id"`
	TenantName string `json:"tenant_name"`
	Address    string `json:"address"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type TenantRequest struct {
	TenantID   string `json:"tenant_id"`
	TenantName string `json:"tenant_name"`
	Address    string `json:"address"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type FindTenant struct {
	TenantID   string `json:"tenant_id"`
	TenantName string `json:"tenant_name"`
	Address    string `json:"address"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
