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
