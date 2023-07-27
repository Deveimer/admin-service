package prescription

import (
	"github.com/Deveimer/goofy/pkg/goofy"
	"main/internal/filters"
	"main/internal/models"
	"main/internal/stores"
)

type PrescriptionService struct {
	store stores.Prescription
}

func New(store stores.Prescription) *PrescriptionService {
	return &PrescriptionService{
		store: store,
	}
}

func (s *PrescriptionService) Get(ctx *goofy.Context, patientID, doctorID string) (*models.PatientPrescription, error) {
	// todo validate doctor and patient
	filter := validateAndGetFilter(patientID, doctorID)
	prescription, err := s.store.Get(ctx, filter)

	if err != nil {
		return nil, err
	}

	return prescription, nil
}

func (s *PrescriptionService) Update(ctx *goofy.Context, req *models.PatientPrescription, patientID, doctorID string) (*models.PatientPrescription, error) {
	prescription, err := s.Get(ctx, patientID, doctorID)
	if err != nil || prescription == nil {
		return nil, err
	}

	filter := validateAndGetFilter(patientID, doctorID)
	err = s.store.Update(ctx, req, filter)

	if err != nil {
		return nil, err
	}

	updatedPrescription, err := s.store.GetByID(ctx, prescription.ID)
	if err != nil || updatedPrescription == nil {
		return nil, err
	}

	return updatedPrescription, nil
}

func (s *PrescriptionService) Delete(ctx *goofy.Context, req *models.PatientPrescription, doctorID string) error {
	prescription, err := s.Get(ctx, "", doctorID)
	if err != nil || prescription == nil {
		return err
	}

	err = s.store.Delete(ctx, doctorID)

	if err != nil {
		return err
	}

	return nil
}

func validateAndGetFilter(patientID, doctorID string) *filters.Prescription {
	filter := filters.Prescription{}
	if patientID != "" {
		filter.PatientID = patientID
	}

	if doctorID != "" {
		filter.DoctorID = doctorID
	}

	return &filter
}
