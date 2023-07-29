package stores

import (
	"github.com/Deveimer/goofy/pkg/goofy"

	"main/internal/filters"
	"main/internal/models"
)

type Prescription interface {
	GetByID(ctx *goofy.Context, prescriptionID string) (*models.PatientPrescription, error)
	Get(ctx *goofy.Context, filter *filters.Prescription) (*models.PatientPrescription, error)
	Create(ctx *goofy.Context, prescription *models.PatientPrescription) (int, error)
	Update(ctx *goofy.Context, requestData *models.PatientPrescription, filter *filters.Prescription) error
	Delete(ctx *goofy.Context, doctorID string) error
}

type Patient interface {
	Create(ctx *goofy.Context, patient *models.PatientDetails) (interface{}, error)
	Get(ctx *goofy.Context, id string) (*models.PatientDetails, error)
	GetPatientByPhoneAndEmail(ctx *goofy.Context, phone, email string) (string, error)
	Update(ctx *goofy.Context, patientDetails *models.PatientRequest, id string) (*models.PatientDetails, error)
	Delete(ctx *goofy.Context, id string) error
}
