// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package dbsqlc

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type ConcurrencyLimitStrategy string

const (
	ConcurrencyLimitStrategyCANCELINPROGRESS ConcurrencyLimitStrategy = "CANCEL_IN_PROGRESS"
	ConcurrencyLimitStrategyDROPNEWEST       ConcurrencyLimitStrategy = "DROP_NEWEST"
	ConcurrencyLimitStrategyQUEUENEWEST      ConcurrencyLimitStrategy = "QUEUE_NEWEST"
	ConcurrencyLimitStrategyGROUPROUNDROBIN  ConcurrencyLimitStrategy = "GROUP_ROUND_ROBIN"
)

func (e *ConcurrencyLimitStrategy) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ConcurrencyLimitStrategy(s)
	case string:
		*e = ConcurrencyLimitStrategy(s)
	default:
		return fmt.Errorf("unsupported scan type for ConcurrencyLimitStrategy: %T", src)
	}
	return nil
}

type NullConcurrencyLimitStrategy struct {
	ConcurrencyLimitStrategy ConcurrencyLimitStrategy `json:"ConcurrencyLimitStrategy"`
	Valid                    bool                     `json:"valid"` // Valid is true if ConcurrencyLimitStrategy is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullConcurrencyLimitStrategy) Scan(value interface{}) error {
	if value == nil {
		ns.ConcurrencyLimitStrategy, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ConcurrencyLimitStrategy.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullConcurrencyLimitStrategy) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ConcurrencyLimitStrategy), nil
}

type InviteLinkStatus string

const (
	InviteLinkStatusPENDING  InviteLinkStatus = "PENDING"
	InviteLinkStatusACCEPTED InviteLinkStatus = "ACCEPTED"
	InviteLinkStatusREJECTED InviteLinkStatus = "REJECTED"
)

func (e *InviteLinkStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = InviteLinkStatus(s)
	case string:
		*e = InviteLinkStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for InviteLinkStatus: %T", src)
	}
	return nil
}

type NullInviteLinkStatus struct {
	InviteLinkStatus InviteLinkStatus `json:"InviteLinkStatus"`
	Valid            bool             `json:"valid"` // Valid is true if InviteLinkStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullInviteLinkStatus) Scan(value interface{}) error {
	if value == nil {
		ns.InviteLinkStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.InviteLinkStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullInviteLinkStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.InviteLinkStatus), nil
}

type JobRunStatus string

const (
	JobRunStatusPENDING   JobRunStatus = "PENDING"
	JobRunStatusRUNNING   JobRunStatus = "RUNNING"
	JobRunStatusSUCCEEDED JobRunStatus = "SUCCEEDED"
	JobRunStatusFAILED    JobRunStatus = "FAILED"
	JobRunStatusCANCELLED JobRunStatus = "CANCELLED"
)

func (e *JobRunStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = JobRunStatus(s)
	case string:
		*e = JobRunStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for JobRunStatus: %T", src)
	}
	return nil
}

type NullJobRunStatus struct {
	JobRunStatus JobRunStatus `json:"JobRunStatus"`
	Valid        bool         `json:"valid"` // Valid is true if JobRunStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullJobRunStatus) Scan(value interface{}) error {
	if value == nil {
		ns.JobRunStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.JobRunStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullJobRunStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.JobRunStatus), nil
}

type StepRunStatus string

const (
	StepRunStatusPENDING           StepRunStatus = "PENDING"
	StepRunStatusPENDINGASSIGNMENT StepRunStatus = "PENDING_ASSIGNMENT"
	StepRunStatusASSIGNED          StepRunStatus = "ASSIGNED"
	StepRunStatusRUNNING           StepRunStatus = "RUNNING"
	StepRunStatusSUCCEEDED         StepRunStatus = "SUCCEEDED"
	StepRunStatusFAILED            StepRunStatus = "FAILED"
	StepRunStatusCANCELLED         StepRunStatus = "CANCELLED"
)

func (e *StepRunStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = StepRunStatus(s)
	case string:
		*e = StepRunStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for StepRunStatus: %T", src)
	}
	return nil
}

type NullStepRunStatus struct {
	StepRunStatus StepRunStatus `json:"StepRunStatus"`
	Valid         bool          `json:"valid"` // Valid is true if StepRunStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStepRunStatus) Scan(value interface{}) error {
	if value == nil {
		ns.StepRunStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.StepRunStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStepRunStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.StepRunStatus), nil
}

type TenantMemberRole string

const (
	TenantMemberRoleOWNER  TenantMemberRole = "OWNER"
	TenantMemberRoleADMIN  TenantMemberRole = "ADMIN"
	TenantMemberRoleMEMBER TenantMemberRole = "MEMBER"
)

func (e *TenantMemberRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TenantMemberRole(s)
	case string:
		*e = TenantMemberRole(s)
	default:
		return fmt.Errorf("unsupported scan type for TenantMemberRole: %T", src)
	}
	return nil
}

type NullTenantMemberRole struct {
	TenantMemberRole TenantMemberRole `json:"TenantMemberRole"`
	Valid            bool             `json:"valid"` // Valid is true if TenantMemberRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTenantMemberRole) Scan(value interface{}) error {
	if value == nil {
		ns.TenantMemberRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TenantMemberRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTenantMemberRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TenantMemberRole), nil
}

type VcsProvider string

const (
	VcsProviderGITHUB VcsProvider = "GITHUB"
)

func (e *VcsProvider) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VcsProvider(s)
	case string:
		*e = VcsProvider(s)
	default:
		return fmt.Errorf("unsupported scan type for VcsProvider: %T", src)
	}
	return nil
}

type NullVcsProvider struct {
	VcsProvider VcsProvider `json:"VcsProvider"`
	Valid       bool        `json:"valid"` // Valid is true if VcsProvider is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVcsProvider) Scan(value interface{}) error {
	if value == nil {
		ns.VcsProvider, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VcsProvider.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVcsProvider) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VcsProvider), nil
}

type WorkerStatus string

const (
	WorkerStatusACTIVE   WorkerStatus = "ACTIVE"
	WorkerStatusINACTIVE WorkerStatus = "INACTIVE"
)

func (e *WorkerStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = WorkerStatus(s)
	case string:
		*e = WorkerStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for WorkerStatus: %T", src)
	}
	return nil
}

type NullWorkerStatus struct {
	WorkerStatus WorkerStatus `json:"WorkerStatus"`
	Valid        bool         `json:"valid"` // Valid is true if WorkerStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullWorkerStatus) Scan(value interface{}) error {
	if value == nil {
		ns.WorkerStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.WorkerStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullWorkerStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.WorkerStatus), nil
}

type WorkflowRunStatus string

const (
	WorkflowRunStatusPENDING   WorkflowRunStatus = "PENDING"
	WorkflowRunStatusRUNNING   WorkflowRunStatus = "RUNNING"
	WorkflowRunStatusSUCCEEDED WorkflowRunStatus = "SUCCEEDED"
	WorkflowRunStatusFAILED    WorkflowRunStatus = "FAILED"
	WorkflowRunStatusQUEUED    WorkflowRunStatus = "QUEUED"
)

func (e *WorkflowRunStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = WorkflowRunStatus(s)
	case string:
		*e = WorkflowRunStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for WorkflowRunStatus: %T", src)
	}
	return nil
}

type NullWorkflowRunStatus struct {
	WorkflowRunStatus WorkflowRunStatus `json:"WorkflowRunStatus"`
	Valid             bool              `json:"valid"` // Valid is true if WorkflowRunStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullWorkflowRunStatus) Scan(value interface{}) error {
	if value == nil {
		ns.WorkflowRunStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.WorkflowRunStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullWorkflowRunStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.WorkflowRunStatus), nil
}

type APIToken struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	ExpiresAt pgtype.Timestamp `json:"expiresAt"`
	Revoked   bool             `json:"revoked"`
	Name      pgtype.Text      `json:"name"`
	TenantId  pgtype.UUID      `json:"tenantId"`
}

type Action struct {
	Description pgtype.Text `json:"description"`
	TenantId    pgtype.UUID `json:"tenantId"`
	ActionId    string      `json:"actionId"`
	ID          pgtype.UUID `json:"id"`
}

type ActionToWorker struct {
	B pgtype.UUID `json:"B"`
	A pgtype.UUID `json:"A"`
}

type Dispatcher struct {
	ID              pgtype.UUID      `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	DeletedAt       pgtype.Timestamp `json:"deletedAt"`
	LastHeartbeatAt pgtype.Timestamp `json:"lastHeartbeatAt"`
	IsActive        bool             `json:"isActive"`
}

type Event struct {
	ID             pgtype.UUID      `json:"id"`
	CreatedAt      pgtype.Timestamp `json:"createdAt"`
	UpdatedAt      pgtype.Timestamp `json:"updatedAt"`
	DeletedAt      pgtype.Timestamp `json:"deletedAt"`
	Key            string           `json:"key"`
	TenantId       pgtype.UUID      `json:"tenantId"`
	ReplayedFromId pgtype.UUID      `json:"replayedFromId"`
	Data           []byte           `json:"data"`
}

type GetGroupKeyRun struct {
	ID                pgtype.UUID      `json:"id"`
	CreatedAt         pgtype.Timestamp `json:"createdAt"`
	UpdatedAt         pgtype.Timestamp `json:"updatedAt"`
	DeletedAt         pgtype.Timestamp `json:"deletedAt"`
	TenantId          pgtype.UUID      `json:"tenantId"`
	WorkerId          pgtype.UUID      `json:"workerId"`
	TickerId          pgtype.UUID      `json:"tickerId"`
	Status            StepRunStatus    `json:"status"`
	Input             []byte           `json:"input"`
	Output            pgtype.Text      `json:"output"`
	RequeueAfter      pgtype.Timestamp `json:"requeueAfter"`
	Error             pgtype.Text      `json:"error"`
	StartedAt         pgtype.Timestamp `json:"startedAt"`
	FinishedAt        pgtype.Timestamp `json:"finishedAt"`
	TimeoutAt         pgtype.Timestamp `json:"timeoutAt"`
	CancelledAt       pgtype.Timestamp `json:"cancelledAt"`
	CancelledReason   pgtype.Text      `json:"cancelledReason"`
	CancelledError    pgtype.Text      `json:"cancelledError"`
	WorkflowRunId     pgtype.UUID      `json:"workflowRunId"`
	ScheduleTimeoutAt pgtype.Timestamp `json:"scheduleTimeoutAt"`
}

type GithubAppInstallation struct {
	ID                      pgtype.UUID      `json:"id"`
	CreatedAt               pgtype.Timestamp `json:"createdAt"`
	UpdatedAt               pgtype.Timestamp `json:"updatedAt"`
	DeletedAt               pgtype.Timestamp `json:"deletedAt"`
	GithubAppOAuthId        pgtype.UUID      `json:"githubAppOAuthId"`
	InstallationId          int32            `json:"installationId"`
	AccountName             string           `json:"accountName"`
	AccountId               int32            `json:"accountId"`
	AccountAvatarURL        pgtype.Text      `json:"accountAvatarURL"`
	InstallationSettingsURL pgtype.Text      `json:"installationSettingsURL"`
	Config                  []byte           `json:"config"`
	TenantId                pgtype.UUID      `json:"tenantId"`
	TenantVcsProviderId     pgtype.UUID      `json:"tenantVcsProviderId"`
}

type GithubAppInstallationToGithubWebhook struct {
	A pgtype.UUID `json:"A"`
	B pgtype.UUID `json:"B"`
}

type GithubAppOAuth struct {
	ID           pgtype.UUID      `json:"id"`
	CreatedAt    pgtype.Timestamp `json:"createdAt"`
	UpdatedAt    pgtype.Timestamp `json:"updatedAt"`
	DeletedAt    pgtype.Timestamp `json:"deletedAt"`
	GithubUserID int32            `json:"githubUserID"`
	AccessToken  []byte           `json:"accessToken"`
	RefreshToken []byte           `json:"refreshToken"`
	ExpiresAt    pgtype.Timestamp `json:"expiresAt"`
}

type GithubAppOAuthToUser struct {
	A pgtype.UUID `json:"A"`
	B pgtype.UUID `json:"B"`
}

type GithubPullRequest struct {
	ID                    pgtype.UUID      `json:"id"`
	CreatedAt             pgtype.Timestamp `json:"createdAt"`
	UpdatedAt             pgtype.Timestamp `json:"updatedAt"`
	DeletedAt             pgtype.Timestamp `json:"deletedAt"`
	TenantId              pgtype.UUID      `json:"tenantId"`
	RepositoryOwner       string           `json:"repositoryOwner"`
	RepositoryName        string           `json:"repositoryName"`
	PullRequestID         int32            `json:"pullRequestID"`
	PullRequestTitle      string           `json:"pullRequestTitle"`
	PullRequestNumber     int32            `json:"pullRequestNumber"`
	PullRequestHeadBranch string           `json:"pullRequestHeadBranch"`
	PullRequestBaseBranch string           `json:"pullRequestBaseBranch"`
	PullRequestState      string           `json:"pullRequestState"`
}

type GithubPullRequestComment struct {
	ID            pgtype.UUID      `json:"id"`
	CreatedAt     pgtype.Timestamp `json:"createdAt"`
	UpdatedAt     pgtype.Timestamp `json:"updatedAt"`
	DeletedAt     pgtype.Timestamp `json:"deletedAt"`
	TenantId      pgtype.UUID      `json:"tenantId"`
	PullRequestID pgtype.UUID      `json:"pullRequestID"`
	ModuleID      string           `json:"moduleID"`
	CommentID     int32            `json:"commentID"`
}

type GithubPullRequestToWorkflowRun struct {
	A pgtype.UUID `json:"A"`
	B pgtype.UUID `json:"B"`
}

type GithubWebhook struct {
	ID              pgtype.UUID      `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	DeletedAt       pgtype.Timestamp `json:"deletedAt"`
	TenantId        pgtype.UUID      `json:"tenantId"`
	RepositoryOwner string           `json:"repositoryOwner"`
	RepositoryName  string           `json:"repositoryName"`
	SigningSecret   []byte           `json:"signingSecret"`
}

type Job struct {
	ID                pgtype.UUID      `json:"id"`
	CreatedAt         pgtype.Timestamp `json:"createdAt"`
	UpdatedAt         pgtype.Timestamp `json:"updatedAt"`
	DeletedAt         pgtype.Timestamp `json:"deletedAt"`
	TenantId          pgtype.UUID      `json:"tenantId"`
	WorkflowVersionId pgtype.UUID      `json:"workflowVersionId"`
	Name              string           `json:"name"`
	Description       pgtype.Text      `json:"description"`
	Timeout           pgtype.Text      `json:"timeout"`
}

type JobRun struct {
	ID              pgtype.UUID      `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	DeletedAt       pgtype.Timestamp `json:"deletedAt"`
	TenantId        pgtype.UUID      `json:"tenantId"`
	JobId           pgtype.UUID      `json:"jobId"`
	TickerId        pgtype.UUID      `json:"tickerId"`
	Status          JobRunStatus     `json:"status"`
	Result          []byte           `json:"result"`
	StartedAt       pgtype.Timestamp `json:"startedAt"`
	FinishedAt      pgtype.Timestamp `json:"finishedAt"`
	TimeoutAt       pgtype.Timestamp `json:"timeoutAt"`
	CancelledAt     pgtype.Timestamp `json:"cancelledAt"`
	CancelledReason pgtype.Text      `json:"cancelledReason"`
	CancelledError  pgtype.Text      `json:"cancelledError"`
	WorkflowRunId   pgtype.UUID      `json:"workflowRunId"`
}

type JobRunLookupData struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	DeletedAt pgtype.Timestamp `json:"deletedAt"`
	JobRunId  pgtype.UUID      `json:"jobRunId"`
	TenantId  pgtype.UUID      `json:"tenantId"`
	Data      []byte           `json:"data"`
}

type Service struct {
	ID          pgtype.UUID      `json:"id"`
	CreatedAt   pgtype.Timestamp `json:"createdAt"`
	UpdatedAt   pgtype.Timestamp `json:"updatedAt"`
	DeletedAt   pgtype.Timestamp `json:"deletedAt"`
	Name        string           `json:"name"`
	Description pgtype.Text      `json:"description"`
	TenantId    pgtype.UUID      `json:"tenantId"`
}

type ServiceToWorker struct {
	A pgtype.UUID `json:"A"`
	B pgtype.UUID `json:"B"`
}

type Step struct {
	ID              pgtype.UUID      `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	DeletedAt       pgtype.Timestamp `json:"deletedAt"`
	ReadableId      pgtype.Text      `json:"readableId"`
	TenantId        pgtype.UUID      `json:"tenantId"`
	JobId           pgtype.UUID      `json:"jobId"`
	ActionId        string           `json:"actionId"`
	Timeout         pgtype.Text      `json:"timeout"`
	CustomUserData  []byte           `json:"customUserData"`
	Retries         int32            `json:"retries"`
	ScheduleTimeout string           `json:"scheduleTimeout"`
}

type StepOrder struct {
	A pgtype.UUID `json:"A"`
	B pgtype.UUID `json:"B"`
}

type StepRun struct {
	ID                pgtype.UUID      `json:"id"`
	CreatedAt         pgtype.Timestamp `json:"createdAt"`
	UpdatedAt         pgtype.Timestamp `json:"updatedAt"`
	DeletedAt         pgtype.Timestamp `json:"deletedAt"`
	TenantId          pgtype.UUID      `json:"tenantId"`
	JobRunId          pgtype.UUID      `json:"jobRunId"`
	StepId            pgtype.UUID      `json:"stepId"`
	Order             int64            `json:"order"`
	WorkerId          pgtype.UUID      `json:"workerId"`
	TickerId          pgtype.UUID      `json:"tickerId"`
	Status            StepRunStatus    `json:"status"`
	Input             []byte           `json:"input"`
	Output            []byte           `json:"output"`
	RequeueAfter      pgtype.Timestamp `json:"requeueAfter"`
	ScheduleTimeoutAt pgtype.Timestamp `json:"scheduleTimeoutAt"`
	Error             pgtype.Text      `json:"error"`
	StartedAt         pgtype.Timestamp `json:"startedAt"`
	FinishedAt        pgtype.Timestamp `json:"finishedAt"`
	TimeoutAt         pgtype.Timestamp `json:"timeoutAt"`
	CancelledAt       pgtype.Timestamp `json:"cancelledAt"`
	CancelledReason   pgtype.Text      `json:"cancelledReason"`
	CancelledError    pgtype.Text      `json:"cancelledError"`
	InputSchema       []byte           `json:"inputSchema"`
	CallerFiles       []byte           `json:"callerFiles"`
	GitRepoBranch     pgtype.Text      `json:"gitRepoBranch"`
	RetryCount        int32            `json:"retryCount"`
}

type StepRunOrder struct {
	A pgtype.UUID `json:"A"`
	B pgtype.UUID `json:"B"`
}

type StepRunResultArchive struct {
	ID              pgtype.UUID      `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	DeletedAt       pgtype.Timestamp `json:"deletedAt"`
	StepRunId       pgtype.UUID      `json:"stepRunId"`
	Order           int64            `json:"order"`
	Input           []byte           `json:"input"`
	Output          []byte           `json:"output"`
	Error           pgtype.Text      `json:"error"`
	StartedAt       pgtype.Timestamp `json:"startedAt"`
	FinishedAt      pgtype.Timestamp `json:"finishedAt"`
	TimeoutAt       pgtype.Timestamp `json:"timeoutAt"`
	CancelledAt     pgtype.Timestamp `json:"cancelledAt"`
	CancelledReason pgtype.Text      `json:"cancelledReason"`
	CancelledError  pgtype.Text      `json:"cancelledError"`
}

type Tenant struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	DeletedAt pgtype.Timestamp `json:"deletedAt"`
	Name      string           `json:"name"`
	Slug      string           `json:"slug"`
}

type TenantInviteLink struct {
	ID           pgtype.UUID      `json:"id"`
	CreatedAt    pgtype.Timestamp `json:"createdAt"`
	UpdatedAt    pgtype.Timestamp `json:"updatedAt"`
	TenantId     pgtype.UUID      `json:"tenantId"`
	InviterEmail string           `json:"inviterEmail"`
	InviteeEmail string           `json:"inviteeEmail"`
	Expires      pgtype.Timestamp `json:"expires"`
	Status       InviteLinkStatus `json:"status"`
	Role         TenantMemberRole `json:"role"`
}

type TenantMember struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	TenantId  pgtype.UUID      `json:"tenantId"`
	UserId    pgtype.UUID      `json:"userId"`
	Role      TenantMemberRole `json:"role"`
}

type TenantVcsProvider struct {
	ID          pgtype.UUID      `json:"id"`
	CreatedAt   pgtype.Timestamp `json:"createdAt"`
	UpdatedAt   pgtype.Timestamp `json:"updatedAt"`
	DeletedAt   pgtype.Timestamp `json:"deletedAt"`
	TenantId    pgtype.UUID      `json:"tenantId"`
	VcsProvider VcsProvider      `json:"vcsProvider"`
	Config      []byte           `json:"config"`
}

type Ticker struct {
	ID              pgtype.UUID      `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	LastHeartbeatAt pgtype.Timestamp `json:"lastHeartbeatAt"`
	IsActive        bool             `json:"isActive"`
}

type User struct {
	ID            pgtype.UUID      `json:"id"`
	CreatedAt     pgtype.Timestamp `json:"createdAt"`
	UpdatedAt     pgtype.Timestamp `json:"updatedAt"`
	DeletedAt     pgtype.Timestamp `json:"deletedAt"`
	Email         string           `json:"email"`
	EmailVerified bool             `json:"emailVerified"`
	Name          pgtype.Text      `json:"name"`
}

type UserOAuth struct {
	ID             pgtype.UUID      `json:"id"`
	CreatedAt      pgtype.Timestamp `json:"createdAt"`
	UpdatedAt      pgtype.Timestamp `json:"updatedAt"`
	UserId         pgtype.UUID      `json:"userId"`
	Provider       string           `json:"provider"`
	ProviderUserId string           `json:"providerUserId"`
	ExpiresAt      pgtype.Timestamp `json:"expiresAt"`
	AccessToken    []byte           `json:"accessToken"`
	RefreshToken   []byte           `json:"refreshToken"`
}

type UserPassword struct {
	Hash   string      `json:"hash"`
	UserId pgtype.UUID `json:"userId"`
}

type UserSession struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	UserId    pgtype.UUID      `json:"userId"`
	Data      []byte           `json:"data"`
	ExpiresAt pgtype.Timestamp `json:"expiresAt"`
}

type Worker struct {
	ID              pgtype.UUID      `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	DeletedAt       pgtype.Timestamp `json:"deletedAt"`
	TenantId        pgtype.UUID      `json:"tenantId"`
	LastHeartbeatAt pgtype.Timestamp `json:"lastHeartbeatAt"`
	Name            string           `json:"name"`
	Status          WorkerStatus     `json:"status"`
	DispatcherId    pgtype.UUID      `json:"dispatcherId"`
	MaxRuns         pgtype.Int4      `json:"maxRuns"`
}

type Workflow struct {
	ID          pgtype.UUID      `json:"id"`
	CreatedAt   pgtype.Timestamp `json:"createdAt"`
	UpdatedAt   pgtype.Timestamp `json:"updatedAt"`
	DeletedAt   pgtype.Timestamp `json:"deletedAt"`
	TenantId    pgtype.UUID      `json:"tenantId"`
	Name        string           `json:"name"`
	Description pgtype.Text      `json:"description"`
}

type WorkflowConcurrency struct {
	ID                    pgtype.UUID              `json:"id"`
	CreatedAt             pgtype.Timestamp         `json:"createdAt"`
	UpdatedAt             pgtype.Timestamp         `json:"updatedAt"`
	WorkflowVersionId     pgtype.UUID              `json:"workflowVersionId"`
	GetConcurrencyGroupId pgtype.UUID              `json:"getConcurrencyGroupId"`
	MaxRuns               int32                    `json:"maxRuns"`
	LimitStrategy         ConcurrencyLimitStrategy `json:"limitStrategy"`
}

type WorkflowDeploymentConfig struct {
	ID                      pgtype.UUID      `json:"id"`
	CreatedAt               pgtype.Timestamp `json:"createdAt"`
	UpdatedAt               pgtype.Timestamp `json:"updatedAt"`
	DeletedAt               pgtype.Timestamp `json:"deletedAt"`
	WorkflowId              pgtype.UUID      `json:"workflowId"`
	GitRepoName             string           `json:"gitRepoName"`
	GitRepoOwner            string           `json:"gitRepoOwner"`
	GitRepoBranch           string           `json:"gitRepoBranch"`
	GithubAppInstallationId pgtype.UUID      `json:"githubAppInstallationId"`
}

type WorkflowRun struct {
	CreatedAt          pgtype.Timestamp  `json:"createdAt"`
	UpdatedAt          pgtype.Timestamp  `json:"updatedAt"`
	DeletedAt          pgtype.Timestamp  `json:"deletedAt"`
	TenantId           pgtype.UUID       `json:"tenantId"`
	WorkflowVersionId  pgtype.UUID       `json:"workflowVersionId"`
	Status             WorkflowRunStatus `json:"status"`
	Error              pgtype.Text       `json:"error"`
	StartedAt          pgtype.Timestamp  `json:"startedAt"`
	FinishedAt         pgtype.Timestamp  `json:"finishedAt"`
	ConcurrencyGroupId pgtype.Text       `json:"concurrencyGroupId"`
	DisplayName        pgtype.Text       `json:"displayName"`
	ID                 pgtype.UUID       `json:"id"`
	GitRepoBranch      pgtype.Text       `json:"gitRepoBranch"`
}

type WorkflowRunTriggeredBy struct {
	ID           pgtype.UUID      `json:"id"`
	CreatedAt    pgtype.Timestamp `json:"createdAt"`
	UpdatedAt    pgtype.Timestamp `json:"updatedAt"`
	DeletedAt    pgtype.Timestamp `json:"deletedAt"`
	TenantId     pgtype.UUID      `json:"tenantId"`
	EventId      pgtype.UUID      `json:"eventId"`
	CronParentId pgtype.UUID      `json:"cronParentId"`
	CronSchedule pgtype.Text      `json:"cronSchedule"`
	ScheduledId  pgtype.UUID      `json:"scheduledId"`
	Input        []byte           `json:"input"`
	ParentId     pgtype.UUID      `json:"parentId"`
}

type WorkflowTag struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"createdAt"`
	UpdatedAt pgtype.Timestamp `json:"updatedAt"`
	TenantId  pgtype.UUID      `json:"tenantId"`
	Name      string           `json:"name"`
	Color     string           `json:"color"`
}

type WorkflowToWorkflowTag struct {
	A pgtype.UUID `json:"A"`
	B pgtype.UUID `json:"B"`
}

type WorkflowTriggerCronRef struct {
	ParentId pgtype.UUID `json:"parentId"`
	Cron     string      `json:"cron"`
	TickerId pgtype.UUID `json:"tickerId"`
	Input    []byte      `json:"input"`
}

type WorkflowTriggerEventRef struct {
	ParentId pgtype.UUID `json:"parentId"`
	EventKey string      `json:"eventKey"`
}

type WorkflowTriggerScheduledRef struct {
	ID        pgtype.UUID      `json:"id"`
	ParentId  pgtype.UUID      `json:"parentId"`
	TriggerAt pgtype.Timestamp `json:"triggerAt"`
	TickerId  pgtype.UUID      `json:"tickerId"`
	Input     []byte           `json:"input"`
}

type WorkflowTriggers struct {
	ID                pgtype.UUID      `json:"id"`
	CreatedAt         pgtype.Timestamp `json:"createdAt"`
	UpdatedAt         pgtype.Timestamp `json:"updatedAt"`
	DeletedAt         pgtype.Timestamp `json:"deletedAt"`
	WorkflowVersionId pgtype.UUID      `json:"workflowVersionId"`
	TenantId          pgtype.UUID      `json:"tenantId"`
}

type WorkflowVersion struct {
	ID              pgtype.UUID      `json:"id"`
	CreatedAt       pgtype.Timestamp `json:"createdAt"`
	UpdatedAt       pgtype.Timestamp `json:"updatedAt"`
	DeletedAt       pgtype.Timestamp `json:"deletedAt"`
	Version         pgtype.Text      `json:"version"`
	Order           int64            `json:"order"`
	WorkflowId      pgtype.UUID      `json:"workflowId"`
	Checksum        string           `json:"checksum"`
	ScheduleTimeout string           `json:"scheduleTimeout"`
}
