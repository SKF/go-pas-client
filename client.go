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

	threshold := models.Threshold{} // nolint:exhaustivestruct

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
