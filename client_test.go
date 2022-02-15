package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
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

func i64p(i int64) *int64 {
	return &i
}

func f64p(f float64) *float64 {
	return &f
}

func stringp(s string) *string {
	return &s
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

func Test_GetAlarmStatus(t *testing.T) {
	var (
		now                   = time.UnixMilli(time.Now().UnixMilli()).UTC()
		triggeringMeasurement = strfmt.UUID(uuid.EmptyUUID.String())
	)

	given := internal_models.ModelsGetAlarmStatusResponse{
		UpdatedAt: now.UnixMilli(),
		Status:    i32p(4), // danger
		OverallAlarm: &internal_models.ModelsGetAlarmStatusResponseGeneric{
			Status:                i32p(2), // good
			TriggeringMeasurement: strfmt.UUID(uuid.EmptyUUID.String()),
		},
		RateOfChangeAlarm: &internal_models.ModelsGetAlarmStatusResponseGeneric{
			Status:                i32p(2), // good
			TriggeringMeasurement: strfmt.UUID(uuid.EmptyUUID.String()),
		},
		InspectionAlarm: &internal_models.ModelsGetAlarmStatusResponseGeneric{
			Status:                i32p(2), // good
			TriggeringMeasurement: strfmt.UUID(uuid.EmptyUUID.String()),
		},
		ExternalAlarm: &internal_models.ModelsGetAlarmStatusResponseExternal{
			Status: i32p(2), // good
		},
		BandAlarms: []*internal_models.ModelsGetAlarmStatusResponseBandAlarm{
			{
				Label:                 "10x RPM",
				Status:                i32p(3), // alert
				TriggeringMeasurement: strfmt.UUID(uuid.EmptyUUID.String()),
				MinFrequency: &internal_models.ModelsGetAlarmStatusResponseFrequency{
					ValueType: i32p(1), // fixed
					Value:     f64p(100),
				},
				MaxFrequency: &internal_models.ModelsGetAlarmStatusResponseFrequency{
					ValueType: i32p(1), // fixed
					Value:     f64p(500),
				},
				CalculatedOverall: &internal_models.ModelsBandCalculatedOverall{
					Unit:  "gE",
					Value: f64p(3.5),
				},
			},
			{
				Label:                 "12 TOOTH SPROCKET",
				Status:                i32p(2), // good
				TriggeringMeasurement: strfmt.UUID(uuid.EmptyUUID.String()),
				MinFrequency: &internal_models.ModelsGetAlarmStatusResponseFrequency{
					ValueType: i32p(2), // speed multiple
					Value:     f64p(1.2),
				},
				MaxFrequency: &internal_models.ModelsGetAlarmStatusResponseFrequency{
					ValueType: i32p(2), // speed multiple
					Value:     f64p(1.5),
				},
				CalculatedOverall: &internal_models.ModelsBandCalculatedOverall{
					Unit:  "gE",
					Value: f64p(3.5),
				},
			},
		},
		HalAlarms: []*internal_models.ModelsGetAlarmStatusResponseHALAlarm{
			{
				Label:                 stringp("10x RPM"),
				Status:                i32p(4), // danger
				TriggeringMeasurement: &triggeringMeasurement,
				HalIndex:              f64p(1.22),
				FaultFrequency:        f64p(122),
				RpmFactor:             f64p(10),
				NumberOfHarmonicsUsed: i64p(15),
			},
			{
				Label:                 stringp("12 TOOTH SPROCKET"),
				Status:                i32p(1), // no data
				TriggeringMeasurement: &triggeringMeasurement,
				FaultFrequency:        f64p(122),
				RpmFactor:             f64p(12),
				ErrorDescription:      stringp("only peaks"),
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

	expected := models.AlarmStatus{
		UpdatedAt: now,
		Status:    models.AlarmStatusDanger,
		Overall: &models.GenericAlarmStatus{
			Status:                models.AlarmStatusGood,
			TriggeringMeasurement: uuid.EmptyUUID,
		},
		RateOfChange: &models.GenericAlarmStatus{
			Status:                models.AlarmStatusGood,
			TriggeringMeasurement: uuid.EmptyUUID,
		},
		Inspection: &models.GenericAlarmStatus{
			Status:                models.AlarmStatusGood,
			TriggeringMeasurement: uuid.EmptyUUID,
		},
		External: &models.ExternalAlarmStatus{
			Status: models.AlarmStatusGood,
		},
		Band: []models.BandAlarmStatus{
			{
				Label: "10x RPM",
				GenericAlarmStatus: models.GenericAlarmStatus{
					Status:                models.AlarmStatusAlert,
					TriggeringMeasurement: uuid.EmptyUUID,
				},
				MinFrequency: models.BandAlarmFrequency{
					ValueType: models.BandAlarmFrequencyFixed,
					Value:     100,
				},
				MaxFrequency: models.BandAlarmFrequency{
					ValueType: models.BandAlarmFrequencyFixed,
					Value:     500,
				},
				CalculatedOverall: &models.BandAlarmStatusCalculatedOverall{
					Unit:  "gE",
					Value: 3.5,
				},
			},
			{
				Label: "12 TOOTH SPROCKET",
				GenericAlarmStatus: models.GenericAlarmStatus{
					Status:                models.AlarmStatusGood,
					TriggeringMeasurement: uuid.EmptyUUID,
				},
				MinFrequency: models.BandAlarmFrequency{
					ValueType: models.BandAlarmFrequencySpeedMultiple,
					Value:     1.2,
				},
				MaxFrequency: models.BandAlarmFrequency{
					ValueType: models.BandAlarmFrequencySpeedMultiple,
					Value:     1.5,
				},
				CalculatedOverall: &models.BandAlarmStatusCalculatedOverall{
					Unit:  "gE",
					Value: 3.5,
				},
			},
		},
		HAL: []models.HALAlarmStatus{
			{
				Label: "10x RPM",
				GenericAlarmStatus: models.GenericAlarmStatus{
					Status:                models.AlarmStatusDanger,
					TriggeringMeasurement: uuid.EmptyUUID,
				},
				HALIndex:              f64p(1.22),
				FaultFrequency:        f64p(122),
				RPMFactor:             f64p(10),
				NumberOfHarmonicsUsed: i64p(15),
			},
			{
				Label: "12 TOOTH SPROCKET",
				GenericAlarmStatus: models.GenericAlarmStatus{
					Status:                models.AlarmStatusNoData,
					TriggeringMeasurement: uuid.EmptyUUID,
				},
				FaultFrequency:   f64p(122),
				RPMFactor:        f64p(12),
				ErrorDescription: stringp("only peaks"),
			},
		},
	}

	actual, err := client.GetAlarmStatus(context.TODO(), uuid.EmptyUUID)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

func Test_GetAlarmStatus_ErrorResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := New(rest.WithBaseURL(server.URL))

	_, err := client.GetAlarmStatus(context.TODO(), uuid.EmptyUUID)

	assert.Error(t, err)
}

func Test_UpdateAlarmStatus(t *testing.T) {
	var (
		now                   = time.UnixMilli(time.Now().UnixMilli()).UTC()
		measurementID         = uuid.EmptyUUID
		expectedCreatedAt     = strfmt.DateTime(now)
		expectedMeasurementID = strfmt.UUID(measurementID.String())
		expectedContentType   = string(models.ContentTypeDataPoint)

		given = models.Measurement{
			CreatedAt:     now,
			MeasurementID: measurementID,
			ContentType:   models.ContentTypeDataPoint,
			DataPoint: &models.DataPoint{
				Coordinate: models.Coordinate{
					X: float64(now.UnixMilli()),
					Y: 10.0,
				},
				XUnit: "ms",
				YUnit: "gE",
			},
			Tags: map[string]interface{}{
				"source": "unit-test",
			},
		}
		expected = internal_models.ModelsUpdateAlarmStatusRequest{
			CreatedAt:     &expectedCreatedAt,
			MeasurementID: &expectedMeasurementID,
			ContentType:   &expectedContentType,
			DataPoint: &internal_models.ModelsDataPoint{
				Coordinate: &internal_models.ModelsCoordinate{
					X: f64p(float64(now.UnixMilli())),
					Y: f64p(10.0),
				},
				XUnit: stringp("ms"),
				YUnit: stringp("gE"),
			},
			Tags: map[string]interface{}{
				"source": "unit-test",
			},
		}
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var actual internal_models.ModelsUpdateAlarmStatusRequest

		err := json.NewDecoder(r.Body).Decode(&actual)
		require.NoError(t, err)

		assert.Equal(t, expected, actual)

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := New(rest.WithBaseURL(server.URL))

	err := client.UpdateAlarmStatus(context.TODO(), uuid.EmptyUUID, given)
	require.NoError(t, err)
}

func Test_SetExternalAlarmStatus(t *testing.T) {
	setBy := uuid.EmptyUUID

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request internal_models.ModelsSetExternalAlarmStatusRequest

		err := json.NewDecoder(r.Body).Decode(&request)
		require.NoError(t, err)

		if assert.NotNil(t, request.Status) {
			assert.Equal(t, int32(models.AlarmStatusAlert), *request.Status)
		}

		if assert.NotNil(t, request.SetBy) {
			assert.Equal(t, strfmt.UUID(uuid.EmptyUUID.String()), *request.SetBy)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := New(rest.WithBaseURL(server.URL))

	err := client.SetExternalAlarmStatus(context.TODO(), uuid.EmptyUUID, models.ExternalAlarmStatus{
		Status: models.AlarmStatusAlert,
		SetBy:  &setBy,
	})

	assert.NoError(t, err)
}
