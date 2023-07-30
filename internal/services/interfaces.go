package services

import (
	"github.com/Deveimer/goofy/pkg/goofy"
	"main/internal/models"
	"mime/multipart"
)

type Patient interface {
	Create(ctx *goofy.Context, patient *models.PatientRequest) (interface{}, error)
	Get(ctx *goofy.Context, id string) (*models.PatientDetails, error)
	Update(ctx *goofy.Context, patientDetails *models.PatientRequest, id string) (*models.PatientDetails, error)
	Delete(ctx *goofy.Context, id string) error
}

type Prescription interface {
	Get(ctx *goofy.Context, patientID, doctorID string) (*models.PatientPrescription, error)
	Update(ctx *goofy.Context, req *models.PatientPrescription, patientID, doctorID string) (*models.PatientPrescription, error)
	Delete(ctx *goofy.Context, doctorID string) error
	Create(ctx *goofy.Context, patientID, doctorID, notes string, file multipart.File) (interface{}, error)
}
