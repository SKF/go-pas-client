package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	dd_http "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"

	pas "github.com/SKF/go-pas-client"
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

	threshold, err := client.GetThreshold(context.Background(), nodeID)
	if err != nil {
		panic(err)
	}

	buf, err := json.MarshalIndent(threshold, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf))
}
