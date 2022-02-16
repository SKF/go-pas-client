package models

import (
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/SKF/go-pas-client/internal/events"
	models "github.com/SKF/go-pas-client/internal/models"
	"github.com/SKF/go-utility/v2/uuid"
	pas "github.com/SKF/proto/v2/pas"
)

func Test_BandAlarm_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *models.ModelsBandAlarm
		expected *BandAlarm
	}{
		{
			given:    nil,
			expected: &BandAlarm{},
		},
		{
			given:    &models.ModelsBandAlarm{},
			expected: &BandAlarm{},
		},
		{
			given: &models.ModelsBandAlarm{
				MaxFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(1),
					Value:     f64p(2.0),
				},
			},
			expected: &BandAlarm{
				MaxFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencyFixed,
					Value:     2.0,
				},
			},
		},
		{
			given: &models.ModelsBandAlarm{
				MinFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(1),
					Value:     f64p(1.0),
				},
			},
			expected: &BandAlarm{
				MinFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencyFixed,
					Value:     1.0,
				},
			},
		},
		{
			given: &models.ModelsBandAlarm{
				MaxFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(2),
					Value:     f64p(2.0),
				},
			},
			expected: &BandAlarm{
				MaxFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencySpeedMultiple,
					Value:     2.0,
				},
			},
		},
		{
			given: &models.ModelsBandAlarm{
				MinFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(2),
					Value:     f64p(1.0),
				},
			},
			expected: &BandAlarm{
				MinFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencySpeedMultiple,
					Value:     1.0,
				},
			},
		},
		{
			given: &models.ModelsBandAlarm{
				OverallThreshold: &models.ModelsBandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &models.ModelsBandAlarmThreshold{
						ValueType: i32p(1), // absolute
						Value:     f64p(10),
					},
				},
			},
			expected: &BandAlarm{
				OverallThreshold: &BandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &BandAlarmThreshold{
						ValueType: BandAlarmThresholdTypeAbsolute,
						Value:     10,
					},
				},
			},
		},
		{
			given: &models.ModelsBandAlarm{
				OverallThreshold: &models.ModelsBandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &models.ModelsBandAlarmThreshold{
						ValueType: i32p(2), // relative fullscale
						Value:     f64p(10),
					},
				},
			},
			expected: &BandAlarm{
				OverallThreshold: &BandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &BandAlarmThreshold{
						ValueType: BandAlarmThresholdTypeRelativeFullscale,
						Value:     10,
					},
				},
			},
		},
		{
			given: &models.ModelsBandAlarm{
				OverallThreshold: &models.ModelsBandAlarmOverallThreshold{
					Unit: "gE",
					UpperDanger: &models.ModelsBandAlarmThreshold{
						ValueType: i32p(1), // absolute
						Value:     f64p(10),
					},
				},
			},
			expected: &BandAlarm{
				OverallThreshold: &BandAlarmOverallThreshold{
					Unit: "gE",
					UpperDanger: &BandAlarmThreshold{
						ValueType: BandAlarmThresholdTypeAbsolute,
						Value:     10,
					},
				},
			},
		},
		{
			given: &models.ModelsBandAlarm{
				OverallThreshold: &models.ModelsBandAlarmOverallThreshold{
					Unit: "gE",
					UpperDanger: &models.ModelsBandAlarmThreshold{
						ValueType: i32p(2), // relative fullscale
						Value:     f64p(10),
					},
				},
			},
			expected: &BandAlarm{
				OverallThreshold: &BandAlarmOverallThreshold{
					Unit: "gE",
					UpperDanger: &BandAlarmThreshold{
						ValueType: BandAlarmThresholdTypeRelativeFullscale,
						Value:     10,
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(BandAlarm)

			actual.FromInternal(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_BandAlarm_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var bandAlarm *BandAlarm

		bandAlarm.FromInternal(&models.ModelsBandAlarm{})
	})

	assert.NotPanics(t, func() {
		bandAlarm := new(BandAlarm)

		bandAlarm.FromInternal(nil)
	})
}

func Test_BandAlarm_ToInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *BandAlarm
		expected *models.ModelsBandAlarm
	}{
		{
			given:    nil,
			expected: nil,
		},
		{
			given: &BandAlarm{
				MaxFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencyFixed,
					Value:     2.0,
				},
			},
			expected: &models.ModelsBandAlarm{
				MinFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(0),
					Value:     f64p(0),
				},
				MaxFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(1),
					Value:     f64p(2.0),
				},
			},
		},
		{
			given: &BandAlarm{
				MinFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencyFixed,
					Value:     1.0,
				},
			},
			expected: &models.ModelsBandAlarm{
				MinFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(1),
					Value:     f64p(1.0),
				},
				MaxFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(0),
					Value:     f64p(0),
				},
			},
		},
		{
			given: &BandAlarm{
				MaxFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencySpeedMultiple,
					Value:     2.0,
				},
			},
			expected: &models.ModelsBandAlarm{
				MinFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(0),
					Value:     f64p(0),
				},
				MaxFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(2),
					Value:     f64p(2.0),
				},
			},
		},
		{
			given: &BandAlarm{
				MinFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencySpeedMultiple,
					Value:     1.0,
				},
			},
			expected: &models.ModelsBandAlarm{
				MinFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(2),
					Value:     f64p(1.0),
				},
				MaxFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(0),
					Value:     f64p(0),
				},
			},
		},
		{
			given: &BandAlarm{
				OverallThreshold: &BandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &BandAlarmThreshold{
						ValueType: BandAlarmThresholdTypeAbsolute,
						Value:     10,
					},
				},
			},
			expected: &models.ModelsBandAlarm{
				MinFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(0),
					Value:     f64p(0),
				},
				MaxFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(0),
					Value:     f64p(0),
				},
				OverallThreshold: &models.ModelsBandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &models.ModelsBandAlarmThreshold{
						ValueType: i32p(1), // absolute
						Value:     f64p(10),
					},
				},
			},
		},
		{
			given: &BandAlarm{
				OverallThreshold: &BandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &BandAlarmThreshold{
						ValueType: BandAlarmThresholdTypeRelativeFullscale,
						Value:     10,
					},
				},
			},
			expected: &models.ModelsBandAlarm{
				MinFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(0),
					Value:     f64p(0),
				},
				MaxFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(0),
					Value:     f64p(0),
				},
				OverallThreshold: &models.ModelsBandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &models.ModelsBandAlarmThreshold{
						ValueType: i32p(2), // relative fullscale
						Value:     f64p(10),
					},
				},
			},
		},
		{
			given: &BandAlarm{
				OverallThreshold: &BandAlarmOverallThreshold{
					Unit: "gE",
					UpperDanger: &BandAlarmThreshold{
						ValueType: BandAlarmThresholdTypeAbsolute,
						Value:     10,
					},
				},
			},
			expected: &models.ModelsBandAlarm{
				MinFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(0),
					Value:     f64p(0),
				},
				MaxFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(0),
					Value:     f64p(0),
				},
				OverallThreshold: &models.ModelsBandAlarmOverallThreshold{
					Unit: "gE",
					UpperDanger: &models.ModelsBandAlarmThreshold{
						ValueType: i32p(1), // absolute
						Value:     f64p(10),
					},
				},
			},
		},
		{
			given: &BandAlarm{
				OverallThreshold: &BandAlarmOverallThreshold{
					Unit: "gE",
					UpperDanger: &BandAlarmThreshold{
						ValueType: BandAlarmThresholdTypeRelativeFullscale,
						Value:     10,
					},
				},
			},
			expected: &models.ModelsBandAlarm{
				MinFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(0),
					Value:     f64p(0),
				},
				MaxFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(0),
					Value:     f64p(0),
				},
				OverallThreshold: &models.ModelsBandAlarmOverallThreshold{
					Unit: "gE",
					UpperDanger: &models.ModelsBandAlarmThreshold{
						ValueType: i32p(2), // relative fullscale
						Value:     f64p(10),
					},
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

func Test_BandAlarm_ToInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var bandAlarm *BandAlarm

		actual := bandAlarm.ToInternal()

		assert.Nil(t, actual)
	})
}

func Test_BandAlarm_FromProto(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *pas.BandAlarm
		expected *BandAlarm
	}{
		{
			given: &pas.BandAlarm{
				MaxFrequency: &pas.Frequency{
					ValueType: pas.Frequency_FIXED,
					Value: &pas.DoubleObject{
						Value: 2.0,
					},
				},
			},
			expected: &BandAlarm{
				MaxFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencyFixed,
					Value:     2.0,
				},
			},
		},
		{
			given: &pas.BandAlarm{
				MinFrequency: &pas.Frequency{
					ValueType: pas.Frequency_FIXED,
					Value: &pas.DoubleObject{
						Value: 1.0,
					},
				},
			},
			expected: &BandAlarm{
				MinFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencyFixed,
					Value:     1.0,
				},
			},
		},
		{
			given: &pas.BandAlarm{
				MaxFrequency: &pas.Frequency{
					ValueType: pas.Frequency_SPEED_MULTIPLE,
					Value: &pas.DoubleObject{
						Value: 2.0,
					},
				},
			},
			expected: &BandAlarm{
				MaxFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencySpeedMultiple,
					Value:     2.0,
				},
			},
		},
		{
			given: &pas.BandAlarm{
				MinFrequency: &pas.Frequency{
					ValueType: pas.Frequency_SPEED_MULTIPLE,
					Value: &pas.DoubleObject{
						Value: 1.0,
					},
				},
			},
			expected: &BandAlarm{
				MinFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencySpeedMultiple,
					Value:     1.0,
				},
			},
		},
		{
			given: &pas.BandAlarm{
				OverallThreshold: &pas.BandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &pas.ThresholdValue{
						ValueType: pas.ThresholdValue_ABSOLUTE,
						Value: &pas.DoubleObject{
							Value: 10,
						},
					},
				},
			},
			expected: &BandAlarm{
				OverallThreshold: &BandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &BandAlarmThreshold{
						ValueType: BandAlarmThresholdTypeAbsolute,
						Value:     10,
					},
				},
			},
		},
		{
			given: &pas.BandAlarm{
				OverallThreshold: &pas.BandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &pas.ThresholdValue{
						ValueType: pas.ThresholdValue_RELATIVE_FULLSCALE,
						Value: &pas.DoubleObject{
							Value: 10,
						},
					},
				},
			},
			expected: &BandAlarm{
				OverallThreshold: &BandAlarmOverallThreshold{
					Unit: "gE",
					UpperAlert: &BandAlarmThreshold{
						ValueType: BandAlarmThresholdTypeRelativeFullscale,
						Value:     10,
					},
				},
			},
		},
		{
			given: &pas.BandAlarm{
				OverallThreshold: &pas.BandAlarmOverallThreshold{
					Unit: "gE",
					UpperDanger: &pas.ThresholdValue{
						ValueType: pas.ThresholdValue_ABSOLUTE,
						Value: &pas.DoubleObject{
							Value: 10,
						},
					},
				},
			},
			expected: &BandAlarm{
				OverallThreshold: &BandAlarmOverallThreshold{
					Unit: "gE",
					UpperDanger: &BandAlarmThreshold{
						ValueType: BandAlarmThresholdTypeAbsolute,
						Value:     10,
					},
				},
			},
		},
		{
			given: &pas.BandAlarm{
				OverallThreshold: &pas.BandAlarmOverallThreshold{
					Unit: "gE",
					UpperDanger: &pas.ThresholdValue{
						ValueType: pas.ThresholdValue_RELATIVE_FULLSCALE,
						Value: &pas.DoubleObject{
							Value: 10,
						},
					},
				},
			},
			expected: &BandAlarm{
				OverallThreshold: &BandAlarmOverallThreshold{
					Unit: "gE",
					UpperDanger: &BandAlarmThreshold{
						ValueType: BandAlarmThresholdTypeRelativeFullscale,
						Value:     10,
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			buf, err := proto.Marshal(test.given)
			require.NoError(t, err)

			actual := new(BandAlarm)

			err = actual.FromProto(buf)
			require.NoError(t, err)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_BandAlarm_FromProto_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var actual *BandAlarm

		err := actual.FromProto([]byte{})

		assert.NoError(t, err)
	})

	assert.NotPanics(t, func() {
		actual := new(BandAlarm)

		err := actual.FromProto(nil)

		assert.NoError(t, err)
	})
}

func Test_BandAlarm_FromProto_InvalidBody(t *testing.T) {
	t.Parallel()

	actual := new(BandAlarm)

	err := actual.FromProto([]byte("not-valid"))

	assert.Error(t, err)
}

func Test_BandAlarmFrequency_FromInternalThreshold_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var frequency *BandAlarmFrequency

		frequency.FromInternalThreshold(&models.ModelsBandAlarmFrequency{})
	})

	assert.NotPanics(t, func() {
		frequency := &BandAlarmFrequency{}

		frequency.FromInternalThreshold(nil)
	})
}

func Test_BandAlarmFrequency_FromInternalAlarmStatus_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var frequency *BandAlarmFrequency

		frequency.FromInternalAlarmStatus(&models.ModelsGetAlarmStatusResponseFrequency{})
	})

	assert.NotPanics(t, func() {
		frequency := new(BandAlarmFrequency)

		frequency.FromInternalAlarmStatus(nil)
	})
}

func Test_BandAlarmFrequency_ToInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var frequency *BandAlarmFrequency

		actual := frequency.ToInternal()

		assert.Nil(t, actual)
	})
}

func Test_BandAlarmFrequency_FromProto_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var frequency *BandAlarmFrequency

		frequency.FromProto(&pas.Frequency{})
	})

	assert.NotPanics(t, func() {
		frequency := new(BandAlarmFrequency)

		frequency.FromProto(nil)
	})
}

func Test_BandAlarmOverallThreshold_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *models.ModelsBandAlarmOverallThreshold
		expected *BandAlarmOverallThreshold
	}{
		{
			given:    nil,
			expected: &BandAlarmOverallThreshold{},
		},
		{
			given:    &models.ModelsBandAlarmOverallThreshold{},
			expected: &BandAlarmOverallThreshold{},
		},
		{
			given: &models.ModelsBandAlarmOverallThreshold{
				UpperAlert: &models.ModelsBandAlarmThreshold{},
			},
			expected: &BandAlarmOverallThreshold{
				UpperAlert: &BandAlarmThreshold{
					ValueType: BandAlarmThresholdTypeUnknown,
					Value:     0,
				},
			},
		},
		{
			given: &models.ModelsBandAlarmOverallThreshold{
				UpperAlert: &models.ModelsBandAlarmThreshold{
					ValueType: i32p(1), // absolute
					Value:     f64p(1),
				},
			},
			expected: &BandAlarmOverallThreshold{
				UpperAlert: &BandAlarmThreshold{
					ValueType: BandAlarmThresholdTypeAbsolute,
					Value:     1,
				},
			},
		},
		{
			given: &models.ModelsBandAlarmOverallThreshold{
				UpperAlert: &models.ModelsBandAlarmThreshold{
					ValueType: i32p(2), // relative fullscale
					Value:     f64p(2),
				},
			},
			expected: &BandAlarmOverallThreshold{
				UpperAlert: &BandAlarmThreshold{
					ValueType: BandAlarmThresholdTypeRelativeFullscale,
					Value:     2,
				},
			},
		},
		{
			given: &models.ModelsBandAlarmOverallThreshold{
				UpperDanger: &models.ModelsBandAlarmThreshold{},
			},
			expected: &BandAlarmOverallThreshold{
				UpperDanger: &BandAlarmThreshold{
					ValueType: BandAlarmThresholdTypeUnknown,
					Value:     0,
				},
			},
		},
		{
			given: &models.ModelsBandAlarmOverallThreshold{
				UpperDanger: &models.ModelsBandAlarmThreshold{
					ValueType: i32p(1), // absolute
					Value:     f64p(1),
				},
			},
			expected: &BandAlarmOverallThreshold{
				UpperDanger: &BandAlarmThreshold{
					ValueType: BandAlarmThresholdTypeAbsolute,
					Value:     1,
				},
			},
		},
		{
			given: &models.ModelsBandAlarmOverallThreshold{
				UpperDanger: &models.ModelsBandAlarmThreshold{
					ValueType: i32p(2), // relative fullscale
					Value:     f64p(2),
				},
			},
			expected: &BandAlarmOverallThreshold{
				UpperDanger: &BandAlarmThreshold{
					ValueType: BandAlarmThresholdTypeRelativeFullscale,
					Value:     2,
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(BandAlarmOverallThreshold)

			actual.FromInternal(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_BandAlarmOverallThreshold_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var threshold *BandAlarmOverallThreshold

		threshold.FromInternal(&models.ModelsBandAlarmOverallThreshold{})
	})

	assert.NotPanics(t, func() {
		threshold := new(BandAlarmOverallThreshold)

		threshold.FromInternal(nil)
	})
}

func Test_BandAlarmOverallThreshold_ToInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *BandAlarmOverallThreshold
		expected *models.ModelsBandAlarmOverallThreshold
	}{
		{
			given:    nil,
			expected: nil,
		},
		{
			given:    &BandAlarmOverallThreshold{},
			expected: &models.ModelsBandAlarmOverallThreshold{},
		},
		{
			given: &BandAlarmOverallThreshold{
				UpperAlert: &BandAlarmThreshold{
					ValueType: BandAlarmThresholdTypeAbsolute,
					Value:     1,
				},
			},
			expected: &models.ModelsBandAlarmOverallThreshold{
				UpperAlert: &models.ModelsBandAlarmThreshold{
					ValueType: i32p(1), // absolute
					Value:     f64p(1),
				},
			},
		},
		{
			given: &BandAlarmOverallThreshold{
				UpperAlert: &BandAlarmThreshold{
					ValueType: BandAlarmThresholdTypeRelativeFullscale,
					Value:     2,
				},
			},
			expected: &models.ModelsBandAlarmOverallThreshold{
				UpperAlert: &models.ModelsBandAlarmThreshold{
					ValueType: i32p(2), // relative fullscale
					Value:     f64p(2),
				},
			},
		},
		{
			given: &BandAlarmOverallThreshold{
				UpperDanger: &BandAlarmThreshold{
					ValueType: BandAlarmThresholdTypeAbsolute,
					Value:     1,
				},
			},
			expected: &models.ModelsBandAlarmOverallThreshold{
				UpperDanger: &models.ModelsBandAlarmThreshold{
					ValueType: i32p(1), // absolute
					Value:     f64p(1),
				},
			},
		},
		{
			given: &BandAlarmOverallThreshold{
				UpperDanger: &BandAlarmThreshold{
					ValueType: BandAlarmThresholdTypeRelativeFullscale,
					Value:     2,
				},
			},
			expected: &models.ModelsBandAlarmOverallThreshold{
				UpperDanger: &models.ModelsBandAlarmThreshold{
					ValueType: i32p(2), // relative fullscale
					Value:     f64p(2),
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

func Test_BandAlarmOverallThreshold_ToInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var threshold *BandAlarmOverallThreshold

		actual := threshold.ToInternal()

		assert.Nil(t, actual)
	})
}

func Test_BandAlarmOverallThreshold_FromProto(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *pas.BandAlarmOverallThreshold
		expected *BandAlarmOverallThreshold
	}{
		{
			given: &pas.BandAlarmOverallThreshold{
				UpperAlert: &pas.ThresholdValue{
					ValueType: pas.ThresholdValue_ABSOLUTE,
					Value: &pas.DoubleObject{
						Value: 1,
					},
				},
			},
			expected: &BandAlarmOverallThreshold{
				UpperAlert: &BandAlarmThreshold{
					ValueType: BandAlarmThresholdTypeAbsolute,
					Value:     1,
				},
			},
		},
		{
			given: &pas.BandAlarmOverallThreshold{
				UpperAlert: &pas.ThresholdValue{
					ValueType: pas.ThresholdValue_RELATIVE_FULLSCALE,
					Value: &pas.DoubleObject{
						Value: 2,
					},
				},
			},
			expected: &BandAlarmOverallThreshold{
				UpperAlert: &BandAlarmThreshold{
					ValueType: BandAlarmThresholdTypeRelativeFullscale,
					Value:     2,
				},
			},
		},
		{
			given: &pas.BandAlarmOverallThreshold{
				UpperDanger: &pas.ThresholdValue{
					ValueType: pas.ThresholdValue_ABSOLUTE,
					Value: &pas.DoubleObject{
						Value: 1,
					},
				},
			},
			expected: &BandAlarmOverallThreshold{
				UpperDanger: &BandAlarmThreshold{
					ValueType: BandAlarmThresholdTypeAbsolute,
					Value:     1,
				},
			},
		},
		{
			given: &pas.BandAlarmOverallThreshold{
				UpperDanger: &pas.ThresholdValue{
					ValueType: pas.ThresholdValue_RELATIVE_FULLSCALE,
					Value: &pas.DoubleObject{
						Value: 2,
					},
				},
			},
			expected: &BandAlarmOverallThreshold{
				UpperDanger: &BandAlarmThreshold{
					ValueType: BandAlarmThresholdTypeRelativeFullscale,
					Value:     2,
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(BandAlarmOverallThreshold)

			actual.FromProto(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_BandAlarmOverallThreshold_FromProto_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var threshold *BandAlarmOverallThreshold

		threshold.FromProto(&pas.BandAlarmOverallThreshold{})
	})

	assert.NotPanics(t, func() {
		threshold := new(BandAlarmOverallThreshold)

		threshold.FromProto(nil)
	})
}

func Test_BandAlarmThreshold_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var threshold *BandAlarmThreshold

		threshold.FromInternal(&models.ModelsBandAlarmThreshold{})
	})

	assert.NotPanics(t, func() {
		threshold := new(BandAlarmThreshold)

		threshold.FromInternal(nil)
	})
}

func Test_BandAlarmThreshold_ToInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var threshold *BandAlarmThreshold

		actual := threshold.ToInternal()

		assert.Nil(t, actual)
	})
}

func Test_BandAlarmThreshold_FromProto_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var threshold *BandAlarmThreshold

		threshold.FromProto(&pas.ThresholdValue{})
	})

	assert.NotPanics(t, func() {
		threshold := &BandAlarmThreshold{}

		threshold.FromProto(nil)
	})
}

func Test_BandAlarmStatus_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *models.ModelsGetAlarmStatusResponseBandAlarm
		expected *BandAlarmStatus
	}{
		{
			given:    &models.ModelsGetAlarmStatusResponseBandAlarm{},
			expected: &BandAlarmStatus{},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseBandAlarm{
				Label: "10x RPM",
			},
			expected: &BandAlarmStatus{
				Label: "10x RPM",
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseBandAlarm{
				Status: i32p(1), // no data
			},
			expected: &BandAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusNoData,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseBandAlarm{
				Status: i32p(2), // good
			},
			expected: &BandAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusGood,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseBandAlarm{
				Status: i32p(3), // alert
			},
			expected: &BandAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusAlert,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseBandAlarm{
				Status: i32p(4), // danger
			},
			expected: &BandAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusDanger,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseBandAlarm{
				TriggeringMeasurement: strfmt.UUID(uuid.EmptyUUID.String()),
			},
			expected: &BandAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					TriggeringMeasurement: uuid.EmptyUUID,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseBandAlarm{
				MinFrequency: &models.ModelsGetAlarmStatusResponseFrequency{
					ValueType: i32p(1), // fixed
					Value:     f64p(100),
				},
			},
			expected: &BandAlarmStatus{
				MinFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencyFixed,
					Value:     100,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseBandAlarm{
				MinFrequency: &models.ModelsGetAlarmStatusResponseFrequency{
					ValueType: i32p(2), // speed multiple
					Value:     f64p(200),
				},
			},
			expected: &BandAlarmStatus{
				MinFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencySpeedMultiple,
					Value:     200,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseBandAlarm{
				MaxFrequency: &models.ModelsGetAlarmStatusResponseFrequency{
					ValueType: i32p(1), // fixed
					Value:     f64p(100),
				},
			},
			expected: &BandAlarmStatus{
				MaxFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencyFixed,
					Value:     100,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseBandAlarm{
				MaxFrequency: &models.ModelsGetAlarmStatusResponseFrequency{
					ValueType: i32p(2), // speed multiple
					Value:     f64p(200),
				},
			},
			expected: &BandAlarmStatus{
				MaxFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencySpeedMultiple,
					Value:     200,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseBandAlarm{
				CalculatedOverall: &models.ModelsBandCalculatedOverall{
					Unit:  "gE",
					Value: f64p(5),
				},
			},
			expected: &BandAlarmStatus{
				CalculatedOverall: &BandAlarmStatusCalculatedOverall{
					Unit:  "gE",
					Value: 5,
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(BandAlarmStatus)

			actual.FromInternal(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_BandAlarmStatus_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var status *BandAlarmStatus

		status.FromInternal(&models.ModelsGetAlarmStatusResponseBandAlarm{})
	})

	assert.NotPanics(t, func() {
		status := new(BandAlarmStatus)

		status.FromInternal(nil)
	})
}

func Test_BandAlarmStatus_FromEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    events.BandAlarmStatus
		expected *BandAlarmStatus
	}{
		{
			given:    events.BandAlarmStatus{},
			expected: &BandAlarmStatus{},
		},
		{
			given: events.BandAlarmStatus{
				Label: "10x RPM",
			},
			expected: &BandAlarmStatus{
				Label: "10x RPM",
			},
		},
		{
			given: events.BandAlarmStatus{
				Status: 1, // no data
			},
			expected: &BandAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusNoData,
				},
			},
		},
		{
			given: events.BandAlarmStatus{
				Status: 2, // good
			},
			expected: &BandAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusGood,
				},
			},
		},
		{
			given: events.BandAlarmStatus{
				Status: 3, // alert
			},
			expected: &BandAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusAlert,
				},
			},
		},
		{
			given: events.BandAlarmStatus{
				Status: 4, // danger
			},
			expected: &BandAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusDanger,
				},
			},
		},
		{
			given: events.BandAlarmStatus{
				TriggeringMeasurement: uuid.EmptyUUID,
			},
			expected: &BandAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					TriggeringMeasurement: uuid.EmptyUUID,
				},
			},
		},
		{
			given: events.BandAlarmStatus{
				MinFrequency: events.Frequency{
					ValueType: 1, // fixed
					Value:     100,
				},
			},
			expected: &BandAlarmStatus{
				MinFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencyFixed,
					Value:     100,
				},
			},
		},
		{
			given: events.BandAlarmStatus{
				MinFrequency: events.Frequency{
					ValueType: 2, // speed multiple
					Value:     200,
				},
			},
			expected: &BandAlarmStatus{
				MinFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencySpeedMultiple,
					Value:     200,
				},
			},
		},
		{
			given: events.BandAlarmStatus{
				MaxFrequency: events.Frequency{
					ValueType: 1, // fixed
					Value:     100,
				},
			},
			expected: &BandAlarmStatus{
				MaxFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencyFixed,
					Value:     100,
				},
			},
		},
		{
			given: events.BandAlarmStatus{
				MaxFrequency: events.Frequency{
					ValueType: 2, // speed multiple
					Value:     200,
				},
			},
			expected: &BandAlarmStatus{
				MaxFrequency: BandAlarmFrequency{
					ValueType: BandAlarmFrequencySpeedMultiple,
					Value:     200,
				},
			},
		},
		{
			given: events.BandAlarmStatus{
				CalculatedOverall: &events.CalculatedOverall{
					Unit:  "gE",
					Value: 5,
				},
			},
			expected: &BandAlarmStatus{
				CalculatedOverall: &BandAlarmStatusCalculatedOverall{
					Unit:  "gE",
					Value: 5,
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(BandAlarmStatus)

			actual.FromEvent(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_BandAlarmStatus_FromEvent_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var status *BandAlarmStatus

		status.FromEvent(events.BandAlarmStatus{})
	})
}
