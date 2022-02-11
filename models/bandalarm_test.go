package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	models "github.com/SKF/go-pas-client/internal/models"
)

func Test_BandAlarm_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *models.ModelsBandAlarm
		expected BandAlarm
	}{
		{
			given:    nil,
			expected: BandAlarm{},
		},
		{
			given:    &models.ModelsBandAlarm{},
			expected: BandAlarm{},
		},
		{
			given: &models.ModelsBandAlarm{
				MaxFrequency: &models.ModelsBandAlarmFrequency{
					ValueType: i32p(1),
					Value:     f64p(2.0),
				},
			},
			expected: BandAlarm{
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
			expected: BandAlarm{
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
			expected: BandAlarm{
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
			expected: BandAlarm{
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
			expected: BandAlarm{
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
			expected: BandAlarm{
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
			expected: BandAlarm{
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
			expected: BandAlarm{
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
			actual := BandAlarm{}

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
		bandAlarm := &BandAlarm{}

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

func Test_BandAlarmFrequency_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var frequency *BandAlarmFrequency

		frequency.FromInternal(&models.ModelsBandAlarmFrequency{})
	})

	assert.NotPanics(t, func() {
		frequency := &BandAlarmFrequency{}

		frequency.FromInternal(nil)
	})
}

func Test_BandAlarmFrequency_ToInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var thisIsNil *BandAlarmFrequency

		actual := thisIsNil.ToInternal()

		assert.Nil(t, actual)
	})
}

func Test_BandAlarmOverallThreshold_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *models.ModelsBandAlarmOverallThreshold
		expected BandAlarmOverallThreshold
	}{
		{
			given:    nil,
			expected: BandAlarmOverallThreshold{},
		},
		{
			given:    &models.ModelsBandAlarmOverallThreshold{},
			expected: BandAlarmOverallThreshold{},
		},
		{
			given: &models.ModelsBandAlarmOverallThreshold{
				UpperAlert: &models.ModelsBandAlarmThreshold{},
			},
			expected: BandAlarmOverallThreshold{
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
			expected: BandAlarmOverallThreshold{
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
			expected: BandAlarmOverallThreshold{
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
			expected: BandAlarmOverallThreshold{
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
			expected: BandAlarmOverallThreshold{
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
			expected: BandAlarmOverallThreshold{
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
			actual := BandAlarmOverallThreshold{}

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
		threshold := &BandAlarmOverallThreshold{}

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

func Test_BandAlarmThreshold_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var threshold *BandAlarmThreshold

		threshold.FromInternal(&models.ModelsBandAlarmThreshold{})
	})

	assert.NotPanics(t, func() {
		threshold := &BandAlarmThreshold{}

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
