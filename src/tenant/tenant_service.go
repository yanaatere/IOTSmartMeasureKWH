package tenant

import (
	"context"
	"github.com/retere/IOTSmartMeasureKWH/entity/tenantentity"
	"github.com/retere/IOTSmartMeasureKWH/helpers/errorcodehandling"
	repositories "github.com/retere/IOTSmartMeasureKWH/model"
)

type service struct {
	repo *repositories.Repository
	err  *errorcodehandling.CodeError
}

func NewService(repo *repositories.Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) insertNewTenant(ctx context.Context, payload *tenantentity.Tenant) (*tenantentity.Tenant, error) {
	var err = s.repo.Tenant.save(ctx, payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
