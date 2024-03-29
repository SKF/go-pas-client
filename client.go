package client

import (
	"context"
	"fmt"

	internal_models "github.com/SKF/go-pas-client/internal/models"
	"github.com/SKF/go-pas-client/models"
	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/stages"
	"github.com/SKF/go-utility/v2/uuid"
)

type API interface {
	GetThreshold(context.Context, uuid.UUID) (models.Threshold, error)
	SetThreshold(context.Context, uuid.UUID, models.Threshold) error
	PatchThreshold(context.Context, uuid.UUID, models.Patch) (models.Threshold, error)
	GetAlarmStatus(context.Context, uuid.UUID) (models.AlarmStatus, error)
	SetExternalAlarmStatus(context.Context, uuid.UUID, models.ExternalAlarmStatus) error
	UpdateAlarmStatus(context.Context, uuid.UUID, *models.Measurement) error
}

type Client struct {
	*rest.Client
}

var _ API = &Client{Client: nil}

func WithStage(stage string) rest.Option {
	if stage == stages.StageProd {
		return rest.WithBaseURL("https://api.point-alarm-status.iot.enlight.skf.com")
	}

	return rest.WithBaseURL(fmt.Sprintf("https://api.point-alarm-status.%s.iot.enlight.skf.com", stage))
}

func New(opts ...rest.Option) *Client {
	restClient := rest.NewClient(
		append([]rest.Option{
			// Defaults to production stage if no option is supplied
			WithStage(stages.StageProd),
			rest.WithProblemDecoder(&ProblemDecoder{}),
		}, opts...)...,
	)

	return &Client{restClient}
}

func (c *Client) GetThreshold(ctx context.Context, nodeID uuid.UUID) (models.Threshold, error) {
	request := rest.Get("v1/point-alarm-threshold/{nodeId}").
		Assign("nodeId", nodeID).
		SetHeader("Accept", "application/json")

	var response internal_models.ModelsGetPointAlarmThresholdResponse

	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.Threshold{}, fmt.Errorf("getting threshold failed: %w", err)
	}

	threshold := models.Threshold{} //nolint:exhaustruct

	if err := threshold.FromInternal(response); err != nil {
		return models.Threshold{}, fmt.Errorf("converting threshold failed: %w", err)
	}

	return threshold, nil
}

func (c *Client) SetThreshold(ctx context.Context, nodeID uuid.UUID, threshold models.Threshold) error {
	request := rest.Put("v1/point-alarm-threshold/{nodeId}").
		Assign("nodeId", nodeID).
		WithJSONPayload(threshold.ToInternal()).
		SetHeader("Accept", "application/json")

	if _, err := c.Do(ctx, request); err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	return nil
}

func (c *Client) PatchThreshold(ctx context.Context, nodeID uuid.UUID, patch models.Patch) (models.Threshold, error) {
	request := rest.Patch("v1/point-alarm-threshold/{nodeId}").
		Assign("nodeId", nodeID).
		WithJSONPayload(patch).
		SetHeader("Content-Type", "application/json-patch+json").
		SetHeader("Accept", "application/json")

	var response internal_models.ModelsGetPointAlarmThresholdResponse

	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.Threshold{}, fmt.Errorf("patching threshold failed: %w", err)
	}

	threshold := models.Threshold{}

	if err := threshold.FromInternal(response); err != nil {
		return models.Threshold{}, fmt.Errorf("converting threshold failed: %w", err)
	}

	return threshold, nil
}

func (c *Client) GetAlarmStatus(ctx context.Context, nodeID uuid.UUID) (alarmStatus models.AlarmStatus, err error) {
	request := rest.Get("v1/alarm-status/{nodeId}").
		Assign("nodeId", nodeID).
		SetHeader("Accept", "application/json")

	var response internal_models.ModelsGetAlarmStatusResponse

	if err = c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.AlarmStatus{}, fmt.Errorf("getting alarm status failed: %w", err)
	}

	alarmStatus.FromInternal(response)

	return
}

func (c *Client) UpdateAlarmStatus(
	ctx context.Context,
	nodeID uuid.UUID,
	measurement *models.Measurement,
) (err error) {
	request := rest.Put("v1/alarm-status/{nodeId}").
		Assign("nodeId", nodeID).
		SetHeader("Accept", "application/json")

	if measurement != nil {
		request = request.WithJSONPayload(measurement.ToInternal())
	}

	_, err = c.Do(ctx, request)

	return
}

func (c *Client) SetExternalAlarmStatus(
	ctx context.Context,
	nodeID uuid.UUID,
	status models.ExternalAlarmStatus,
) (err error) {
	payload := status.ToSetRequest()

	request := rest.Put("v1/alarm-status/{nodeId}/status/external").
		Assign("nodeId", nodeID).
		WithJSONPayload(payload).
		SetHeader("Accept", "application/json")

	_, err = c.Do(ctx, request)

	return
}
