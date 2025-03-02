package ingestor

import (
	"context"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hatchet-dev/hatchet/internal/datautils"
	"github.com/hatchet-dev/hatchet/internal/msgqueue"
	"github.com/hatchet-dev/hatchet/internal/repository"
	"github.com/hatchet-dev/hatchet/internal/repository/prisma/dbsqlc"
	"github.com/hatchet-dev/hatchet/internal/repository/prisma/sqlchelpers"
	"github.com/hatchet-dev/hatchet/internal/services/ingestor/contracts"
	"github.com/hatchet-dev/hatchet/internal/services/shared/tasktypes"
)

func (i *IngestorImpl) Push(ctx context.Context, req *contracts.PushEventRequest) (*contracts.Event, error) {
	tenant := ctx.Value("tenant").(*dbsqlc.Tenant)

	tenantId := sqlchelpers.UUIDToStr(tenant.ID)

	event, err := i.IngestEvent(ctx, tenantId, req.Key, []byte(req.Payload))

	if err != nil {
		return nil, err
	}

	e, err := toEvent(event)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (i *IngestorImpl) ReplaySingleEvent(ctx context.Context, req *contracts.ReplayEventRequest) (*contracts.Event, error) {
	tenant := ctx.Value("tenant").(*dbsqlc.Tenant)

	tenantId := sqlchelpers.UUIDToStr(tenant.ID)

	oldEvent, err := i.eventRepository.GetEventForEngine(ctx, tenantId, req.EventId)

	if err != nil {
		return nil, err
	}

	newEvent, err := i.IngestReplayedEvent(ctx, tenantId, oldEvent)

	if err != nil {
		return nil, err
	}

	e, err := toEvent(newEvent)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (i *IngestorImpl) PutStreamEvent(ctx context.Context, req *contracts.PutStreamEventRequest) (*contracts.PutStreamEventResponse, error) {
	tenant := ctx.Value("tenant").(*dbsqlc.Tenant)

	tenantId := sqlchelpers.UUIDToStr(tenant.ID)

	var createdAt *time.Time

	if t := req.CreatedAt.AsTime().UTC(); !t.IsZero() {
		createdAt = &t
	}

	var metadata []byte

	if req.Metadata != "" {
		metadata = []byte(req.Metadata)
	}

	streamEvent, err := i.streamEventRepository.PutStreamEvent(ctx, tenantId, &repository.CreateStreamEventOpts{
		StepRunId: req.StepRunId,
		CreatedAt: createdAt,
		Message:   req.Message,
		Metadata:  metadata,
	})

	if err != nil {
		return nil, err
	}

	q, err := msgqueue.TenantEventConsumerQueue(tenantId)

	if err != nil {
		return nil, err
	}

	err = i.mq.AddMessage(context.Background(), q, streamEventToTask(streamEvent))

	if err != nil {
		return nil, err
	}

	return &contracts.PutStreamEventResponse{}, nil
}

func (i *IngestorImpl) PutLog(ctx context.Context, req *contracts.PutLogRequest) (*contracts.PutLogResponse, error) {
	tenant := ctx.Value("tenant").(*dbsqlc.Tenant)

	tenantId := sqlchelpers.UUIDToStr(tenant.ID)

	var createdAt *time.Time

	if t := req.CreatedAt.AsTime(); !t.IsZero() {
		createdAt = &t
	}

	var metadata []byte

	if req.Metadata != "" {
		metadata = []byte(req.Metadata)
	}

	_, err := i.logRepository.PutLog(ctx, tenantId, &repository.CreateLogLineOpts{
		StepRunId: req.StepRunId,
		CreatedAt: createdAt,
		Message:   req.Message,
		Level:     req.Level,
		Metadata:  metadata,
	})

	if err != nil {
		return nil, err
	}

	return &contracts.PutLogResponse{}, nil
}

func toEvent(e *dbsqlc.Event) (*contracts.Event, error) {
	tenantId := sqlchelpers.UUIDToStr(e.TenantId)
	eventId := sqlchelpers.UUIDToStr(e.ID)

	return &contracts.Event{
		TenantId:       tenantId,
		EventId:        eventId,
		Key:            e.Key,
		Payload:        string(e.Data),
		EventTimestamp: timestamppb.New(e.CreatedAt.Time),
	}, nil
}

func streamEventToTask(e *dbsqlc.StreamEvent) *msgqueue.Message {
	tenantId := sqlchelpers.UUIDToStr(e.TenantId)

	payloadTyped := tasktypes.StepRunStreamEventTaskPayload{
		StepRunId:     sqlchelpers.UUIDToStr(e.StepRunId),
		CreatedAt:     e.CreatedAt.Time.String(),
		StreamEventId: strconv.FormatInt(e.ID, 10),
	}

	payload, _ := datautils.ToJSONMap(payloadTyped)

	metadata, _ := datautils.ToJSONMap(tasktypes.StepRunStreamEventTaskMetadata{
		TenantId:      tenantId,
		StreamEventId: strconv.FormatInt(e.ID, 10),
	})

	return &msgqueue.Message{
		ID:       "step-run-stream-event",
		Payload:  payload,
		Metadata: metadata,
		Retries:  3,
	}
}
