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

type Doctor interface {
	Create(ctx *goofy.Context, doctor *models.DoctorRequest) (interface{}, error)
	Get(ctx *goofy.Context, id string) (*models.DoctorDetails, error)
	Update(ctx *goofy.Context, doctorDetails *models.DoctorRequest, id string) (*models.DoctorDetails, error)
	Delete(ctx *goofy.Context, id string) error
}
