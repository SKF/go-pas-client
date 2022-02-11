package models

import models "github.com/SKF/go-pas-client/internal/models"

type (
	Inspection struct {
		Choices []InspectionChoice
	}

	InspectionChoice struct {
		Answer      string
		Instruction string
		Status      AlarmStatusType
	}
)

func (i *Inspection) FromInternal(internal *models.ModelsInspection) {
	if i == nil || internal == nil {
		return
	}

	i.Choices = make([]InspectionChoice, len(internal.Choices))

	for idx, choice := range internal.Choices {
		i.Choices[idx].FromInternal(choice)
	}
}

func (i *InspectionChoice) FromInternal(internal *models.ModelsInspectionChoice) {
	if i == nil || internal == nil {
		return
	}

	i.Answer = internal.Answer
	i.Instruction = internal.Instruction

	if internal.Status != nil {
		i.Status = AlarmStatusType(*internal.Status)
	}
}

func (i *Inspection) ToInternal() *models.ModelsInspection {
	if i == nil {
		return nil
	}

	inspection := &models.ModelsInspection{
		Choices: make([]*models.ModelsInspectionChoice, 0, len(i.Choices)),
	}

	for _, choice := range i.Choices {
		inspection.Choices = append(inspection.Choices, choice.ToInternal())
	}

	return inspection
}

func (i *InspectionChoice) ToInternal() *models.ModelsInspectionChoice {
	if i == nil {
		return nil
	}

	status := int32(i.Status)

	return &models.ModelsInspectionChoice{
		Answer:      i.Answer,
		Instruction: i.Instruction,
		Status:      &status,
	}
}
