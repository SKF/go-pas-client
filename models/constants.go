package models

type AlarmStatusType int32

const (
	AlarmStatusNotConfigured AlarmStatusType = iota
	AlarmStatusNoData
	AlarmStatusGood
	AlarmStatusAlert
	AlarmStatusDanger
)

type ThresholdType int32

const (
	ThresholdTypeNone ThresholdType = iota
	ThresholdTypeOverallInWindow
	ThresholdTypeOverallOutOfWindow
	ThresholdTypeInspection
)

type BandAlarmFrequencyValueType int32

const (
	BandAlarmFrequencyUnknown BandAlarmFrequencyValueType = iota
	BandAlarmFrequencyFixed
	BandAlarmFrequencySpeedMultiple
)

type BandAlarmThresholdType int32

const (
	BandAlarmThresholdTypeUnknown BandAlarmThresholdType = iota
	BandAlarmThresholdTypeAbsolute
	BandAlarmThresholdTypeRelativeFullscale
)

type HALAlarmType string

const (
	HALAlarmTypeGlobal         HALAlarmType = "GLOBAL"
	HALAlarmTypeFaultFrequency HALAlarmType = "FREQUENCY"
)
