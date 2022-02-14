package models

import (
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"

	models "github.com/SKF/go-pas-client/internal/models"
	"github.com/SKF/go-utility/v2/uuid"
)

func strfmtDateTimep(d strfmt.DateTime) *strfmt.DateTime {
	return &d
}

func strfmtUUIDp(u strfmt.UUID) *strfmt.UUID {
	return &u
}

func Test_Measurement_ToInternal(t *testing.T) {
	var (
		now           = time.UnixMilli(time.Now().UnixMilli()).UTC()
		measurementID = uuid.EmptyUUID
	)

	tests := []struct {
		given    *Measurement
		expected models.ModelsUpdateAlarmStatusRequest
	}{
		{
			given: &Measurement{
				CreatedAt:     now,
				MeasurementID: measurementID,
				ContentType:   ContentTypeDataPoint,
				DataPoint: &DataPoint{
					Coordinate: Coordinate{
						X: float64(now.UnixMilli()),
						Y: 10.0,
					},
					XUnit: "ms",
					YUnit: "gE",
				},
				Tags: map[string]interface{}{
					"source": "unit-test",
				},
			},
			expected: models.ModelsUpdateAlarmStatusRequest{
				CreatedAt:     strfmtDateTimep(strfmt.DateTime(now)),
				MeasurementID: strfmtUUIDp(strfmt.UUID(measurementID.String())),
				ContentType:   stringp("DATA_POINT"),
				DataPoint: &models.ModelsDataPoint{
					Coordinate: &models.ModelsCoordinate{
						X: f64p(float64(now.UnixMilli())),
						Y: f64p(10.0),
					},
					XUnit: stringp("ms"),
					YUnit: stringp("gE"),
				},
				Tags: map[string]interface{}{
					"source": "unit-test",
				},
			},
		},
		{
			given: &Measurement{
				CreatedAt:     now,
				MeasurementID: measurementID,
				ContentType:   ContentTypeSpectrum,
				Spectrum: &Spectrum{
					XUnit: "ms",
					YUnit: "gE",
					Speed: 1780.02,
				},
				Tags: map[string]interface{}{
					"source": "unit-test",
				},
			},
			expected: models.ModelsUpdateAlarmStatusRequest{
				CreatedAt:     strfmtDateTimep(strfmt.DateTime(now)),
				MeasurementID: strfmtUUIDp(strfmt.UUID(measurementID.String())),
				ContentType:   stringp("SPECTRUM"),
				Spectrum: &models.ModelsSpectrum{
					XUnit: stringp("ms"),
					YUnit: stringp("gE"),
					Speed: f64p(1780.02),
				},
				Tags: map[string]interface{}{
					"source": "unit-test",
				},
			},
		},
		{
			given: &Measurement{
				CreatedAt:     now,
				MeasurementID: measurementID,
				ContentType:   ContentTypeQuestionAnswers,
				QuestionAnswers: []string{
					"good",
				},
				Tags: map[string]interface{}{
					"source": "unit-test",
				},
			},
			expected: models.ModelsUpdateAlarmStatusRequest{
				CreatedAt:     strfmtDateTimep(strfmt.DateTime(now)),
				MeasurementID: strfmtUUIDp(strfmt.UUID(measurementID.String())),
				ContentType:   stringp("QUESTION_ANSWERS"),
				QuestionAnswers: []string{
					"good",
				},
				Tags: map[string]interface{}{
					"source": "unit-test",
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := test.given.ToInternal()

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_Measurement_ToInternal_IsNil(t *testing.T) {
	assert.NotPanics(t, func() {
		var m *Measurement

		_ = m.ToInternal()
	})
}
