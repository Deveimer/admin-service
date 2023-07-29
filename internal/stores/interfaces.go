package stores

import (
	"github.com/Deveimer/goofy/pkg/goofy"

	"main/internal/filters"
	"main/internal/models"
)

type Patient interface {
	Create(ctx *goofy.Context, patient *models.PatientDetails) (interface{}, error)
	Get(ctx *goofy.Context, id string) (*models.PatientDetails, error)
	GetPatientByPhoneAndEmail(ctx *goofy.Context, phone, email string) (string, error)
	Update(ctx *goofy.Context, patientDetails *models.PatientRequest, id string) (*models.PatientDetails, error)
	Delete(ctx *goofy.Context, id string) error
}

type DoctorOPDSchedule interface {
	Create(ctx *goofy.Context, request *models.DoctorOPDScheduleCreateRequest) (*models.DoctorOPDSchedule, error)
	GetByID(ctx *goofy.Context, id int) (*models.DoctorOPDSchedule, error)
	GetAll(ctx *goofy.Context, filter *filters.DoctorOPDSchedule) ([]*models.DoctorOPDSchedule, error)
	Update(ctx *goofy.Context, status string, id int) (*models.DoctorOPDSchedule, error)
}

