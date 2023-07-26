package services

import (
	"github.com/Deveimer/goofy/pkg/goofy"
	"main/internal/models"
)

type Patient interface {
	Create(ctx *goofy.Context, patient *models.PatientRequest) (interface{}, error)
	Get(ctx *goofy.Context, patient *models.PatientRequest) (*models.PatientDetails, error)
	Update(ctx *goofy.Context, patient *models.PatientRequest) (*models.PatientDetails, error)
	Delete(ctx *goofy.Context, patient *models.PatientRequest) error
}
