package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	models "github.com/SKF/go-pas-client/internal/models"
)

func Test_HALAlarm_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *models.ModelsHALAlarm
		expected HALAlarm
	}{
		{
			given:    nil,
			expected: HALAlarm{},
		},
		{
			given:    &models.ModelsHALAlarm{},
			expected: HALAlarm{},
		},
		{
			given: &models.ModelsHALAlarm{
				HalAlarmType: models.ModelsHALAlarmHalAlarmTypeGLOBAL,
			},
			expected: HALAlarm{
				HALAlarmType: HALAlarmTypeGlobal,
			},
		},
		{
			given: &models.ModelsHALAlarm{
				HalAlarmType: models.ModelsHALAlarmHalAlarmTypeFREQUENCY,
			},
			expected: HALAlarm{
				HALAlarmType: HALAlarmTypeFaultFrequency,
			},
		},
		{
			given: &models.ModelsHALAlarm{
				Bearing: &models.ModelsBearing{
					Manufacturer: stringp("SKF"),
					ModelNumber:  stringp("2222"),
				},
			},
			expected: HALAlarm{
				Bearing: &Bearing{
					Manufacturer: "SKF",
					ModelNumber:  "2222",
				},
			},
		},
		{
			given: &models.ModelsHALAlarm{
				UpperAlert: f64p(10),
			},
			expected: HALAlarm{
				UpperAlert: f64p(10),
			},
		},
		{
			given: &models.ModelsHALAlarm{
				UpperDanger: f64p(20),
			},
			expected: HALAlarm{
				UpperDanger: f64p(20),
			},
		},
		{
			given: &models.ModelsHALAlarm{
				UpperAlert:  f64p(10),
				UpperDanger: f64p(20),
			},
			expected: HALAlarm{
				UpperAlert:  f64p(10),
				UpperDanger: f64p(20),
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := HALAlarm{}

			actual.FromInternal(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_HALAlarm_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var halAlarm *HALAlarm

		halAlarm.FromInternal(&models.ModelsHALAlarm{})
	})

	assert.NotPanics(t, func() {
		halAlarm := &HALAlarm{}

		halAlarm.FromInternal(nil)
	})
}

func Test_HALAlarm_ToInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *HALAlarm
		expected *models.ModelsHALAlarm
	}{
		{
			given:    nil,
			expected: nil,
		},
		{
			given:    &HALAlarm{},
			expected: &models.ModelsHALAlarm{},
		},
		{
			given: &HALAlarm{
				HALAlarmType: HALAlarmTypeGlobal,
			},
			expected: &models.ModelsHALAlarm{
				HalAlarmType: models.ModelsHALAlarmHalAlarmTypeGLOBAL,
			},
		},
		{
			given: &HALAlarm{
				HALAlarmType: HALAlarmTypeFaultFrequency,
			},
			expected: &models.ModelsHALAlarm{
				HalAlarmType: models.ModelsHALAlarmHalAlarmTypeFREQUENCY,
			},
		},
		{
			given: &HALAlarm{
				Bearing: &Bearing{
					Manufacturer: "SKF",
					ModelNumber:  "2222",
				},
			},
			expected: &models.ModelsHALAlarm{
				Bearing: &models.ModelsBearing{
					Manufacturer: stringp("SKF"),
					ModelNumber:  stringp("2222"),
				},
			},
		},
		{
			given: &HALAlarm{
				UpperAlert: f64p(10),
			},
			expected: &models.ModelsHALAlarm{
				UpperAlert: f64p(10),
			},
		},
		{
			given: &HALAlarm{
				UpperDanger: f64p(20),
			},
			expected: &models.ModelsHALAlarm{
				UpperDanger: f64p(20),
			},
		},
		{
			given: &HALAlarm{
				UpperAlert:  f64p(10),
				UpperDanger: f64p(20),
			},
			expected: &models.ModelsHALAlarm{
				UpperAlert:  f64p(10),
				UpperDanger: f64p(20),
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

func Test_HALAlarm_ToInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var halAlarm *HALAlarm

		actual := halAlarm.ToInternal()

		assert.Nil(t, actual)
	})
}
