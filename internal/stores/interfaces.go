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
