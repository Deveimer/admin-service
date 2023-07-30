package prescription

import (
	"github.com/Deveimer/goofy/pkg/goofy"
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"main/clients/ftp"
	"main/internal/filters"
	"main/internal/models"
	"main/internal/stores"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
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
	// todo validate doctor and patient
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

func (s *PrescriptionService) Delete(ctx *goofy.Context, doctorID string) error {
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

func (s *PrescriptionService) Create(ctx *goofy.Context, patientID, doctorID, notes string, file multipart.File) (interface{}, error) {
	fileName := "-Customers-" + time.Now().Format("202303131545") + ".pdf"
	path := "uploads/prescription/" + time.Now().Format("2006-01-02")
	location := path + "/" + fileName

	ftpObj, err := ftp.NewClient(ctx)
	if err != nil {
		ctx.Logger.Errorf("Unable to connect to the ftp client, Error: %v", err)
		return "", err
	}

	err = ftpObj.Append(location, file)
	if err != nil {
		return "", &errors.Response{
			StatusCode: http.StatusInternalServerError,
			Code:       http.StatusText(http.StatusInternalServerError),
			Reason:     "Could not upload the file",
		}
	}

	prescriptionDetails := models.PatientPrescription{
		PatientID:            patientID,
		DoctorID:             doctorID,
		PrescriptionLocation: location,
		Notes:                notes,
	}

	prescriptionId, err := s.store.Create(ctx, &prescriptionDetails)
	if err != nil {
		ctx.Logger.Errorf("Error while fetching from the DB, Error: %s", err.Error())
		return nil, &errors.Response{
			StatusCode: http.StatusInternalServerError,
			Code:       strconv.Itoa(http.StatusInternalServerError),
			Reason:     "Sorry, an unexpected error has occurred. Please try again later.",
		}
	}

	prescription, err := s.store.GetByID(ctx, strconv.Itoa(prescriptionId))
	if err != nil {
		return nil, &errors.Response{
			StatusCode: http.StatusInternalServerError,
			Code:       strconv.Itoa(http.StatusInternalServerError),
			Reason:     "Sorry, an unexpected error has occurred. Please try again later.",
		}
	}
	return prescription, err
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
