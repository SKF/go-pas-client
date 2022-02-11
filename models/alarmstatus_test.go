package models

import (
	"math"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"

	models "github.com/SKF/go-pas-client/internal/models"
	"github.com/SKF/go-utility/v2/uuid"
)

func Test_AlarmStatus_FromInternal(t *testing.T) {
	// This strange creation of a timestamp is to make
	// sure nanoseconds in the time struct doesn't
	// fail the equals assertion.
	now := time.UnixMilli(time.Now().UTC().UnixMilli()).UTC()

	triggeringMeasurement := strfmt.UUID(uuid.EmptyUUID.String())

	tests := []struct {
		given    models.ModelsGetAlarmStatusResponse
		expected AlarmStatus
	}{
		{
			given: models.ModelsGetAlarmStatusResponse{
				UpdatedAt: now.UnixMilli(),
			},
			expected: AlarmStatus{
				UpdatedAt: now,
				Band:      []BandAlarmStatus{},
				HAL:       []HALAlarmStatus{},
			},
		},
		{
			given: models.ModelsGetAlarmStatusResponse{
				UpdatedAt: now.UnixMilli(),
				Status:    i32p(0), // not configured
			},
			expected: AlarmStatus{
				UpdatedAt: now,
				Status:    AlarmStatusNotConfigured,
				Band:      []BandAlarmStatus{},
				HAL:       []HALAlarmStatus{},
			},
		},
		{
			given: models.ModelsGetAlarmStatusResponse{
				UpdatedAt: now.UnixMilli(),
				Status:    i32p(1), // no data
			},
			expected: AlarmStatus{
				UpdatedAt: now,
				Status:    AlarmStatusNoData,
				Band:      []BandAlarmStatus{},
				HAL:       []HALAlarmStatus{},
			},
		},
		{
			given: models.ModelsGetAlarmStatusResponse{
				UpdatedAt: now.UnixMilli(),
				Status:    i32p(2), // good
			},
			expected: AlarmStatus{
				UpdatedAt: now,
				Status:    AlarmStatusGood,
				Band:      []BandAlarmStatus{},
				HAL:       []HALAlarmStatus{},
			},
		},
		{
			given: models.ModelsGetAlarmStatusResponse{
				UpdatedAt: now.UnixMilli(),
				Status:    i32p(3), // alert
			},
			expected: AlarmStatus{
				UpdatedAt: now,
				Status:    AlarmStatusAlert,
				Band:      []BandAlarmStatus{},
				HAL:       []HALAlarmStatus{},
			},
		},
		{
			given: models.ModelsGetAlarmStatusResponse{
				UpdatedAt: now.UnixMilli(),
				Status:    i32p(4), // danger
			},
			expected: AlarmStatus{
				UpdatedAt: now,
				Status:    AlarmStatusDanger,
				Band:      []BandAlarmStatus{},
				HAL:       []HALAlarmStatus{},
			},
		},
		{
			given: models.ModelsGetAlarmStatusResponse{
				UpdatedAt: now.UnixMilli(),
				OverallAlarm: &models.ModelsGetAlarmStatusResponseGeneric{
					Status:                i32p(2), // good
					TriggeringMeasurement: strfmt.UUID(uuid.EmptyUUID.String()),
				},
			},
			expected: AlarmStatus{
				UpdatedAt: now,
				Overall: &GenericAlarmStatus{
					Status:                AlarmStatusGood,
					TriggeringMeasurement: uuid.EmptyUUID,
				},
				Band: []BandAlarmStatus{},
				HAL:  []HALAlarmStatus{},
			},
		},
		{
			given: models.ModelsGetAlarmStatusResponse{
				UpdatedAt: now.UnixMilli(),
				RateOfChangeAlarm: &models.ModelsGetAlarmStatusResponseGeneric{
					Status:                i32p(2), // good
					TriggeringMeasurement: strfmt.UUID(uuid.EmptyUUID.String()),
				},
			},
			expected: AlarmStatus{
				UpdatedAt: now,
				RateOfChange: &GenericAlarmStatus{
					Status:                AlarmStatusGood,
					TriggeringMeasurement: uuid.EmptyUUID,
				},
				Band: []BandAlarmStatus{},
				HAL:  []HALAlarmStatus{},
			},
		},
		{
			given: models.ModelsGetAlarmStatusResponse{
				UpdatedAt: now.UnixMilli(),
				InspectionAlarm: &models.ModelsGetAlarmStatusResponseGeneric{
					Status:                i32p(2), // good
					TriggeringMeasurement: strfmt.UUID(uuid.EmptyUUID.String()),
				},
			},
			expected: AlarmStatus{
				UpdatedAt: now,
				Inspection: &GenericAlarmStatus{
					Status:                AlarmStatusGood,
					TriggeringMeasurement: uuid.EmptyUUID,
				},
				Band: []BandAlarmStatus{},
				HAL:  []HALAlarmStatus{},
			},
		},
		{
			given: models.ModelsGetAlarmStatusResponse{
				UpdatedAt: now.UnixMilli(),
				ExternalAlarm: &models.ModelsGetAlarmStatusResponseExternal{
					Status: i32p(2), // good
				},
			},
			expected: AlarmStatus{
				UpdatedAt: now,
				External: &ExternalAlarmStatus{
					Status: AlarmStatusGood,
				},
				Band: []BandAlarmStatus{},
				HAL:  []HALAlarmStatus{},
			},
		},
		{
			given: models.ModelsGetAlarmStatusResponse{
				UpdatedAt: now.UnixMilli(),
				BandAlarms: []*models.ModelsGetAlarmStatusResponseBandAlarm{
					{
						Label:                 "10x RPM",
						Status:                i32p(2),
						TriggeringMeasurement: strfmt.UUID(uuid.EmptyUUID.String()),
						MinFrequency: &models.ModelsGetAlarmStatusResponseFrequency{
							ValueType: i32p(1), // fixed
							Value:     f64p(100),
						},
						MaxFrequency: &models.ModelsGetAlarmStatusResponseFrequency{
							ValueType: i32p(1), // fixed
							Value:     f64p(500),
						},
						CalculatedOverall: &models.ModelsBandCalculatedOverall{
							Unit:  "gE",
							Value: f64p(3.5),
						},
					},
					{
						Label:                 "12 TOOTH SPROCKET",
						Status:                i32p(2),
						TriggeringMeasurement: strfmt.UUID(uuid.EmptyUUID.String()),
						MinFrequency: &models.ModelsGetAlarmStatusResponseFrequency{
							ValueType: i32p(2), // speed multiple
							Value:     f64p(1.2),
						},
						MaxFrequency: &models.ModelsGetAlarmStatusResponseFrequency{
							ValueType: i32p(2), // speed multiple
							Value:     f64p(1.5),
						},
						CalculatedOverall: &models.ModelsBandCalculatedOverall{
							Unit:  "gE",
							Value: f64p(3.5),
						},
					},
				},
			},
			expected: AlarmStatus{
				UpdatedAt: now,
				Band: []BandAlarmStatus{
					{
						Label: "10x RPM",
						GenericAlarmStatus: GenericAlarmStatus{
							Status:                AlarmStatusGood,
							TriggeringMeasurement: uuid.EmptyUUID,
						},
						MinFrequency: &BandAlarmFrequency{
							ValueType: BandAlarmFrequencyFixed,
							Value:     100,
						},
						MaxFrequency: &BandAlarmFrequency{
							ValueType: BandAlarmFrequencyFixed,
							Value:     500,
						},
						CalculatedOverall: BandAlarmStatusCalculatedOverall{
							Unit:  "gE",
							Value: 3.5,
						},
					},
					{
						Label: "12 TOOTH SPROCKET",
						GenericAlarmStatus: GenericAlarmStatus{
							Status:                AlarmStatusGood,
							TriggeringMeasurement: uuid.EmptyUUID,
						},
						MinFrequency: &BandAlarmFrequency{
							ValueType: BandAlarmFrequencySpeedMultiple,
							Value:     1.2,
						},
						MaxFrequency: &BandAlarmFrequency{
							ValueType: BandAlarmFrequencySpeedMultiple,
							Value:     1.5,
						},
						CalculatedOverall: BandAlarmStatusCalculatedOverall{
							Unit:  "gE",
							Value: 3.5,
						},
					},
				},
				HAL: []HALAlarmStatus{},
			},
		},
		{
			given: models.ModelsGetAlarmStatusResponse{
				UpdatedAt: now.UnixMilli(),
				HalAlarms: []*models.ModelsGetAlarmStatusResponseHALAlarm{
					{
						Label:                 stringp("10x RPM"),
						Status:                i32p(2), // good
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
			},
			expected: AlarmStatus{
				UpdatedAt: now,
				Band:      []BandAlarmStatus{},
				HAL: []HALAlarmStatus{
					{
						Label: "10x RPM",
						GenericAlarmStatus: GenericAlarmStatus{
							Status:                AlarmStatusGood,
							TriggeringMeasurement: uuid.EmptyUUID,
						},
						HALIndex:              f64p(1.22),
						FaultFrequency:        f64p(122),
						RPMFactor:             f64p(10),
						NumberOfHarmonicsUsed: i64p(15),
					},
					{
						Label: "12 TOOTH SPROCKET",
						GenericAlarmStatus: GenericAlarmStatus{
							Status:                AlarmStatusNoData,
							TriggeringMeasurement: uuid.EmptyUUID,
						},
						FaultFrequency:   f64p(122),
						RPMFactor:        f64p(12),
						ErrorDescription: stringp("only peaks"),
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := AlarmStatus{}

			actual.FromInternal(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_AlarmStatus_FromInternal_IsNil(t *testing.T) {
	assert.NotPanics(t, func() {
		var status *AlarmStatus

		status.FromInternal(models.ModelsGetAlarmStatusResponse{})
	})
}

func Test_GenericAlarmStatus(t *testing.T) {
	tests := []struct {
		given    *models.ModelsGetAlarmStatusResponseGeneric
		expected *GenericAlarmStatus
	}{
		{
			given:    &models.ModelsGetAlarmStatusResponseGeneric{},
			expected: &GenericAlarmStatus{},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseGeneric{
				Status: i32p(0), // not configured
			},
			expected: &GenericAlarmStatus{
				Status: AlarmStatusNotConfigured,
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseGeneric{
				Status: i32p(1), // no data
			},
			expected: &GenericAlarmStatus{
				Status: AlarmStatusNoData,
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseGeneric{
				Status: i32p(2), // good
			},
			expected: &GenericAlarmStatus{
				Status: AlarmStatusGood,
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseGeneric{
				Status: i32p(3), // alert
			},
			expected: &GenericAlarmStatus{
				Status: AlarmStatusAlert,
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseGeneric{
				Status: i32p(4), // danger
			},
			expected: &GenericAlarmStatus{
				Status: AlarmStatusDanger,
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseGeneric{
				TriggeringMeasurement: strfmt.UUID(uuid.EmptyUUID.String()),
			},
			expected: &GenericAlarmStatus{
				TriggeringMeasurement: uuid.EmptyUUID,
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(GenericAlarmStatus)

			actual.FromInternal(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_GenericAlarmStatus_IsNil(t *testing.T) {
	assert.NotPanics(t, func() {
		var status *GenericAlarmStatus

		status.FromInternal(&models.ModelsGetAlarmStatusResponseGeneric{})
	})

	assert.NotPanics(t, func() {
		status := GenericAlarmStatus{}

		status.FromInternal(nil)
	})
}

func Test_ExternalAlarmStatus(t *testing.T) {
	var (
		setByStrfmt = strfmt.UUID(uuid.EmptyUUID.String())
		setBy       = uuid.EmptyUUID
	)

	tests := []struct {
		given    *models.ModelsGetAlarmStatusResponseExternal
		expected *ExternalAlarmStatus
	}{
		{
			given:    &models.ModelsGetAlarmStatusResponseExternal{},
			expected: &ExternalAlarmStatus{},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseExternal{
				Status: i32p(0), // not configured
			},
			expected: &ExternalAlarmStatus{
				Status: AlarmStatusNotConfigured,
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseExternal{
				Status: i32p(1), // no data
			},
			expected: &ExternalAlarmStatus{
				Status: AlarmStatusNoData,
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseExternal{
				Status: i32p(2), // good
			},
			expected: &ExternalAlarmStatus{
				Status: AlarmStatusGood,
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseExternal{
				Status: i32p(3), // alert
			},
			expected: &ExternalAlarmStatus{
				Status: AlarmStatusAlert,
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseExternal{
				Status: i32p(4), // danger
			},
			expected: &ExternalAlarmStatus{
				Status: AlarmStatusDanger,
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseExternal{
				SetBy: &setByStrfmt,
			},
			expected: &ExternalAlarmStatus{
				SetBy: &setBy,
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(ExternalAlarmStatus)

			actual.FromInternal(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_ExternalAlarmStatus_IsNil(t *testing.T) {
	assert.NotPanics(t, func() {
		var status *ExternalAlarmStatus

		status.FromInternal(&models.ModelsGetAlarmStatusResponseExternal{})
	})

	assert.NotPanics(t, func() {
		status := ExternalAlarmStatus{}

		status.FromInternal(nil)
	})
}

func Test_BandAlarmStatus(t *testing.T) {
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
				MinFrequency: &BandAlarmFrequency{
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
				MinFrequency: &BandAlarmFrequency{
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
				MaxFrequency: &BandAlarmFrequency{
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
				MaxFrequency: &BandAlarmFrequency{
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
				CalculatedOverall: BandAlarmStatusCalculatedOverall{
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

func Test_BandAlarmStatus_IsNil(t *testing.T) {
	assert.NotPanics(t, func() {
		var status *BandAlarmStatus

		status.FromInternal(&models.ModelsGetAlarmStatusResponseBandAlarm{})
	})

	assert.NotPanics(t, func() {
		status := BandAlarmStatus{}

		status.FromInternal(nil)
	})
}

func Test_HALAlarmStatus(t *testing.T) {
	triggeringMeasurement := strfmt.UUID(uuid.EmptyUUID.String())

	tests := []struct {
		given    *models.ModelsGetAlarmStatusResponseHALAlarm
		expected *HALAlarmStatus
	}{
		{
			given:    &models.ModelsGetAlarmStatusResponseHALAlarm{},
			expected: &HALAlarmStatus{},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				Label: stringp("10x RPM"),
			},
			expected: &HALAlarmStatus{
				Label: "10x RPM",
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				Status: i32p(0), // not configured
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusNotConfigured,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				Status: i32p(1), // no data
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusNoData,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				Status: i32p(2), // good
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusGood,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				Status: i32p(3), // alert
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusAlert,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				Status: i32p(4), // danger
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					Status: AlarmStatusDanger,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				TriggeringMeasurement: &triggeringMeasurement,
			},
			expected: &HALAlarmStatus{
				GenericAlarmStatus: GenericAlarmStatus{
					TriggeringMeasurement: uuid.EmptyUUID,
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				Bearing: &models.ModelsBearing{
					Manufacturer: stringp("SKF"),
					ModelNumber:  stringp("2222"),
				},
			},
			expected: &HALAlarmStatus{
				Bearing: &Bearing{
					Manufacturer: "SKF",
					ModelNumber:  "2222",
				},
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				HalIndex: f64p(math.Pi),
			},
			expected: &HALAlarmStatus{
				HALIndex: f64p(math.Pi),
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				FaultFrequency: f64p(math.Pi),
			},
			expected: &HALAlarmStatus{
				FaultFrequency: f64p(math.Pi),
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				NumberOfHarmonicsUsed: i64p(25),
			},
			expected: &HALAlarmStatus{
				NumberOfHarmonicsUsed: i64p(25),
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				RpmFactor: f64p(10),
			},
			expected: &HALAlarmStatus{
				RPMFactor: f64p(10),
			},
		},
		{
			given: &models.ModelsGetAlarmStatusResponseHALAlarm{
				ErrorDescription: stringp("only peaks"),
			},
			expected: &HALAlarmStatus{
				ErrorDescription: stringp("only peaks"),
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := new(HALAlarmStatus)

			actual.FromInternal(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_HALAlarmStatus_IsNil(t *testing.T) {
	assert.NotPanics(t, func() {
		var status *HALAlarmStatus

		status.FromInternal(&models.ModelsGetAlarmStatusResponseHALAlarm{})
	})

	assert.NotPanics(t, func() {
		status := HALAlarmStatus{}

		status.FromInternal(nil)
	})
}
