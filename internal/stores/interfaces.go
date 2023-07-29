package stores

import (
	"github.com/Deveimer/goofy/pkg/goofy"

	"main/internal/models"
)

type Patient interface {
	Create(ctx *goofy.Context, patient *models.PatientDetails) (interface{}, error)
	Get(ctx *goofy.Context, id string) (*models.PatientDetails, error)
	GetPatientByPhoneAndEmail(ctx *goofy.Context, phone, email string) (string, error)
	Update(ctx *goofy.Context, patientDetails *models.PatientRequest, id string) (*models.PatientDetails, error)
	Delete(ctx *goofy.Context, id string) error
}
