package models

import (
	"time"

	"github.com/go-openapi/strfmt"

	models "github.com/SKF/go-pas-client/internal/models"
	"github.com/SKF/go-utility/v2/uuid"
)

type ContentType string

const (
	ContentTypeDataPoint       ContentType = "DATA_POINT"
	ContentTypeSpectrum        ContentType = "SPECTRUM"
	ContentTypeQuestionAnswers ContentType = "QUESTION_ANSWERS"
)

type (
	Measurement struct {
		MeasurementID   uuid.UUID
		CreatedAt       time.Time
		ContentType     ContentType
		DataPoint       *DataPoint
		Spectrum        *Spectrum
		QuestionAnswers []string
		RateOfChange    *float64
		Tags            map[string]interface{}
	}

	Coordinate struct {
		X float64
		Y float64
	}

	DataPoint struct {
		Coordinate Coordinate
		XUnit      string
		YUnit      string
	}

	Spectrum struct {
		XUnit string
		YUnit string
		Speed float64
	}
)

func (m *Measurement) ToInternal() models.ModelsUpdateAlarmStatusRequest {
	if m == nil {
		return models.ModelsUpdateAlarmStatusRequest{} // nolint:exhaustivestruct
	}

	var (
		measurementID = strfmt.UUID(m.MeasurementID.String())
		createdAt     = strfmt.DateTime(m.CreatedAt)
		contentType   = string(m.ContentType)
	)

	internal := models.ModelsUpdateAlarmStatusRequest{
		MeasurementID:   &measurementID,
		CreatedAt:       &createdAt,
		ContentType:     &contentType,
		RateOfChange:    m.RateOfChange,
		DataPoint:       nil,
		Spectrum:        nil,
		QuestionAnswers: nil,
		Tags:            m.Tags,
	}

	if length := len(m.QuestionAnswers); length > 0 {
		internal.QuestionAnswers = make([]string, length)

		copy(internal.QuestionAnswers, m.QuestionAnswers)
	}

	if m.DataPoint != nil {
		internal.DataPoint = &models.ModelsDataPoint{
			XUnit: &m.DataPoint.XUnit,
			YUnit: &m.DataPoint.YUnit,
			Coordinate: &models.ModelsCoordinate{
				X: &m.DataPoint.Coordinate.X,
				Y: &m.DataPoint.Coordinate.Y,
			},
		}
	}

	if m.Spectrum != nil {
		internal.Spectrum = &models.ModelsSpectrum{
			XUnit: &m.Spectrum.XUnit,
			YUnit: &m.Spectrum.YUnit,
			Speed: &m.Spectrum.Speed,
		}
	}

	return internal
}
