package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	internal_models "github.com/SKF/go-pas-client/internal/models"
	"github.com/SKF/go-pas-client/models"
	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/uuid"
)

func i32p(i int32) *int32 {
	return &i
}

func f64p(f float64) *float64 {
	return &f
}

func Test_BaseURL(t *testing.T) {
	c := New()

	require.NotNil(t, c.Client.BaseURL)
	assert.Equal(t, "api.point-alarm-status.iot.enlight.skf.com", c.Client.BaseURL.Host)

	c = New(WithStage("sandbox"))

	require.NotNil(t, c.Client.BaseURL)
	assert.Equal(t, "api.point-alarm-status.sandbox.iot.enlight.skf.com", c.Client.BaseURL.Host)
}

func Test_GetThreshold(t *testing.T) {
	given := internal_models.ModelsGetPointAlarmThresholdResponse{
		ThresholdType: i32p(2),
		Overall: &internal_models.ModelsOverall{
			Unit:      "gE",
			OuterHigh: f64p(8),
			InnerHigh: f64p(6),
			InnerLow:  f64p(2),
			OuterLow:  f64p(1),
		},
		RateOfChange: &internal_models.ModelsRateOfChange{
			Unit:      "gE",
			OuterHigh: f64p(10),
			InnerHigh: f64p(5),
			InnerLow:  f64p(-5),
			OuterLow:  f64p(-10),
		},
		BandAlarms: []*internal_models.ModelsBandAlarm{
			{
				Label: "10x RPM",
				MinFrequency: &internal_models.ModelsBandAlarmFrequency{
					ValueType: i32p(1),
					Value:     f64p(0),
				},
				MaxFrequency: &internal_models.ModelsBandAlarmFrequency{
					ValueType: i32p(1),
					Value:     f64p(2000),
				},
				OverallThreshold: &internal_models.ModelsBandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &internal_models.ModelsBandAlarmThreshold{
						ValueType: i32p(1),
						Value:     f64p(1),
					},
					UpperDanger: &internal_models.ModelsBandAlarmThreshold{
						ValueType: i32p(1),
						Value:     f64p(2),
					},
				},
			},
		},
		HalAlarms: []*internal_models.ModelsHALAlarm{
			{
				Label:        "10x RPM",
				HalAlarmType: "FREQUENCY",
				UpperAlert:   f64p(3.5),
				UpperDanger:  f64p(5),
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(given)
		require.NoError(t, err)
	}))
	defer server.Close()

	client := New(rest.WithBaseURL(server.URL))

	actual, err := client.GetThreshold(context.TODO(), uuid.EmptyUUID)
	require.NoError(t, err)

	expected := models.Threshold{
		ThresholdType: models.ThresholdTypeOverallOutOfWindow,
		Overall: &models.Overall{
			Unit:      "gE",
			OuterHigh: f64p(8),
			InnerHigh: f64p(6),
			InnerLow:  f64p(2),
			OuterLow:  f64p(1),
		},
		RateOfChange: &models.RateOfChange{
			Unit:      "gE",
			OuterHigh: f64p(10),
			InnerHigh: f64p(5),
			InnerLow:  f64p(-5),
			OuterLow:  f64p(-10),
		},
		BandAlarms: []models.BandAlarm{
			{
				Label: "10x RPM",
				MinFrequency: models.BandAlarmFrequency{
					ValueType: models.BandAlarmFrequencyFixed,
					Value:     0,
				},
				MaxFrequency: models.BandAlarmFrequency{
					ValueType: models.BandAlarmFrequencyFixed,
					Value:     2000,
				},
				OverallThreshold: &models.BandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &models.BandAlarmThreshold{
						ValueType: models.BandAlarmThresholdTypeAbsolute,
						Value:     1,
					},
					UpperDanger: &models.BandAlarmThreshold{
						ValueType: models.BandAlarmThresholdTypeAbsolute,
						Value:     2,
					},
				},
			},
		},
		HALAlarms: []models.HALAlarm{
			{
				Label:        "10x RPM",
				HALAlarmType: models.HALAlarmTypeFaultFrequency,
				UpperAlert:   f64p(3.5),
				UpperDanger:  f64p(5),
			},
		},
	}

	assert.Equal(t, expected, actual)
}

func Test_GetThreshold_ErrorResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := New(rest.WithBaseURL(server.URL))

	_, err := client.GetThreshold(context.TODO(), uuid.EmptyUUID)

	assert.Error(t, err)
}

func Test_GetThreshold_InvalidNodeIDResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{"nodeId": "boop"}`))
	}))
	defer server.Close()

	client := New(rest.WithBaseURL(server.URL))

	_, err := client.GetThreshold(context.TODO(), uuid.EmptyUUID)

	assert.Error(t, err)
}

func Test_SetThreshold(t *testing.T) {
	expected := internal_models.ModelsSetPointAlarmThresholdRequest{
		ThresholdType: i32p(2),
		Overall: &internal_models.ModelsOverall{
			Unit:      "gE",
			OuterHigh: f64p(8),
			InnerHigh: f64p(6),
			InnerLow:  f64p(2),
			OuterLow:  f64p(1),
		},
		RateOfChange: &internal_models.ModelsRateOfChange{
			Unit:      "gE",
			OuterHigh: f64p(10),
			InnerHigh: f64p(5),
			InnerLow:  f64p(-5),
			OuterLow:  f64p(-10),
		},
		BandAlarms: []*internal_models.ModelsBandAlarm{
			{
				Label: "10x RPM",
				MinFrequency: &internal_models.ModelsBandAlarmFrequency{
					ValueType: i32p(1),
					Value:     f64p(0),
				},
				MaxFrequency: &internal_models.ModelsBandAlarmFrequency{
					ValueType: i32p(1),
					Value:     f64p(2000),
				},
				OverallThreshold: &internal_models.ModelsBandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &internal_models.ModelsBandAlarmThreshold{
						ValueType: i32p(1),
						Value:     f64p(1),
					},
					UpperDanger: &internal_models.ModelsBandAlarmThreshold{
						ValueType: i32p(1),
						Value:     f64p(2),
					},
				},
			},
		},
		HalAlarms: []*internal_models.ModelsHALAlarm{
			{
				Label:        "10x RPM",
				HalAlarmType: "FREQUENCY",
				UpperAlert:   f64p(3.5),
				UpperDanger:  f64p(5),
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var actual internal_models.ModelsSetPointAlarmThresholdRequest

		err := json.NewDecoder(r.Body).Decode(&actual)
		require.NoError(t, err)

		assert.Equal(t, expected, actual)

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := New(rest.WithBaseURL(server.URL))

	given := models.Threshold{
		ThresholdType: models.ThresholdTypeOverallOutOfWindow,
		Overall: &models.Overall{
			Unit:      "gE",
			OuterHigh: f64p(8),
			InnerHigh: f64p(6),
			InnerLow:  f64p(2),
			OuterLow:  f64p(1),
		},
		RateOfChange: &models.RateOfChange{
			Unit:      "gE",
			OuterHigh: f64p(10),
			InnerHigh: f64p(5),
			InnerLow:  f64p(-5),
			OuterLow:  f64p(-10),
		},
		BandAlarms: []models.BandAlarm{
			{
				Label: "10x RPM",
				MinFrequency: models.BandAlarmFrequency{
					ValueType: models.BandAlarmFrequencyFixed,
					Value:     0,
				},
				MaxFrequency: models.BandAlarmFrequency{
					ValueType: models.BandAlarmFrequencyFixed,
					Value:     2000,
				},
				OverallThreshold: &models.BandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &models.BandAlarmThreshold{
						ValueType: models.BandAlarmThresholdTypeAbsolute,
						Value:     1,
					},
					UpperDanger: &models.BandAlarmThreshold{
						ValueType: models.BandAlarmThresholdTypeAbsolute,
						Value:     2,
					},
				},
			},
		},
		HALAlarms: []models.HALAlarm{
			{
				Label:        "10x RPM",
				HALAlarmType: models.HALAlarmTypeFaultFrequency,
				UpperAlert:   f64p(3.5),
				UpperDanger:  f64p(5),
			},
		},
	}

	err := client.SetThreshold(context.TODO(), uuid.EmptyUUID, given)

	assert.NoError(t, err)
}

func Test_SetThreshold_ErrorResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := New(rest.WithBaseURL(server.URL))

	err := client.SetThreshold(context.TODO(), uuid.EmptyUUID, models.Threshold{})

	assert.Error(t, err)
}

func Test_PatchThreshold(t *testing.T) {
	given := internal_models.ModelsGetPointAlarmThresholdResponse{
		ThresholdType: i32p(2),
		Overall: &internal_models.ModelsOverall{
			Unit:      "gE",
			OuterHigh: f64p(8),
			InnerHigh: f64p(6),
			InnerLow:  f64p(2),
			OuterLow:  f64p(1),
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if assert.Contains(t, r.Header, "Content-Type") {
			assert.Equal(t, []string{"application/json-patch+json"}, r.Header["Content-Type"])
		}

		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(given)
		require.NoError(t, err)
	}))
	defer server.Close()

	client := New(rest.WithBaseURL(server.URL))

	expected := models.Threshold{
		ThresholdType: models.ThresholdTypeOverallOutOfWindow,
		Overall: &models.Overall{
			Unit:      "gE",
			OuterHigh: f64p(8),
			InnerHigh: f64p(6),
			InnerLow:  f64p(2),
			OuterLow:  f64p(1),
		},
		BandAlarms: []models.BandAlarm{},
		HALAlarms:  []models.HALAlarm{},
	}

	actual, err := client.PatchThreshold(context.TODO(), uuid.EmptyUUID, models.Patch{})
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func Test_PatchThreshold_ErrorResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := New(rest.WithBaseURL(server.URL))

	_, err := client.PatchThreshold(context.TODO(), uuid.EmptyUUID, models.Patch{})

	assert.Error(t, err)
}

func Test_PatchThreshold_InvalidNodeIDResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{"nodeId": "boop"}`))
	}))
	defer server.Close()

	client := New(rest.WithBaseURL(server.URL))

	_, err := client.PatchThreshold(context.TODO(), uuid.EmptyUUID, models.Patch{})

	assert.Error(t, err)
}
