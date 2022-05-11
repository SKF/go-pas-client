package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	dd_http "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"

	pas "github.com/SKF/go-pas-client"
	"github.com/SKF/go-pas-client/models"
	"github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-rest-utility/client/auth"
	"github.com/SKF/go-utility/v2/stages"
	"github.com/SKF/go-utility/v2/uuid"
)

const serviceName = "example-service"

type tokenProvider struct{}

func (t *tokenProvider) GetRawToken(ctx context.Context) (auth.RawToken, error) {
	return auth.RawToken(mustGetEnv("TOKEN")), nil
}

type api struct {
	client *pas.Client
	nodeID uuid.UUID
}

func main() {
	a := &api{
		client: pas.New(
			pas.WithStage(stages.StageSandbox),
			client.WithDatadogTracing(dd_http.RTWithServiceName(serviceName)),
			client.WithTokenProvider(&tokenProvider{}),
		),
		nodeID: uuid.UUID(os.Args[1]),
	}

	if err := a.nodeID.Validate(); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.setThreshold(ctx); err != nil {
		panic(err)
	}

	threshold, err := a.getThreshold(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("The threshold:")
	dbg(threshold)

	threshold, err = a.patchThreshold(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("The patched threshold:")
	dbg(threshold)

	alarmStatus, err := a.getAlarmStatus(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("The alarm status:")
	dbg(alarmStatus)

	if err = a.updateAlarmStatus(ctx); err != nil {
		panic(err)
	}

	if err = a.setExternalAlarmStatus(ctx); err != nil {
		panic(err)
	}

	alarmStatus, err = a.getAlarmStatus(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("The alarm status:")
	dbg(alarmStatus)
}

func dbg(data interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	encoder.Encode(data)
	fmt.Println()
}

func (a *api) setThreshold(ctx context.Context) error {
	var (
		outerHigh = 70.0
		innerHigh = 50.0
		innerLow  = 20.0
		outerLow  = 10.0
	)

	return a.client.SetThreshold(ctx, a.nodeID, models.Threshold{
		ThresholdType: models.ThresholdTypeOverallOutOfWindow,
		Overall: &models.Overall{
			Unit:      "C",
			OuterHigh: &outerHigh,
			InnerHigh: &innerHigh,
			InnerLow:  &innerLow,
			OuterLow:  &outerLow,
		},
	})
}

func (a *api) getThreshold(ctx context.Context) (models.Threshold, error) {
	return a.client.GetThreshold(ctx, a.nodeID)
}

func (a *api) patchThreshold(ctx context.Context) (models.Threshold, error) {
	patch := models.Patch{
		{Type: "test", Path: "/overall/outerHigh", Value: 70},
		{Type: "replace", Path: "/overall/outerHigh", Value: 80},
	}

	return a.client.PatchThreshold(ctx, a.nodeID, patch)
}

func (a *api) getAlarmStatus(ctx context.Context) (models.AlarmStatus, error) {
	return a.client.GetAlarmStatus(ctx, a.nodeID)
}

func (a *api) updateAlarmStatus(ctx context.Context) error {
	var (
		createdAt   = time.Now().Add(-time.Minute).UTC()
		measurement = models.Measurement{
			CreatedAt:     createdAt,
			MeasurementID: uuid.EmptyUUID,
			ContentType:   models.ContentTypeDataPoint,
			DataPoint: &models.DataPoint{
				Coordinate: models.Coordinate{
					X: float64(createdAt.UnixMilli()),
					Y: 65,
				},
				XUnit: "ms",
				YUnit: "C",
			},
			Tags: map[string]interface{}{
				"source": "client-example",
			},
		}
	)

	return a.client.UpdateAlarmStatus(ctx, a.nodeID, &measurement)
}

func (a *api) setExternalAlarmStatus(ctx context.Context) error {
	return a.client.SetExternalAlarmStatus(ctx, a.nodeID, models.ExternalAlarmStatus{
		Status: models.AlarmStatusDanger,
	})
}

func mustGetEnv(key string) string {
	value, found := os.LookupEnv(key)
	if !found {
		panic(fmt.Errorf("environment variable %q is not set", key))
	}

	return value
}
