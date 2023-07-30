package services

import (
	"github.com/Deveimer/goofy/pkg/goofy"

	"main/internal/filters"
	"main/internal/models"
)

type Patient interface {
	Create(ctx *goofy.Context, patient *models.PatientRequest) (interface{}, error)
	Get(ctx *goofy.Context, id string) (*models.PatientDetails, error)
	Update(ctx *goofy.Context, patientDetails *models.PatientRequest, id string) (*models.PatientDetails, error)
	Delete(ctx *goofy.Context, id string) error
}

type DoctorOPDScheduler interface {
	Create(ctx *goofy.Context, request *models.DoctorOPDScheduleCreateRequest) (*models.DoctorOPDSchedule, error)
	GetAll(ctx *goofy.Context, filter *filters.DoctorOPDSchedule) ([]*models.DoctorOPDSchedule, error)
	GetById(ctx *goofy.Context, id string) (*models.DoctorOPDSchedule, error)
	Update(ctx *goofy.Context, id string, status string, reason string) (*models.DoctorOPDSchedule, error)
}
