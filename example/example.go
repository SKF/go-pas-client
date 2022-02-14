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
	return auth.RawToken(os.Getenv("ENLIGHT_TOKEN")), nil
}

func main() {
	client := pas.New(
		pas.WithStage(stages.StageSandbox),
		client.WithDatadogTracing(dd_http.RTWithServiceName(serviceName)),
		client.WithTokenProvider(&tokenProvider{}),
	)

	nodeID := uuid.UUID(os.Args[1])

	if err := nodeID.Validate(); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	setThreshold(ctx, client, nodeID)

	fmt.Println("The treshold:")
	getThreshold(ctx, client, nodeID)

	fmt.Println("The patched treshold:")
	patchThreshold(ctx, client, nodeID)

	fmt.Println("The alarm status:")
	getAlarmStatus(ctx, client, nodeID)

	setExternalAlarmStatus(ctx, client, nodeID)

	updateAlarmStatus(ctx, client, nodeID)

	fmt.Println("The alarm status:")
	getAlarmStatus(ctx, client, nodeID)
}

func print(data interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	encoder.Encode(data)
	fmt.Println()
}

func setThreshold(ctx context.Context, client *pas.Client, nodeID uuid.UUID) {
	var (
		outerHigh = 70.0
		innerHigh = 50.0
		innerLow  = 20.0
		outerLow  = 10.0
	)

	err := client.SetThreshold(ctx, nodeID, models.Threshold{
		ThresholdType: models.ThresholdTypeOverallOutOfWindow,
		Overall: &models.Overall{
			Unit:      "C",
			OuterHigh: &outerHigh,
			InnerHigh: &innerHigh,
			InnerLow:  &innerLow,
			OuterLow:  &outerLow,
		},
	})
	if err != nil {
		panic(err)
	}
}

func patchThreshold(ctx context.Context, client *pas.Client, nodeID uuid.UUID) {
	patch := models.Patch{
		{Type: "test", Path: "/overall/outerHigh", Value: 70},
		{Type: "replace", Path: "/overall/outerHigh", Value: 80},
	}

	threshold, err := client.PatchThreshold(ctx, nodeID, patch)
	if err != nil {
		panic(err)
	}

	print(threshold)
}

func getThreshold(ctx context.Context, client *pas.Client, nodeID uuid.UUID) {
	threshold, err := client.GetThreshold(ctx, nodeID)
	if err != nil {
		panic(err)
	}

	print(threshold)
}

func getAlarmStatus(ctx context.Context, client *pas.Client, nodeID uuid.UUID) {
	alarmStatus, err := client.GetAlarmStatus(ctx, nodeID)
	if err != nil {
		panic(err)
	}

	print(alarmStatus)
}

func updateAlarmStatus(ctx context.Context, client *pas.Client, nodeID uuid.UUID) {
	createdAt := time.Now().Add(-time.Minute).UTC()
	measurement := models.Measurement{
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
			"source": "unit-test",
		},
	}

	err := client.UpdateAlarmStatus(ctx, nodeID, measurement)
	if err != nil {
		panic(err)
	}
}

func setExternalAlarmStatus(ctx context.Context, client *pas.Client, nodeID uuid.UUID) {
	err := client.SetExternalAlarmStatus(ctx, nodeID, models.ExternalAlarmStatus{
		Status: models.AlarmStatusDanger,
	})
	if err != nil {
		panic(err)
	}
}
