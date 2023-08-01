package services

import (
	"github.com/Deveimer/goofy/pkg/goofy"

	"main/internal/models"
)

type Patient interface {
	Create(ctx *goofy.Context, patient *models.PatientRequest) (interface{}, error)
	Get(ctx *goofy.Context, id string) (*models.PatientDetails, error)
	Update(ctx *goofy.Context, patientDetails *models.PatientRequest, id string) (*models.PatientDetails, error)
	Delete(ctx *goofy.Context, id string) error
}
