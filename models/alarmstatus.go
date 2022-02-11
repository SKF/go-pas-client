package models

import (
	"time"

	models "github.com/SKF/go-pas-client/internal/models"
	"github.com/SKF/go-utility/v2/uuid"
)

type AlarmStatus struct {
	Status       AlarmStatusType
	UpdatedAt    time.Time
	Overall      *GenericAlarmStatus
	RateOfChange *GenericAlarmStatus
	Inspection   *GenericAlarmStatus
	Band         []BandAlarmStatus
	HAL          []HALAlarmStatus
	External     *ExternalAlarmStatus
}

type GenericAlarmStatus struct {
	TriggeringMeasurement uuid.UUID
	Status                AlarmStatusType
}

type ExternalAlarmStatus struct {
	Status AlarmStatusType
	SetBy  *uuid.UUID
}

type BandAlarmStatus struct {
	GenericAlarmStatus
	Label             string
	MinFrequency      *BandAlarmFrequency
	MaxFrequency      *BandAlarmFrequency
	CalculatedOverall BandAlarmStatusCalculatedOverall
}

type BandAlarmStatusCalculatedOverall struct {
	Unit  string
	Value float64
}

type HALAlarmStatus struct {
	GenericAlarmStatus
	Label                 string
	Bearing               *Bearing
	HALIndex              *float64
	FaultFrequency        *float64
	RPMFactor             *float64
	NumberOfHarmonicsUsed *int64
	ErrorDescription      *string
}

func (a *AlarmStatus) FromInternal(internal models.ModelsGetAlarmStatusResponse) {
	if a == nil {
		return
	}

	if internal.Status != nil {
		a.Status = AlarmStatusType(*internal.Status)
	}

	a.UpdatedAt = time.UnixMilli(internal.UpdatedAt).UTC()

	if internal.OverallAlarm != nil {
		a.Overall = new(GenericAlarmStatus)
		a.Overall.FromInternal(internal.OverallAlarm)
	}

	if internal.RateOfChangeAlarm != nil {
		a.RateOfChange = new(GenericAlarmStatus)
		a.RateOfChange.FromInternal(internal.RateOfChangeAlarm)
	}

	if internal.InspectionAlarm != nil {
		a.Inspection = new(GenericAlarmStatus)
		a.Inspection.FromInternal(internal.InspectionAlarm)
	}

	if internal.ExternalAlarm != nil {
		a.External = new(ExternalAlarmStatus)
		a.External.FromInternal(internal.ExternalAlarm)
	}

	a.Band = make([]BandAlarmStatus, len(internal.BandAlarms))

	for i, status := range internal.BandAlarms {
		a.Band[i].FromInternal(status)
	}

	a.HAL = make([]HALAlarmStatus, len(internal.HalAlarms))

	for i, status := range internal.HalAlarms {
		a.HAL[i].FromInternal(status)
	}
}

func (g *GenericAlarmStatus) FromInternal(internal *models.ModelsGetAlarmStatusResponseGeneric) {
	if g == nil || internal == nil {
		return
	}

	if internal.Status != nil {
		g.Status = AlarmStatusType(*internal.Status)
	}

	g.TriggeringMeasurement = uuid.UUID(internal.TriggeringMeasurement.String())
}

func (e *ExternalAlarmStatus) FromInternal(internal *models.ModelsGetAlarmStatusResponseExternal) {
	if e == nil || internal == nil {
		return
	}

	if internal.Status != nil {
		e.Status = AlarmStatusType(*internal.Status)
	}

	if internal.SetBy != nil {
		setBy := uuid.UUID(internal.SetBy.String())
		e.SetBy = &setBy
	}
}

func (b *BandAlarmStatus) FromInternal(internal *models.ModelsGetAlarmStatusResponseBandAlarm) {
	if b == nil || internal == nil {
		return
	}

	b.Label = internal.Label

	if internal.Status != nil {
		b.Status = AlarmStatusType(*internal.Status)
	}

	b.TriggeringMeasurement = uuid.UUID(internal.TriggeringMeasurement.String())

	if internal.MinFrequency != nil {
		b.MinFrequency = new(BandAlarmFrequency)
		b.MinFrequency.FromInternalAlarmStatus(internal.MinFrequency)
	}

	if internal.MaxFrequency != nil {
		b.MaxFrequency = new(BandAlarmFrequency)
		b.MaxFrequency.FromInternalAlarmStatus(internal.MaxFrequency)
	}

	if internal.CalculatedOverall != nil {
		b.CalculatedOverall.Unit = internal.CalculatedOverall.Unit

		if internal.CalculatedOverall.Value != nil {
			b.CalculatedOverall.Value = *internal.CalculatedOverall.Value
		}
	}
}

func (h *HALAlarmStatus) FromInternal(internal *models.ModelsGetAlarmStatusResponseHALAlarm) {
	if h == nil || internal == nil {
		return
	}

	if internal.Label != nil {
		h.Label = *internal.Label
	}

	if internal.Status != nil {
		h.Status = AlarmStatusType(*internal.Status)
	}

	if internal.TriggeringMeasurement != nil {
		h.TriggeringMeasurement = uuid.UUID(internal.TriggeringMeasurement.String())
	}

	if internal.Bearing != nil {
		h.Bearing = &Bearing{
			Manufacturer: *internal.Bearing.Manufacturer,
			ModelNumber:  *internal.Bearing.ModelNumber,
		}
	}

	h.FaultFrequency = internal.FaultFrequency
	h.RPMFactor = internal.RpmFactor
	h.HALIndex = internal.HalIndex
	h.NumberOfHarmonicsUsed = internal.NumberOfHarmonicsUsed
	h.ErrorDescription = internal.ErrorDescription
}
