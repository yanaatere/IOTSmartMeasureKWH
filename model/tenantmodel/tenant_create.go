package tenantmodel

import (
	"context"
	"github.com/retere/IOTSmartMeasureKWH/entity/tenantentity"
)

func (r *TenantRepository) saveNewTenant(ctx context.Context, payload *tenantentity.Tenant) error {
	err := r.db.Create(payload).Error()
	if err != nil {
		parsed := r.codeE
	}
	return err
}
