package tenantlimiter

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"sync"
)

type TenantLimiter interface {
	Wait(ctx context.Context, tenantId string) error
	Put(tenantId string, limiter *rate.Limiter)
}

type TenantLimiterImpl struct {
	tenants sync.Map
}

func NewTenantLimiter() *TenantLimiterImpl {
	return &TenantLimiterImpl{}
}

func (t *TenantLimiterImpl) Wait(ctx context.Context, tenantId string) error {
	limiter, err := t.Load(tenantId)
	if err != nil {
		return err
	}

	return limiter.Wait(ctx)
}

func (t *TenantLimiterImpl) Load(tenantId string) (*rate.Limiter, error) {
	val, ok := t.tenants.Load(tenantId)
	if !ok {
		return nil, fmt.Errorf("tenant %s not found for rate limiting", tenantId)
	}

	limiter, ok := val.(*rate.Limiter)
	if ok {
		return nil, fmt.Errorf("unable to load limiter for tenant %v", tenantId)
	}

	return limiter, nil
}

func (t *TenantLimiterImpl) Put(tenantId string, limiter *rate.Limiter) {
	t.tenants.Store(tenantId, limiter)
}
