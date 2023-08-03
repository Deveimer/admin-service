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

type Doctor interface {
	Create(ctx *goofy.Context, doctor *models.DoctorRequest) (interface{}, error)
	Get(ctx *goofy.Context, id string) (*models.DoctorDetails, error)
	Update(ctx *goofy.Context, doctorDetails *models.DoctorRequest, id string) (*models.DoctorDetails, error)
	Delete(ctx *goofy.Context, id string) error
}

type DoctorOPDScheduler interface {
	Create(ctx *goofy.Context, request *models.OPDScheduleCreateRequest) (*models.OPDSchedule, error)
	GetAll(ctx *goofy.Context, filter *filters.DoctorOPDSchedule) ([]*models.OPDSchedule, error)
	GetById(ctx *goofy.Context, id string) (*models.OPDSchedule, error)
	Update(ctx *goofy.Context, id string, status string, reason string) (*models.OPDSchedule, error)
}

type AppointmentScheduler interface {
	Create(ctx *goofy.Context, request *models.AppointmentCreateRequest) (*models.Appointment, error)
	GetAll(ctx *goofy.Context, filter *filters.Appointment) ([]*models.Appointment, error)
	GetById(ctx *goofy.Context, id string) (*models.Appointment, error)
	Update(ctx *goofy.Context, request *models.AppointmentUpdateRequest) (*models.Appointment, error)
}
