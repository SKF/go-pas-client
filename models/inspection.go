package models

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	models "github.com/SKF/go-pas-client/internal/models"
	pas "github.com/SKF/proto/v2/pas"
)

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

func (i Inspection) ToInternal() *models.ModelsInspection {
	inspection := &models.ModelsInspection{
		Choices: make([]*models.ModelsInspectionChoice, len(i.Choices)),
	}

	for i, choice := range i.Choices {
		inspection.Choices[i] = choice.ToInternal()
	}

	return inspection
}

func (i InspectionChoice) ToInternal() *models.ModelsInspectionChoice {
	status := int32(i.Status)

	return &models.ModelsInspectionChoice{
		Answer:      i.Answer,
		Instruction: i.Instruction,
		Status:      &status,
	}
}

func (i *Inspection) FromProto(buf []byte) error {
	var internal pas.Inspection

	if err := proto.Unmarshal(buf, &internal); err != nil {
		return fmt.Errorf("decoding inspection alarm failed: %w", err)
	}

	i.Choices = make([]InspectionChoice, len(internal.Choices))

	for idx, choice := range internal.Choices {
		i.Choices[idx].FromProto(choice)
	}

	return nil
}

func (i *InspectionChoice) FromProto(internal *pas.InspectionChoice) {
	if i == nil || internal == nil {
		return
	}

	i.Answer = internal.Answer
	i.Instruction = internal.Instruction
	i.Status = AlarmStatusType(internal.Status)
}
