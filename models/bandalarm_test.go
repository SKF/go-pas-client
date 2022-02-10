package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	models "github.com/SKF/go-pas-client/internal/models"
)

func Test_bandAlarmFromInternal(t *testing.T) {
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
