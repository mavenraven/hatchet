package repository

import (
	"context"

	"github.com/hatchet-dev/hatchet/internal/datautils"
	"github.com/hatchet-dev/hatchet/internal/encryption"
	"github.com/hatchet-dev/hatchet/internal/repository/prisma/db"
	"github.com/hatchet-dev/hatchet/internal/repository/prisma/dbsqlc"
	"github.com/hatchet-dev/hatchet/internal/repository/prisma/sqlchelpers"
)

type CreateWorkflowRunOpts struct {
	// (optional) the workflow run display name
	DisplayName *string

	// (required) the workflow version id
	WorkflowVersionId string `validate:"required,uuid"`

	ManualTriggerInput *string `validate:"omitnil,required_without=TriggeringEventId,required_without=Cron,required_without=ScheduledWorkflowId,excluded_with=TriggeringEventId,excluded_with=Cron,excluded_with=ScheduledWorkflowId"`

	// (optional) the event id that triggered the workflow run
	TriggeringEventId *string `validate:"omitnil,uuid,required_without=ManualTriggerInput,required_without=Cron,required_without=ScheduledWorkflowId,excluded_with=ManualTriggerInput,excluded_with=Cron,excluded_with=ScheduledWorkflowId"`

	// (optional) the cron schedule that triggered the workflow run
	Cron         *string `validate:"omitnil,cron,required_without=ManualTriggerInput,required_without=TriggeringEventId,required_without=ScheduledWorkflowId,excluded_with=ManualTriggerInput,excluded_with=TriggeringEventId,excluded_with=ScheduledWorkflowId"`
	CronParentId *string `validate:"omitnil,uuid,required_without=ManualTriggerInput,required_without=TriggeringEventId,required_without=ScheduledWorkflowId,excluded_with=ManualTriggerInput,excluded_with=TriggeringEventId,excluded_with=ScheduledWorkflowId"`

	// (optional) the scheduled trigger
	ScheduledWorkflowId *string `validate:"omitnil,uuid,required_without=ManualTriggerInput,required_without=TriggeringEventId,required_without=Cron,excluded_with=ManualTriggerInput,excluded_with=TriggeringEventId,excluded_with=Cron"`

	InputData []byte

	TriggeredBy string

	GetGroupKeyRun *CreateGroupKeyRunOpts `validate:"omitempty"`

	// (optional) the parent workflow run which this workflow run was triggered from
	ParentId *string `validate:"omitempty,uuid"`

	// (optional) the parent step run id which this workflow run was triggered from
	ParentStepRunId *string `validate:"omitempty,uuid"`

	// (optional) the child key of the workflow run, if this is a child run of a different workflow
	ChildKey *string

	// (optional) the child index of the workflow run, if this is a child run of a different workflow
	ChildIndex *int
}

type CreateGroupKeyRunOpts struct {
	// (optional) the input data
	Input []byte
}

type CreateWorkflowRunOpt func(*CreateWorkflowRunOpts)

func WithParent(
	parentId, parentStepRunId string,
	childIndex int,
	childKey *string,
) CreateWorkflowRunOpt {
	return func(opts *CreateWorkflowRunOpts) {
		opts.ParentId = &parentId
		opts.ParentStepRunId = &parentStepRunId
		opts.ChildIndex = &childIndex
		opts.ChildKey = childKey
	}
}

func GetCreateWorkflowRunOptsFromManual(
	workflowVersion *dbsqlc.GetWorkflowVersionForEngineRow,
	input []byte,
) (*CreateWorkflowRunOpts, error) {
	opts := &CreateWorkflowRunOpts{
		DisplayName:        StringPtr(getWorkflowRunDisplayName(workflowVersion.WorkflowName)),
		WorkflowVersionId:  sqlchelpers.UUIDToStr(workflowVersion.WorkflowVersion.ID),
		ManualTriggerInput: StringPtr(string(input)),
		TriggeredBy:        string(datautils.TriggeredByManual),
		InputData:          input,
	}

	if input != nil {
		if workflowVersion.ConcurrencyLimitStrategy.Valid {
			opts.GetGroupKeyRun = &CreateGroupKeyRunOpts{
				Input: input,
			}
		}
	}

	return opts, nil
}

func GetCreateWorkflowRunOptsFromParent(
	workflowVersion *dbsqlc.GetWorkflowVersionForEngineRow,
	input []byte,
	parentId, parentStepRunId string,
	childIndex int,
	childKey *string,
) (*CreateWorkflowRunOpts, error) {
	opts := &CreateWorkflowRunOpts{
		DisplayName:        StringPtr(getWorkflowRunDisplayName(workflowVersion.WorkflowName)),
		WorkflowVersionId:  sqlchelpers.UUIDToStr(workflowVersion.WorkflowVersion.ID),
		ManualTriggerInput: StringPtr(string(input)),
		TriggeredBy:        string(datautils.TriggeredByParent),
		InputData:          input,
	}

	WithParent(parentId, parentStepRunId, childIndex, childKey)(opts)

	if input != nil {
		if workflowVersion.ConcurrencyLimitStrategy.Valid {
			opts.GetGroupKeyRun = &CreateGroupKeyRunOpts{
				Input: input,
			}
		}
	}

	return opts, nil
}

func GetCreateWorkflowRunOptsFromEvent(eventId string, workflowVersion *dbsqlc.GetWorkflowVersionForEngineRow, input []byte) (*CreateWorkflowRunOpts, error) {
	opts := &CreateWorkflowRunOpts{
		DisplayName:       StringPtr(getWorkflowRunDisplayName(workflowVersion.WorkflowName)),
		WorkflowVersionId: sqlchelpers.UUIDToStr(workflowVersion.WorkflowVersion.ID),
		TriggeringEventId: &eventId,
		TriggeredBy:       string(datautils.TriggeredByEvent),
		InputData:         input,
	}

	if input != nil {
		if workflowVersion.ConcurrencyLimitStrategy.Valid {
			opts.GetGroupKeyRun = &CreateGroupKeyRunOpts{
				Input: input,
			}
		}
	}

	return opts, nil
}

func GetCreateWorkflowRunOptsFromCron(cron, cronParentId string, workflowVersion *dbsqlc.GetWorkflowVersionForEngineRow, input []byte) (*CreateWorkflowRunOpts, error) {
	opts := &CreateWorkflowRunOpts{
		DisplayName:       StringPtr(getWorkflowRunDisplayName(workflowVersion.WorkflowName)),
		WorkflowVersionId: sqlchelpers.UUIDToStr(workflowVersion.WorkflowVersion.ID),
		Cron:              &cron,
		CronParentId:      &cronParentId,
		TriggeredBy:       string(datautils.TriggeredByCron),
		InputData:         input,
	}

	if input != nil {
		if workflowVersion.ConcurrencyLimitStrategy.Valid {
			opts.GetGroupKeyRun = &CreateGroupKeyRunOpts{
				Input: input,
			}
		}
	}

	return opts, nil
}

func GetCreateWorkflowRunOptsFromSchedule(
	scheduledWorkflowId string,
	workflowVersion *dbsqlc.GetWorkflowVersionForEngineRow,
	input []byte,
	fs ...CreateWorkflowRunOpt,
) (*CreateWorkflowRunOpts, error) {
	opts := &CreateWorkflowRunOpts{
		DisplayName:         StringPtr(getWorkflowRunDisplayName(workflowVersion.WorkflowName)),
		WorkflowVersionId:   sqlchelpers.UUIDToStr(workflowVersion.WorkflowVersion.ID),
		ScheduledWorkflowId: &scheduledWorkflowId,
		TriggeredBy:         string(datautils.TriggeredBySchedule),
		InputData:           input,
	}

	if input != nil {
		if workflowVersion.ConcurrencyLimitStrategy.Valid {
			opts.GetGroupKeyRun = &CreateGroupKeyRunOpts{
				Input: input,
			}
		}
	}

	for _, f := range fs {
		f(opts)
	}

	return opts, nil
}

func getWorkflowRunDisplayName(workflowName string) string {
	workflowSuffix, _ := encryption.GenerateRandomBytes(3)

	return workflowName + "-" + workflowSuffix
}

type ListWorkflowRunsOpts struct {
	// (optional) the workflow id
	WorkflowId *string `validate:"omitempty,uuid"`

	// (optional) the workflow version id
	WorkflowVersionId *string `validate:"omitempty,uuid"`

	// (optional) a list of workflow run ids to filter by
	Ids []string `validate:"omitempty,dive,uuid"`

	// (optional) the parent workflow run id
	ParentId *string `validate:"omitempty,uuid"`

	// (optional) the parent step run id
	ParentStepRunId *string `validate:"omitempty,uuid"`

	// (optional) the event id that triggered the workflow run
	EventId *string `validate:"omitempty,uuid"`

	// (optional) the group key for the workflow run
	GroupKey *string

	// (optional) the status of the workflow run
	Statuses *[]db.WorkflowRunStatus

	// (optional) number of events to skip
	Offset *int

	// (optional) number of events to return
	Limit *int

	// (optional) the order by field
	OrderBy *string `validate:"omitempty,oneof=createdAt"`

	// (optional) the order direction
	OrderDirection *string `validate:"omitempty,oneof=ASC DESC"`
}

type WorkflowRunsMetricsOpts struct {
	// (optional) the workflow id
	WorkflowId *string `validate:"omitempty,uuid"`

	// (optional) the workflow version id
	WorkflowVersionId *string `validate:"omitempty,uuid"`

	// (optional) the parent workflow run id
	ParentId *string `validate:"omitempty,uuid"`

	// (optional) the parent step run id
	ParentStepRunId *string `validate:"omitempty,uuid"`

	// (optional) the event id that triggered the workflow run
	EventId *string `validate:"omitempty,uuid"`
}

type ListWorkflowRunsResult struct {
	Rows  []*dbsqlc.ListWorkflowRunsRow
	Count int
}

type CreateWorkflowRunPullRequestOpts struct {
	RepositoryOwner       string
	RepositoryName        string
	PullRequestID         int
	PullRequestTitle      string
	PullRequestNumber     int
	PullRequestHeadBranch string
	PullRequestBaseBranch string
	PullRequestState      string
}

type ListPullRequestsForWorkflowRunOpts struct {
	State *string
}

type ListWorkflowRunRoundRobinsOpts struct {
	// (optional) the workflow id
	WorkflowId *string `validate:"omitempty,uuid"`

	// (optional) the workflow version id
	WorkflowVersionId *string `validate:"omitempty,uuid"`

	// (optional) the status of the workflow run
	Status *db.WorkflowRunStatus

	// (optional) number of events to skip
	Offset *int

	// (optional) number of events to return
	Limit *int
}

type WorkflowRunMetricsCountOpts struct {
	// (optional) the workflow id
	WorkflowId *string `validate:"omitempty,uuid"`

	// (optional) the workflow version id
	WorkflowVersionId *string `validate:"omitempty,uuid"`
}

type WorkflowRunAPIRepository interface {
	// ListWorkflowRuns returns workflow runs for a given workflow version id.
	ListWorkflowRuns(tenantId string, opts *ListWorkflowRunsOpts) (*ListWorkflowRunsResult, error)

	// Counts by status
	WorkflowRunMetricsCount(tenantId string, opts *WorkflowRunsMetricsOpts) (*dbsqlc.WorkflowRunsMetricsCountRow, error)

	// CreateNewWorkflowRun creates a new workflow run for a workflow version.
	CreateNewWorkflowRun(ctx context.Context, tenantId string, opts *CreateWorkflowRunOpts) (*db.WorkflowRunModel, error)

	// GetWorkflowRunById returns a workflow run by id.
	GetWorkflowRunById(tenantId, runId string) (*db.WorkflowRunModel, error)

	CreateWorkflowRunPullRequest(tenantId, workflowRunId string, opts *CreateWorkflowRunPullRequestOpts) (*db.GithubPullRequestModel, error)

	ListPullRequestsForWorkflowRun(tenantId, workflowRunId string, opts *ListPullRequestsForWorkflowRunOpts) ([]db.GithubPullRequestModel, error)
}

type WorkflowRunEngineRepository interface {
	// ListWorkflowRuns returns workflow runs for a given workflow version id.
	ListWorkflowRuns(ctx context.Context, tenantId string, opts *ListWorkflowRunsOpts) (*ListWorkflowRunsResult, error)

	GetChildWorkflowRun(ctx context.Context, parentId, parentStepRunId string, childIndex int, childkey *string) (*dbsqlc.WorkflowRun, error)

	GetScheduledChildWorkflowRun(ctx context.Context, parentId, parentStepRunId string, childIndex int, childkey *string) (*dbsqlc.WorkflowTriggerScheduledRef, error)

	PopWorkflowRunsRoundRobin(ctx context.Context, tenantId, workflowId string, maxRuns int) ([]*dbsqlc.WorkflowRun, error)

	// CreateNewWorkflowRun creates a new workflow run for a workflow version.
	CreateNewWorkflowRun(ctx context.Context, tenantId string, opts *CreateWorkflowRunOpts) (string, error)

	// GetWorkflowRunById returns a workflow run by id.
	GetWorkflowRunById(ctx context.Context, tenantId, runId string) (*dbsqlc.GetWorkflowRunRow, error)
}
