package models

import (
	"fmt"

	"google.golang.org/protobuf/proto" //nolint:gci

	"github.com/SKF/go-pas-client/internal/events"
	models "github.com/SKF/go-pas-client/internal/models"
	"github.com/SKF/go-utility/v2/uuid"
	pas "github.com/SKF/proto/v2/pas"
)

type (
	BandAlarm struct {
		Label            string
		MinFrequency     BandAlarmFrequency
		MaxFrequency     BandAlarmFrequency
		OverallThreshold *BandAlarmOverallThreshold
	}

	BandAlarmFrequency struct {
		ValueType BandAlarmFrequencyValueType
		Value     float64
	}

	BandAlarmThreshold struct {
		ValueType BandAlarmThresholdType
		Value     float64
	}

	BandAlarmOverallThreshold struct {
		Unit        string
		UpperAlert  *BandAlarmThreshold
		UpperDanger *BandAlarmThreshold
	}
)

func (b *BandAlarm) FromInternal(internal *models.ModelsBandAlarm) {
	if b == nil || internal == nil {
		return
	}

	b.Label = internal.Label
	b.MinFrequency.FromInternalThreshold(internal.MinFrequency)
	b.MaxFrequency.FromInternalThreshold(internal.MaxFrequency)

	if internal.OverallThreshold != nil {
		b.OverallThreshold = new(BandAlarmOverallThreshold)
		b.OverallThreshold.FromInternal(internal.OverallThreshold)
	}
}

func (b BandAlarm) ToInternal() *models.ModelsBandAlarm {
	bandAlarm := &models.ModelsBandAlarm{
		Label:            b.Label,
		OverallThreshold: nil,
		MinFrequency:     b.MinFrequency.ToInternal(),
		MaxFrequency:     b.MaxFrequency.ToInternal(),
	}

	if b.OverallThreshold != nil {
		bandAlarm.OverallThreshold = b.OverallThreshold.ToInternal()
	}

	return bandAlarm
}

func (b *BandAlarm) FromProto(buf []byte) error {
	if b == nil {
		return nil
	}

	var internal pas.BandAlarm

	if err := proto.Unmarshal(buf, &internal); err != nil {
		return fmt.Errorf("decoding band alarm failed: %w", err)
	}

	b.Label = internal.Label
	b.MinFrequency.FromProto(internal.MinFrequency)
	b.MaxFrequency.FromProto(internal.MaxFrequency)

	if internal.OverallThreshold != nil {
		b.OverallThreshold = new(BandAlarmOverallThreshold)
		b.OverallThreshold.FromProto(internal.OverallThreshold)
	}

	return nil
}

func (f *BandAlarmFrequency) FromInternalThreshold(internal *models.ModelsBandAlarmFrequency) {
	if f == nil || internal == nil {
		return
	}

	if internal.ValueType != nil {
		f.ValueType = BandAlarmFrequencyValueType(*internal.ValueType)
	}

	if internal.Value != nil {
		f.Value = *internal.Value
	}
}

func (f *BandAlarmFrequency) FromInternalAlarmStatus(internal *models.ModelsGetAlarmStatusResponseFrequency) {
	if f == nil || internal == nil {
		return
	}

	if internal.ValueType != nil {
		f.ValueType = BandAlarmFrequencyValueType(*internal.ValueType)
	}

	if internal.Value != nil {
		f.Value = *internal.Value
	}
}

func (f BandAlarmFrequency) ToInternal() *models.ModelsBandAlarmFrequency {
	valueType := int32(f.ValueType)

	return &models.ModelsBandAlarmFrequency{
		ValueType: &valueType,
		Value:     &f.Value,
	}
}

func (f *BandAlarmFrequency) FromProto(internal *pas.Frequency) {
	if f == nil || internal == nil {
		return
	}

	f.ValueType = BandAlarmFrequencyValueType(internal.ValueType)

	if internal.Value != nil {
		f.Value = internal.Value.Value
	}
}

func (b *BandAlarmOverallThreshold) FromInternal(internal *models.ModelsBandAlarmOverallThreshold) {
	if b == nil || internal == nil {
		return
	}

	b.Unit = internal.Unit

	if internal.UpperAlert != nil {
		b.UpperAlert = new(BandAlarmThreshold)
		b.UpperAlert.FromInternal(internal.UpperAlert)
	}

	if internal.UpperDanger != nil {
		b.UpperDanger = new(BandAlarmThreshold)
		b.UpperDanger.FromInternal(internal.UpperDanger)
	}
}

func (b BandAlarmOverallThreshold) ToInternal() *models.ModelsBandAlarmOverallThreshold {
	threshold := &models.ModelsBandAlarmOverallThreshold{
		Unit:        b.Unit,
		UpperAlert:  nil,
		UpperDanger: nil,
	}

	if b.UpperAlert != nil {
		threshold.UpperAlert = b.UpperAlert.ToInternal()
	}

	if b.UpperDanger != nil {
		threshold.UpperDanger = b.UpperDanger.ToInternal()
	}

	return threshold
}

func (b *BandAlarmOverallThreshold) FromProto(internal *pas.BandAlarmOverallThreshold) {
	if b == nil || internal == nil {
		return
	}

	b.Unit = internal.Unit

	if internal.UpperAlert != nil {
		b.UpperAlert = new(BandAlarmThreshold)
		b.UpperAlert.FromProto(internal.UpperAlert)
	}

	if internal.UpperDanger != nil {
		b.UpperDanger = new(BandAlarmThreshold)
		b.UpperDanger.FromProto(internal.UpperDanger)
	}
}

func (t *BandAlarmThreshold) FromInternal(internal *models.ModelsBandAlarmThreshold) {
	if t == nil || internal == nil {
		return
	}

	if internal.ValueType != nil {
		t.ValueType = BandAlarmThresholdType(*internal.ValueType)
	}

	if internal.Value != nil {
		t.Value = *internal.Value
	}
}

func (t BandAlarmThreshold) ToInternal() *models.ModelsBandAlarmThreshold {
	valueType := int32(t.ValueType)

	return &models.ModelsBandAlarmThreshold{
		ValueType: &valueType,
		Value:     &t.Value,
	}
}

func (t *BandAlarmThreshold) FromProto(internal *pas.ThresholdValue) {
	if t == nil || internal == nil {
		return
	}

	t.ValueType = BandAlarmThresholdType(internal.ValueType)

	if internal.Value != nil {
		t.Value = internal.Value.Value
	}
}

type (
	BandAlarmStatus struct {
		GenericAlarmStatus
		Label             string
		MinFrequency      BandAlarmFrequency
		MaxFrequency      BandAlarmFrequency
		CalculatedOverall *BandAlarmStatusCalculatedOverall
	}

	BandAlarmStatusCalculatedOverall struct {
		Unit  string
		Value float64
	}
)

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
		b.MinFrequency.FromInternalAlarmStatus(internal.MinFrequency)
	}

	if internal.MaxFrequency != nil {
		b.MaxFrequency.FromInternalAlarmStatus(internal.MaxFrequency)
	}

	if internal.CalculatedOverall != nil {
		b.CalculatedOverall = &BandAlarmStatusCalculatedOverall{
			Unit:  internal.CalculatedOverall.Unit,
			Value: 0,
		}

		if internal.CalculatedOverall.Value != nil {
			b.CalculatedOverall.Value = *internal.CalculatedOverall.Value
		}
	}
}

func (b *BandAlarmStatus) FromEvent(internal events.BandAlarmStatus) {
	if b == nil {
		return
	}

	b.Label = internal.Label
	b.Status = AlarmStatusType(internal.Status)
	b.TriggeringMeasurement = internal.TriggeringMeasurement

	b.MinFrequency = BandAlarmFrequency{
		ValueType: BandAlarmFrequencyValueType(internal.MinFrequency.ValueType),
		Value:     internal.MinFrequency.Value,
	}

	b.MaxFrequency = BandAlarmFrequency{
		ValueType: BandAlarmFrequencyValueType(internal.MaxFrequency.ValueType),
		Value:     internal.MaxFrequency.Value,
	}

	if internal.CalculatedOverall != nil {
		b.CalculatedOverall = &BandAlarmStatusCalculatedOverall{
			Unit:  internal.CalculatedOverall.Unit,
			Value: internal.CalculatedOverall.Value,
		}
	}
}
