package models

import models "github.com/SKF/go-pas-client/internal/models"

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

	if internal.OverallThreshold != nil {
		b.OverallThreshold = new(BandAlarmOverallThreshold)
		b.OverallThreshold.FromInternal(internal.OverallThreshold)
	}

	b.MinFrequency.FromInternal(internal.MinFrequency)
	b.MaxFrequency.FromInternal(internal.MaxFrequency)
}

func (b *BandAlarm) ToInternal() *models.ModelsBandAlarm {
	if b == nil {
		return nil
	}

	return &models.ModelsBandAlarm{
		Label:            b.Label,
		OverallThreshold: b.OverallThreshold.ToInternal(),
		MinFrequency:     b.MinFrequency.ToInternal(),
		MaxFrequency:     b.MaxFrequency.ToInternal(),
	}
}

func (f *BandAlarmFrequency) FromInternal(internal *models.ModelsBandAlarmFrequency) {
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

func (f *BandAlarmFrequency) ToInternal() *models.ModelsBandAlarmFrequency {
	if f == nil {
		return nil
	}

	valueType := int32(f.ValueType)

	return &models.ModelsBandAlarmFrequency{
		ValueType: &valueType,
		Value:     &f.Value,
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

func (b *BandAlarmOverallThreshold) ToInternal() *models.ModelsBandAlarmOverallThreshold {
	if b == nil {
		return nil
	}

	return &models.ModelsBandAlarmOverallThreshold{
		Unit:        b.Unit,
		UpperAlert:  b.UpperAlert.ToInternal(),
		UpperDanger: b.UpperDanger.ToInternal(),
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

func (t *BandAlarmThreshold) ToInternal() *models.ModelsBandAlarmThreshold {
	if t == nil {
		return nil
	}

	valueType := int32(t.ValueType)

	return &models.ModelsBandAlarmThreshold{
		ValueType: &valueType,
		Value:     &t.Value,
	}
}
