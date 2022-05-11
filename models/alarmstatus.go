package models

import (
	"time"

	s "github.com/go-openapi/strfmt"

	"github.com/SKF/go-pas-client/internal/events"
	"github.com/SKF/go-pas-client/internal/models"
	"github.com/SKF/go-utility/v2/uuid"
)

type (
	AlarmStatus struct {
		Status       AlarmStatusType
		UpdatedAt    time.Time
		Overall      *GenericAlarmStatus
		RateOfChange *GenericAlarmStatus
		Inspection   *GenericAlarmStatus
		Band         []BandAlarmStatus
		HAL          []HALAlarmStatus
		External     *ExternalAlarmStatus
	}

	GenericAlarmStatus struct {
		TriggeringMeasurement uuid.UUID
		Status                AlarmStatusType
	}

	ExternalAlarmStatus struct {
		Status AlarmStatusType
		SetBy  *uuid.UUID
	}
)

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

func (g *GenericAlarmStatus) FromEvent(internal *events.GenericAlarm) {
	if g == nil || internal == nil {
		return
	}

	g.Status = AlarmStatusType(internal.Status)
	g.TriggeringMeasurement = internal.TriggeringMeasurement
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

func (e *ExternalAlarmStatus) FromEvent(internal *events.ExternalAlarm) {
	if e == nil || internal == nil {
		return
	}

	e.Status = AlarmStatusType(internal.Status)
	e.SetBy = internal.SetBy
}

func (e *ExternalAlarmStatus) ToSetRequest() models.ModelsSetExternalAlarmStatusRequest {
	if e == nil {
		return models.ModelsSetExternalAlarmStatusRequest{} // nolint:exhaustivestruct
	}

	status := int32(e.Status)

	request := models.ModelsSetExternalAlarmStatusRequest{
		Status: &status,
		SetBy:  nil,
	}

	if e.SetBy != nil {
		setBy := s.UUID(e.SetBy.String())

		request.SetBy = &setBy
	}

	return request
}
