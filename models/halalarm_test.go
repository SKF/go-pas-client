package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	models "github.com/SKF/go-pas-client/internal/models"
)

func Test_halAlarmFromInternal(t *testing.T) {
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
