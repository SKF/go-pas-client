package models

import (
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	models "github.com/SKF/go-pas-client/internal/models"
	"github.com/SKF/go-utility/v2/uuid"
)

func f64p(f float64) *float64 {
	return &f
}

func i32p(i int32) *int32 {
	return &i
}

func stringp(s string) *string {
	return &s
}

func Test_ThresholdFromInternal(t *testing.T) {
	t.Parallel()

	nodeID := strfmt.UUID(uuid.EmptyUUID.String())

	tests := []struct {
		given    models.ModelsGetPointAlarmThresholdResponse
		expected *Threshold
	}{
		{
			given: models.ModelsGetPointAlarmThresholdResponse{},
			expected: &Threshold{
				BandAlarms: []BandAlarm{},
				HALAlarms:  []HALAlarm{},
			},
		},
		{
			given: models.ModelsGetPointAlarmThresholdResponse{
				NodeID: &nodeID,
			},
			expected: &Threshold{
				NodeID:     uuid.EmptyUUID,
				BandAlarms: []BandAlarm{},
				HALAlarms:  []HALAlarm{},
			},
		},
		{
			given: models.ModelsGetPointAlarmThresholdResponse{
				ThresholdType: i32p(1), // overall in window
			},
			expected: &Threshold{
				ThresholdType: ThresholdTypeOverallInWindow,
				BandAlarms:    []BandAlarm{},
				HALAlarms:     []HALAlarm{},
			},
		},
		{
			given: models.ModelsGetPointAlarmThresholdResponse{
				ThresholdType: i32p(2), // overall out of window
			},
			expected: &Threshold{
				ThresholdType: ThresholdTypeOverallOutOfWindow,
				BandAlarms:    []BandAlarm{},
				HALAlarms:     []HALAlarm{},
			},
		},
		{
			given: models.ModelsGetPointAlarmThresholdResponse{
				ThresholdType: i32p(3), // inspection
			},
			expected: &Threshold{
				ThresholdType: ThresholdTypeInspection,
				BandAlarms:    []BandAlarm{},
				HALAlarms:     []HALAlarm{},
			},
		},
		{
			given: models.ModelsGetPointAlarmThresholdResponse{
				ThresholdType: i32p(2), // overall out of window
				Overall: &models.ModelsOverall{
					Unit:      "gE",
					OuterHigh: f64p(4),
					InnerHigh: f64p(3),
					InnerLow:  f64p(2),
					OuterLow:  f64p(1),
				},
			},
			expected: &Threshold{
				ThresholdType: ThresholdTypeOverallOutOfWindow,
				Overall: &Overall{
					Unit:      "gE",
					OuterHigh: f64p(4),
					InnerHigh: f64p(3),
					InnerLow:  f64p(2),
					OuterLow:  f64p(1),
				},
				BandAlarms: []BandAlarm{},
				HALAlarms:  []HALAlarm{},
			},
		},
		{
			given: models.ModelsGetPointAlarmThresholdResponse{
				RateOfChange: &models.ModelsRateOfChange{
					Unit:      "gE",
					OuterHigh: f64p(2),
					InnerHigh: f64p(1),
					InnerLow:  f64p(-1),
					OuterLow:  f64p(-2),
				},
			},
			expected: &Threshold{
				RateOfChange: &RateOfChange{
					Unit:      "gE",
					OuterHigh: f64p(2),
					InnerHigh: f64p(1),
					InnerLow:  f64p(-1),
					OuterLow:  f64p(-2),
				},
				BandAlarms: []BandAlarm{},
				HALAlarms:  []HALAlarm{},
			},
		},
		{
			given: models.ModelsGetPointAlarmThresholdResponse{
				Inspection: &models.ModelsInspection{
					Choices: []*models.ModelsInspectionChoice{
						{
							Answer:      "good",
							Instruction: "good?",
							Status:      i32p(2),
						},
					},
				},
			},
			expected: &Threshold{
				Inspection: &Inspection{
					Choices: []InspectionChoice{
						{
							Answer:      "good",
							Instruction: "good?",
							Status:      AlarmStatusGood,
						},
					},
				},
				BandAlarms: []BandAlarm{},
				HALAlarms:  []HALAlarm{},
			},
		},
		{
			given: models.ModelsGetPointAlarmThresholdResponse{
				BandAlarms: []*models.ModelsBandAlarm{
					{},
				},
			},
			expected: &Threshold{
				BandAlarms: []BandAlarm{
					{},
				},
				HALAlarms: []HALAlarm{},
			},
		},
		{
			given: models.ModelsGetPointAlarmThresholdResponse{
				HalAlarms: []*models.ModelsHALAlarm{
					{},
				},
			},
			expected: &Threshold{
				BandAlarms: []BandAlarm{},
				HALAlarms: []HALAlarm{
					{},
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(Threshold)

			err := actual.FromInternal(test.given)
			require.NoError(t, err)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_ThresholdFromInternal_isNil(t *testing.T) {
	t.Parallel()

	var threshold *Threshold

	err := threshold.FromInternal(models.ModelsGetPointAlarmThresholdResponse{})

	assert.NoError(t, err)
}

func Test_ThresholdFromInternal_invalidNodeID(t *testing.T) {
	t.Parallel()

	var (
		threshold = new(Threshold)
		nodeID    = strfmt.UUID("not-valid")
	)

	err := threshold.FromInternal(models.ModelsGetPointAlarmThresholdResponse{
		NodeID: &nodeID,
	})

	assert.Error(t, err)
}

func Test_Threshold_ToInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *Threshold
		expected models.ModelsSetPointAlarmThresholdRequest
	}{
		{
			given: &Threshold{
				BandAlarms: []BandAlarm{},
				HALAlarms:  []HALAlarm{},
			},
			expected: models.ModelsSetPointAlarmThresholdRequest{
				ThresholdType: i32p(0),
				BandAlarms:    []*models.ModelsBandAlarm{},
				HalAlarms:     []*models.ModelsHALAlarm{},
			},
		},
		{
			given: &Threshold{
				ThresholdType: ThresholdTypeOverallInWindow,
				BandAlarms:    []BandAlarm{},
				HALAlarms:     []HALAlarm{},
			},
			expected: models.ModelsSetPointAlarmThresholdRequest{
				ThresholdType: i32p(1), // overall in window
				BandAlarms:    []*models.ModelsBandAlarm{},
				HalAlarms:     []*models.ModelsHALAlarm{},
			},
		},
		{
			given: &Threshold{
				ThresholdType: ThresholdTypeOverallOutOfWindow,
				BandAlarms:    []BandAlarm{},
				HALAlarms:     []HALAlarm{},
			},
			expected: models.ModelsSetPointAlarmThresholdRequest{
				ThresholdType: i32p(2), // overall out of window
				BandAlarms:    []*models.ModelsBandAlarm{},
				HalAlarms:     []*models.ModelsHALAlarm{},
			},
		},
		{
			given: &Threshold{
				ThresholdType: ThresholdTypeInspection,
				BandAlarms:    []BandAlarm{},
				HALAlarms:     []HALAlarm{},
			},
			expected: models.ModelsSetPointAlarmThresholdRequest{
				ThresholdType: i32p(3), // inspection
				BandAlarms:    []*models.ModelsBandAlarm{},
				HalAlarms:     []*models.ModelsHALAlarm{},
			},
		},
		{
			given: &Threshold{
				ThresholdType: ThresholdTypeOverallOutOfWindow,
				Overall: &Overall{
					Unit:      "gE",
					OuterHigh: f64p(4),
					InnerHigh: f64p(3),
					InnerLow:  f64p(2),
					OuterLow:  f64p(1),
				},
				BandAlarms: []BandAlarm{},
				HALAlarms:  []HALAlarm{},
			},
			expected: models.ModelsSetPointAlarmThresholdRequest{
				ThresholdType: i32p(2), // overall out of window
				Overall: &models.ModelsOverall{
					Unit:      "gE",
					OuterHigh: f64p(4),
					InnerHigh: f64p(3),
					InnerLow:  f64p(2),
					OuterLow:  f64p(1),
				},
				BandAlarms: []*models.ModelsBandAlarm{},
				HalAlarms:  []*models.ModelsHALAlarm{},
			},
		},
		{
			given: &Threshold{
				RateOfChange: &RateOfChange{
					Unit:      "gE",
					OuterHigh: f64p(2),
					InnerHigh: f64p(1),
					InnerLow:  f64p(-1),
					OuterLow:  f64p(-2),
				},
				BandAlarms: []BandAlarm{},
				HALAlarms:  []HALAlarm{},
			},
			expected: models.ModelsSetPointAlarmThresholdRequest{
				ThresholdType: i32p(0),
				RateOfChange: &models.ModelsRateOfChange{
					Unit:      "gE",
					OuterHigh: f64p(2),
					InnerHigh: f64p(1),
					InnerLow:  f64p(-1),
					OuterLow:  f64p(-2),
				},
				BandAlarms: []*models.ModelsBandAlarm{},
				HalAlarms:  []*models.ModelsHALAlarm{},
			},
		},
		{
			given: &Threshold{
				ThresholdType: ThresholdTypeInspection,
				Inspection: &Inspection{
					Choices: []InspectionChoice{
						{
							Answer:      "good",
							Instruction: "good?",
							Status:      AlarmStatusGood,
						},
					},
				},
				BandAlarms: []BandAlarm{},
				HALAlarms:  []HALAlarm{},
			},
			expected: models.ModelsSetPointAlarmThresholdRequest{
				ThresholdType: i32p(3), // inspection
				Inspection: &models.ModelsInspection{
					Choices: []*models.ModelsInspectionChoice{
						{
							Answer:      "good",
							Instruction: "good?",
							Status:      i32p(2),
						},
					},
				},
				BandAlarms: []*models.ModelsBandAlarm{},
				HalAlarms:  []*models.ModelsHALAlarm{},
			},
		},
		{
			given: &Threshold{
				BandAlarms: []BandAlarm{
					{},
				},
				HALAlarms: []HALAlarm{},
			},
			expected: models.ModelsSetPointAlarmThresholdRequest{
				ThresholdType: i32p(0),
				BandAlarms: []*models.ModelsBandAlarm{
					{
						MinFrequency: &models.ModelsBandAlarmFrequency{
							ValueType: i32p(0),
							Value:     f64p(0),
						},
						MaxFrequency: &models.ModelsBandAlarmFrequency{
							ValueType: i32p(0),
							Value:     f64p(0),
						},
					},
				},
				HalAlarms: []*models.ModelsHALAlarm{},
			},
		},
		{
			given: &Threshold{
				FullScale:  f64p(0.5),
				BandAlarms: []BandAlarm{},
				HALAlarms:  []HALAlarm{},
			},
			expected: models.ModelsSetPointAlarmThresholdRequest{
				ThresholdType: i32p(0),
				FullScale:     0.5,
				BandAlarms:    []*models.ModelsBandAlarm{},
				HalAlarms:     []*models.ModelsHALAlarm{},
			},
		},
		{
			given: &Threshold{
				BandAlarms: []BandAlarm{},
				HALAlarms: []HALAlarm{
					{},
				},
			},
			expected: models.ModelsSetPointAlarmThresholdRequest{
				ThresholdType: i32p(0),
				BandAlarms:    []*models.ModelsBandAlarm{},
				HalAlarms: []*models.ModelsHALAlarm{
					{},
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

func Test_Threshold_ToInternal_IsNil(t *testing.T) {
	assert.NotPanics(t, func() {
		var threshold *Threshold

		expected := models.ModelsSetPointAlarmThresholdRequest{}

		actual := threshold.ToInternal()

		assert.Equal(t, expected, actual)
	})
}
