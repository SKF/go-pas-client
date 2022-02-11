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

	set(ctx, client, nodeID)

	fmt.Println("The treshold:")
	get(ctx, client, nodeID)

	fmt.Println("The patched treshold:")
	patch(ctx, client, nodeID)
}

func print(threshold models.Threshold) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	encoder.Encode(threshold)
	fmt.Println()
}

func set(ctx context.Context, client *pas.Client, nodeID uuid.UUID) {
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

func patch(ctx context.Context, client *pas.Client, nodeID uuid.UUID) {
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

func get(ctx context.Context, client *pas.Client, nodeID uuid.UUID) {
	threshold, err := client.GetThreshold(ctx, nodeID)
	if err != nil {
		panic(err)
	}

	print(threshold)
}
