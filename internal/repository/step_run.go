package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/hatchet-dev/hatchet/internal/repository/prisma/db"
	"github.com/hatchet-dev/hatchet/internal/repository/prisma/dbsqlc"
)

type ListStepRunsOpts struct {
	WorkflowRunIds []string `validate:"dive,uuid"`

	Status *dbsqlc.StepRunStatus
}

type UpdateStepRunOpts struct {
	IsRerun bool

	RequeueAfter *time.Time

	ScheduleTimeoutAt *time.Time

	Status *db.StepRunStatus

	StartedAt *time.Time

	FailedAt *time.Time

	FinishedAt *time.Time

	CancelledAt *time.Time

	CancelledReason *string

	Error *string

	Input []byte

	Output []byte

	RetryCount *int
}

type UpdateStepRunOverridesDataOpts struct {
	OverrideKey string
	Data        []byte
	CallerFile  *string
}

func StepRunStatusPtr(status db.StepRunStatus) *db.StepRunStatus {
	return &status
}

var ErrStepRunIsNotPending = fmt.Errorf("step run is not pending")
var ErrNoWorkerAvailable = fmt.Errorf("no worker available")
var ErrRateLimitExceeded = fmt.Errorf("rate limit exceeded")

type StepRunUpdateInfo struct {
	JobRunFinalState      bool
	WorkflowRunFinalState bool
	WorkflowRunId         string
	WorkflowRunStatus     string
}

type StepRunAPIRepository interface {
	GetStepRunById(tenantId, stepRunId string) (*db.StepRunModel, error)

	GetFirstArchivedStepRunResult(tenantId, stepRunId string) (*db.StepRunResultArchiveModel, error)
}

type StepRunEngineRepository interface {
	// ListStepRunsForWorkflowRun returns a list of step runs for a workflow run.
	ListStepRuns(ctx context.Context, tenantId string, opts *ListStepRunsOpts) ([]*dbsqlc.GetStepRunForEngineRow, error)

	// ListStepRunsToRequeue returns a list of step runs which are in a requeueable state.
	ListStepRunsToRequeue(ctx context.Context, tenantId string) ([]*dbsqlc.GetStepRunForEngineRow, error)

	// ListStepRunsToReassign returns a list of step runs which are in a reassignable state.
	ListStepRunsToReassign(ctx context.Context, tenantId string) ([]*dbsqlc.GetStepRunForEngineRow, error)

	UpdateStepRun(ctx context.Context, tenantId, stepRunId string, opts *UpdateStepRunOpts) (*dbsqlc.GetStepRunForEngineRow, *StepRunUpdateInfo, error)

	UnlinkStepRunFromWorker(ctx context.Context, tenantId, stepRunId string) error

	// UpdateStepRunOverridesData updates the overrides data field in the input for a step run. This returns the input
	// bytes.
	UpdateStepRunOverridesData(ctx context.Context, tenantId, stepRunId string, opts *UpdateStepRunOverridesDataOpts) ([]byte, error)

	UpdateStepRunInputSchema(ctx context.Context, tenantId, stepRunId string, schema []byte) ([]byte, error)

	AssignStepRunToWorker(ctx context.Context, stepRun *dbsqlc.GetStepRunForEngineRow) (workerId string, dispatcherId string, err error)

	GetStepRunForEngine(ctx context.Context, tenantId, stepRunId string) (*dbsqlc.GetStepRunForEngineRow, error)

	// QueueStepRun is like UpdateStepRun, except that it will only update the step run if it is in
	// a pending state.
	QueueStepRun(ctx context.Context, tenantId, stepRunId string, opts *UpdateStepRunOpts) (*dbsqlc.GetStepRunForEngineRow, error)

	ListStartableStepRuns(ctx context.Context, tenantId, jobRunId string, parentStepRunId *string) ([]*dbsqlc.GetStepRunForEngineRow, error)

	ArchiveStepRunResult(ctx context.Context, tenantId, stepRunId string) error
}
