package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/SKF/go-utility/v2/uuid"
)

func Test_ThresholdEvent_FromInternal(t *testing.T) {
	tests := []struct {
		given    []byte
		expected *ThresholdEvent
	}{
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000"
}
`),
			expected: &ThresholdEvent{
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(ThresholdEvent)

			err := actual.FromInternal(test.given)
			require.NoError(t, err)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_ThresholdEvent_FromInternal_IsNil(t *testing.T) {
	assert.NotPanics(t, func() {
		var event *ThresholdEvent

		err := event.FromInternal([]byte{})
		require.NoError(t, err)
	})
}

func Test_ThresholdEvent_FromInternal_InvalidBody(t *testing.T) {
	event := &ThresholdEvent{}

	err := event.FromInternal([]byte(`not-valid`))

	assert.Error(t, err)
}

func Test_AlarmStatusEvent_FromInternal(t *testing.T) {
	tests := []struct {
		given    []byte
		expected *AlarmStatusEvent
	}{
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"alarmStatus": 0
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Status: AlarmStatusNotConfigured,
					Band:   []BandAlarmStatus{},
					HAL:    []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"alarmStatus": 1
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Status: AlarmStatusNoData,
					Band:   []BandAlarmStatus{},
					HAL:    []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"alarmStatus": 2
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Status: AlarmStatusGood,
					Band:   []BandAlarmStatus{},
					HAL:    []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"alarmStatus": 3
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Status: AlarmStatusAlert,
					Band:   []BandAlarmStatus{},
					HAL:    []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"alarmStatus": 4
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Status: AlarmStatusDanger,
					Band:   []BandAlarmStatus{},
					HAL:    []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"overallAlarm": {
		"status": 2,
		"triggeringMeasurement": "00000000-0000-0000-0000-000000000000"
	}
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Overall: &GenericAlarmStatus{
						Status:                AlarmStatusGood,
						TriggeringMeasurement: uuid.EmptyUUID,
					},
					Band: []BandAlarmStatus{},
					HAL:  []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"rateOfChangeAlarm": {
		"status": 2,
		"triggeringMeasurement": "00000000-0000-0000-0000-000000000000"
	}
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					RateOfChange: &GenericAlarmStatus{
						Status:                AlarmStatusGood,
						TriggeringMeasurement: uuid.EmptyUUID,
					},
					Band: []BandAlarmStatus{},
					HAL:  []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"inspectionAlarm": {
		"status": 2,
		"triggeringMeasurement": "00000000-0000-0000-0000-000000000000"
	}
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Inspection: &GenericAlarmStatus{
						Status:                AlarmStatusGood,
						TriggeringMeasurement: uuid.EmptyUUID,
					},
					Band: []BandAlarmStatus{},
					HAL:  []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"externalAlarm": {
		"status": 2
	}
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					External: &ExternalAlarmStatus{
						Status: AlarmStatusGood,
					},
					Band: []BandAlarmStatus{},
					HAL:  []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"bandAlarms": [
		{
			"label": "10x RPM",
			"status": 2,
			"triggeringMeasurement": "00000000-0000-0000-0000-000000000000",
			"minFrequency": {
				"valueType": 1,
				"value": 100
			},
			"maxFrequency": {
				"valueType": 1,
				"value": 500
			},
			"calculatedOverall": {
				"unit": "gE",
				"value": 0.5
			}
		}
	]
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Band: []BandAlarmStatus{
						{
							Label: "10x RPM",
							GenericAlarmStatus: GenericAlarmStatus{
								Status:                AlarmStatusGood,
								TriggeringMeasurement: uuid.EmptyUUID,
							},
							MinFrequency: BandAlarmFrequency{
								ValueType: BandAlarmFrequencyFixed,
								Value:     100,
							},
							MaxFrequency: BandAlarmFrequency{
								ValueType: BandAlarmFrequencyFixed,
								Value:     500,
							},
							CalculatedOverall: &BandAlarmStatusCalculatedOverall{
								Unit:  "gE",
								Value: 0.5,
							},
						},
					},
					HAL: []HALAlarmStatus{},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
		{
			given: []byte(`
{
	"aggregateId": "00000000-0000-0000-0000-000000000000",
	"userId": "00000000-0000-0000-0000-000000000000",
	"halAlarms": [
		{
			"label": "10x RPM",
			"status": 2,
			"triggeringMeasurement": "00000000-0000-0000-0000-000000000000",
			"halIndex": 1.22,
			"faultFrequency": 12.22,
			"rpmFactor": 10,
			"numberOfHarmonicsUsed": 5
		}
	]
}
`),
			expected: &AlarmStatusEvent{
				AlarmStatus: AlarmStatus{
					Band: []BandAlarmStatus{},
					HAL: []HALAlarmStatus{
						{
							Label: "10x RPM",
							GenericAlarmStatus: GenericAlarmStatus{
								Status:                AlarmStatusGood,
								TriggeringMeasurement: uuid.EmptyUUID,
							},
							HALIndex:              f64p(1.22),
							FaultFrequency:        f64p(12.22),
							RPMFactor:             f64p(10),
							NumberOfHarmonicsUsed: i64p(5),
						},
					},
				},
				AggregateID: uuid.EmptyUUID,
				UserID:      uuid.EmptyUUID,
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(AlarmStatusEvent)

			err := actual.FromInternal(test.given)
			require.NoError(t, err)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_AlarmStatusEvent_FromInternal_IsNil(t *testing.T) {
	assert.NotPanics(t, func() {
		var event *AlarmStatusEvent

		event.FromInternal([]byte{})
	})
}

func Test_AlarmStatusEvent_FromInternal_InvalidBody(t *testing.T) {
	event := &AlarmStatusEvent{}

	err := event.FromInternal([]byte(`not-valid`))

	assert.Error(t, err)
}
