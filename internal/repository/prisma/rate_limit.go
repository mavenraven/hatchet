package prisma

import (
	"context"
	"fmt"
	"github.com/hatchet-dev/hatchet/internal/services/tenantlimiter"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"

	"github.com/hatchet-dev/hatchet/internal/repository"
	"github.com/hatchet-dev/hatchet/internal/repository/prisma/dbsqlc"
	"github.com/hatchet-dev/hatchet/internal/repository/prisma/sqlchelpers"
	"github.com/hatchet-dev/hatchet/internal/validator"
)

type rateLimitEngineRepository struct {
	pool    *pgxpool.Pool
	v       validator.Validator
	queries *dbsqlc.Queries
	l       *zerolog.Logger
}

func NewRateLimitEngineRepository(pool *pgxpool.Pool, v validator.Validator, l *zerolog.Logger, limiter tenantlimiter.TenantLimiter) repository.RateLimitEngineRepository {
	queries := dbsqlc.New()

	return &rateLimitEngineRepository{
		pool:    pool,
		v:       v,
		queries: queries,
		l:       l,
	}
}

func (r *rateLimitEngineRepository) UpsertRateLimit(ctx context.Context, tenantId string, key string, opts *repository.UpsertRateLimitOpts) (*dbsqlc.RateLimit, error) {
	if err := r.v.Validate(opts); err != nil {
		return nil, err
	}

	upsertParams := dbsqlc.UpsertRateLimitParams{
		Tenantid: sqlchelpers.UUIDFromStr(tenantId),
		Key:      key,
		Limit:    int32(opts.Limit),
	}

	if opts.Duration != nil {
		upsertParams.Window = sqlchelpers.TextFromStr(fmt.Sprintf("1 %s", *opts.Duration))
	}

	rateLimit, err := r.queries.UpsertRateLimit(ctx, r.pool, upsertParams)

	if err != nil {
		return nil, fmt.Errorf("could not upsert rate limit: %w", err)
	}

	return rateLimit, nil
}
