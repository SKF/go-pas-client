package models

import (
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"

	"github.com/SKF/go-pas-client/internal/events"
	models "github.com/SKF/go-pas-client/internal/models"
	"github.com/SKF/go-utility/v2/uuid"
)

func Test_AlarmStatus_FromInternal(t *testing.T) {
	t.Parallel()

	// This strange creation of a timestamp is to make
	// sure nanoseconds in the time struct doesn't
	// fail the equals assertion.
	now := time.UnixMilli(time.Now().UTC().UnixMilli()).UTC()

	triggeringMeasurement := strfmt.UUID(uuid.EmptyUUID.String())

	tests := []struct {
		given    models.ModelsGetAlarmStatusResponse
		expected *AlarmStatus
	}{
		{
			given: models.ModelsGetAlarmStatusResponse{
				UpdatedAt: now.UnixMilli(),
			},
			expected: &AlarmStatus{
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
			expected: &AlarmStatus{
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
			expected: &AlarmStatus{
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
			expected: &AlarmStatus{
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
			expected: &AlarmStatus{
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
			expected: &AlarmStatus{
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
			expected: &AlarmStatus{
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
			expected: &AlarmStatus{
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
			expected: &AlarmStatus{
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
			expected: &AlarmStatus{
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
			expected: &AlarmStatus{
				UpdatedAt: now,
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
							Value: 3.5,
						},
					},
					{
						Label: "12 TOOTH SPROCKET",
						GenericAlarmStatus: GenericAlarmStatus{
							Status:                AlarmStatusGood,
							TriggeringMeasurement: uuid.EmptyUUID,
						},
						MinFrequency: BandAlarmFrequency{
							ValueType: BandAlarmFrequencySpeedMultiple,
							Value:     1.2,
						},
						MaxFrequency: BandAlarmFrequency{
							ValueType: BandAlarmFrequencySpeedMultiple,
							Value:     1.5,
						},
						CalculatedOverall: &BandAlarmStatusCalculatedOverall{
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
			expected: &AlarmStatus{
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
			actual := new(AlarmStatus)

			actual.FromInternal(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_AlarmStatus_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var status *AlarmStatus

		status.FromInternal(models.ModelsGetAlarmStatusResponse{})
	})
}

func Test_GenericAlarmStatus_FromInternal(t *testing.T) {
	t.Parallel()

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

func Test_GenericAlarmStatus_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var status *GenericAlarmStatus

		status.FromInternal(&models.ModelsGetAlarmStatusResponseGeneric{})
	})

	assert.NotPanics(t, func() {
		status := new(GenericAlarmStatus)

		status.FromInternal(nil)
	})
}

func Test_GenericAlarmStatus_FromEvent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		given    *events.GenericAlarm
		expected *GenericAlarmStatus
	}{
		{
			given:    &events.GenericAlarm{},
			expected: &GenericAlarmStatus{},
		},
		{
			given: &events.GenericAlarm{
				Status: 0, // not configured
			},
			expected: &GenericAlarmStatus{
				Status: AlarmStatusNotConfigured,
			},
		},
		{
			given: &events.GenericAlarm{
				Status: 1, // no data
			},
			expected: &GenericAlarmStatus{
				Status: AlarmStatusNoData,
			},
		},
		{
			given: &events.GenericAlarm{
				Status: 2, // good
			},
			expected: &GenericAlarmStatus{
				Status: AlarmStatusGood,
			},
		},
		{
			given: &events.GenericAlarm{
				Status: 3, // alert
			},
			expected: &GenericAlarmStatus{
				Status: AlarmStatusAlert,
			},
		},
		{
			given: &events.GenericAlarm{
				Status: 4, // danger
			},
			expected: &GenericAlarmStatus{
				Status: AlarmStatusDanger,
			},
		},
		{
			given: &events.GenericAlarm{
				TriggeringMeasurement: uuid.EmptyUUID,
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

			actual.FromEvent(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_GenericAlarmStatus_FromEvent_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var status *GenericAlarmStatus

		status.FromEvent(&events.GenericAlarm{})
	})

	assert.NotPanics(t, func() {
		status := new(GenericAlarmStatus)

		status.FromEvent(nil)
	})
}

func Test_ExternalAlarmStatus_FromInternal(t *testing.T) {
	t.Parallel()

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

func Test_ExternalAlarmStatus_FromInternal_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var status *ExternalAlarmStatus

		status.FromInternal(&models.ModelsGetAlarmStatusResponseExternal{})
	})

	assert.NotPanics(t, func() {
		status := new(ExternalAlarmStatus)

		status.FromInternal(nil)
	})
}

func Test_ExternalAlarmStatus_FromEvent(t *testing.T) {
	t.Parallel()

	setBy := uuid.EmptyUUID

	tests := []struct {
		given    *events.ExternalAlarm
		expected *ExternalAlarmStatus
	}{
		{
			given:    &events.ExternalAlarm{},
			expected: &ExternalAlarmStatus{},
		},
		{
			given: &events.ExternalAlarm{
				Status: 0, // not configured
			},
			expected: &ExternalAlarmStatus{
				Status: AlarmStatusNotConfigured,
			},
		},
		{
			given: &events.ExternalAlarm{
				Status: 1, // no data
			},
			expected: &ExternalAlarmStatus{
				Status: AlarmStatusNoData,
			},
		},
		{
			given: &events.ExternalAlarm{
				Status: 2, // good
			},
			expected: &ExternalAlarmStatus{
				Status: AlarmStatusGood,
			},
		},
		{
			given: &events.ExternalAlarm{
				Status: 3, // alert
			},
			expected: &ExternalAlarmStatus{
				Status: AlarmStatusAlert,
			},
		},
		{
			given: &events.ExternalAlarm{
				Status: 4, // danger
			},
			expected: &ExternalAlarmStatus{
				Status: AlarmStatusDanger,
			},
		},
		{
			given: &events.ExternalAlarm{
				SetBy: &setBy,
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

			actual.FromEvent(test.given)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_ExternalAlarmStatus_FromEvent_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var status *ExternalAlarmStatus

		status.FromEvent(&events.ExternalAlarm{})
	})

	assert.NotPanics(t, func() {
		status := new(ExternalAlarmStatus)

		status.FromEvent(nil)
	})
}

func Test_ExternalAlarmStatus_ToSetRequest(t *testing.T) {
	t.Parallel()

	var (
		setByStrfmt = strfmt.UUID(uuid.EmptyUUID.String())
		setBy       = uuid.EmptyUUID
	)

	tests := []struct {
		given    *ExternalAlarmStatus
		expected models.ModelsSetExternalAlarmStatusRequest
	}{
		{
			given: &ExternalAlarmStatus{
				Status: AlarmStatusNotConfigured,
			},
			expected: models.ModelsSetExternalAlarmStatusRequest{
				Status: i32p(0), // not configured
			},
		},
		{
			given: &ExternalAlarmStatus{
				Status: AlarmStatusNoData,
			},
			expected: models.ModelsSetExternalAlarmStatusRequest{
				Status: i32p(1), // no data
			},
		},
		{
			given: &ExternalAlarmStatus{
				Status: AlarmStatusGood,
			},
			expected: models.ModelsSetExternalAlarmStatusRequest{
				Status: i32p(2), // good
			},
		},
		{
			given: &ExternalAlarmStatus{
				Status: AlarmStatusAlert,
			},
			expected: models.ModelsSetExternalAlarmStatusRequest{
				Status: i32p(3), // alert
			},
		},
		{
			given: &ExternalAlarmStatus{
				Status: AlarmStatusDanger,
			},
			expected: models.ModelsSetExternalAlarmStatusRequest{
				Status: i32p(4), // danger
			},
		},
		{
			given: &ExternalAlarmStatus{
				Status: AlarmStatusGood,
				SetBy:  &setBy,
			},
			expected: models.ModelsSetExternalAlarmStatusRequest{
				Status: i32p(2), // good
				SetBy:  &setByStrfmt,
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run("", func(t *testing.T) {
			actual := test.given.ToSetRequest()

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_ExternalAlarmStatus_ToSetRequest_IsNil(t *testing.T) {
	t.Parallel()

	assert.NotPanics(t, func() {
		var status *ExternalAlarmStatus

		_ = status.ToSetRequest()
	})
}
