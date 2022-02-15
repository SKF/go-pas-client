package models

import (
	"encoding/json"
	"fmt"

	"github.com/SKF/go-pas-client/internal/events"
	"github.com/SKF/go-utility/v2/uuid"
)

const (
	EventAttributeEventType = "SKF.Hierarchy.EventType"
	EventAttributeAggregate = "SKF.Hierarchy.Aggregate"
)

const (
	EventTypeThreshold   = "SetPointAlarmThresholdEvent"
	EventTypeAlarmStatus = "PointAlarmStatusEvent"
)

type ThresholdEvent struct {
	AggregateID uuid.UUID
	UserID      uuid.UUID
	Threshold   Threshold
}

func (t *ThresholdEvent) FromInternal(buf []byte) error {
	if t == nil {
		return nil
	}

	var internal events.SetPointAlarmThresholdEvent

	if err := json.Unmarshal(buf, &internal); err != nil {
		return fmt.Errorf("failed to decode event: %w", err)
	}

	t.AggregateID = uuid.UUID(internal.BaseEvent.AggregateID)
	t.UserID = uuid.UUID(internal.BaseEvent.UserID)

	return nil
}

type AlarmStatusEvent struct {
	AggregateID uuid.UUID
	UserID      uuid.UUID
	Changed     bool
	AlarmStatus AlarmStatus
}

func (a *AlarmStatusEvent) FromInternal(buf []byte) error {
	if a == nil {
		return nil
	}

	var internal events.PointAlarmStatusEvent

	if err := json.Unmarshal(buf, &internal); err != nil {
		return fmt.Errorf("failed to decode event: %w", err)
	}

	a.AlarmStatus.Status = AlarmStatusType(internal.AlarmStatus)
	a.AggregateID = uuid.UUID(internal.BaseEvent.AggregateID)
	a.UserID = uuid.UUID(internal.BaseEvent.UserID)
	a.Changed = internal.AlarmsChanged

	if internal.OverallAlarm != nil {
		a.AlarmStatus.Overall = new(GenericAlarmStatus)
		a.AlarmStatus.Overall.FromEvent(internal.OverallAlarm)
	}

	if internal.RateOfChangeAlarm != nil {
		a.AlarmStatus.RateOfChange = new(GenericAlarmStatus)
		a.AlarmStatus.RateOfChange.FromEvent(internal.RateOfChangeAlarm)
	}

	if internal.InspectionAlarm != nil {
		a.AlarmStatus.Inspection = new(GenericAlarmStatus)
		a.AlarmStatus.Inspection.FromEvent(internal.InspectionAlarm)
	}

	if internal.ExternalAlarm != nil {
		a.AlarmStatus.External = new(ExternalAlarmStatus)
		a.AlarmStatus.External.FromEvent(internal.ExternalAlarm)
	}

	a.AlarmStatus.Band = make([]BandAlarmStatus, len(internal.BandAlarms))

	for i, status := range internal.BandAlarms {
		a.AlarmStatus.Band[i].FromEvent(status)
	}

	a.AlarmStatus.HAL = make([]HALAlarmStatus, len(internal.HalAlarms))

	for i, status := range internal.HalAlarms {
		a.AlarmStatus.HAL[i].FromEvent(status)
	}

	return nil
}
