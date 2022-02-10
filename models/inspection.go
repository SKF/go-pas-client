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

func (inspection *Inspection) FromInternal(internal *models.ModelsInspection) {
	if inspection == nil || internal == nil {
		return
	}

	inspection.Choices = make([]InspectionChoice, len(internal.Choices))

	for i, choice := range internal.Choices {
		inspection.Choices[i].FromInternal(choice)
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
