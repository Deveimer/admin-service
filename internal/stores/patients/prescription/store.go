package prescription

import (
	"database/sql"
	"fmt"
	"github.com/Deveimer/goofy/pkg/goofy"
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"main/internal/filters"
	"main/internal/models"
	"net/http"
	"strconv"
	"time"
)

type Prescription struct {
	DB *sql.DB
}

func New(db *sql.DB) *Prescription {
	return &Prescription{DB: db}
}

const tableName = "patient_prescription"

func (s *Prescription) GetByID(ctx *goofy.Context, prescriptionID string) (*models.PatientPrescription, error) {
	query := `SELECT * FROM ` + tableName + ` WHERE  id = $1 and deleted_at is NULL`

	var prescriptionData models.PatientPrescription

	row := s.DB.QueryRow(query, prescriptionID)

	if err := row.Scan(&prescriptionData.ID, &prescriptionData.PatientID, &prescriptionData.DoctorID, &prescriptionData.PrescriptionLocation, &prescriptionData.Notes, &prescriptionData.CreatedAt, &prescriptionData.UpdatedAt, &prescriptionData.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Response{
				StatusCode: http.StatusNotFound,
				Status:     http.StatusText(http.StatusNotFound),
				Reason:     "Prescription Not Found",
			}
		}

		ctx.Log("INFO", fmt.Sprintf("Scan Error while reteriving prescription, Error: %v", err.Error()))

		return nil, errors.Response{
			StatusCode: http.StatusInternalServerError,
			Code:       http.StatusText(http.StatusInternalServerError),
			Reason:     "Sorry an unexpected error occurred, please try again later",
		}
	}

	return &prescriptionData, nil
}

func (s *Prescription) Get(ctx *goofy.Context, filter *filters.Prescription) (*models.PatientPrescription, error) {
	query := `SELECT * FROM ` + tableName + ` WHERE `

	q, val := generateWhereClause(filter, 0)
	query += q

	var prescriptionData models.PatientPrescription

	row := s.DB.QueryRow(query, val...)

	if err := row.Scan(&prescriptionData.ID, &prescriptionData.PatientID, &prescriptionData.DoctorID, &prescriptionData.PrescriptionLocation, &prescriptionData.Notes, &prescriptionData.CreatedAt, &prescriptionData.UpdatedAt, &prescriptionData.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Response{
				StatusCode: http.StatusNotFound,
				Status:     http.StatusText(http.StatusNotFound),
				Reason:     "Prescription Not Found",
			}
		}

		ctx.Log("INFO", fmt.Sprintf("Scan Error while reteriving prescription, Error: %v", err.Error()))

		return nil, errors.Response{
			StatusCode: http.StatusInternalServerError,
			Code:       http.StatusText(http.StatusInternalServerError),
			Reason:     "Sorry an unexpected error occurred, please try again later",
		}
	}

	return &prescriptionData, nil
}

func (s *Prescription) Create(ctx *goofy.Context, prescription *models.PatientPrescription) (int, error) {
	var prescriptionID int
	query := `INSERT INTO ` + tableName + ` (patient_id, doctor_id, prescription_location, notes) VALUES ($1,$2,$3,$4) RETURNING id`

	row := s.DB.QueryRow(query, prescription.PatientID, prescription.DoctorID, prescription.PrescriptionLocation, prescription.Notes)
	if err := row.Scan(&prescriptionID); err != nil {

		ctx.Log("INFO", fmt.Sprintf("Error While Inserting to DB, Error: %v", err))

		return 0, errors.Response{
			StatusCode: http.StatusInternalServerError,
			Code:       http.StatusText(http.StatusInternalServerError),
			Reason:     "Sorry an unexpected error occurred, please try again later",
		}
	}

	return prescriptionID, nil
}

func (s *Prescription) Update(ctx *goofy.Context, requestData *models.PatientPrescription, filter *filters.Prescription) error {
	var (
		val        []interface{}
		setQuery   string
		whereQuery string
	)

	setQuery, p := generateSetClause(requestData, 0)
	for i, _ := range p {
		val = append(val, p[i])
	}

	whereQuery, v := generateWhereClause(filter, len(p)+1)
	for i, _ := range v {
		val = append(val, v[i])
	}

	query := `UPDATE ` + tableName + ` SET ` + setQuery + ` WHERE ` + whereQuery

	_, err := s.DB.Exec(query, val...)
	if err != nil {
		ctx.Log("INFO", fmt.Sprintf("Error while updating prescription Error: %v", err))
		return errors.Response{
			StatusCode: http.StatusInternalServerError,
			Code:       http.StatusText(http.StatusInternalServerError),
			Reason:     "Sorry an unexpected error occurred, please try again later",
		}
	}

	return nil
}

func (s *Prescription) Delete(ctx *goofy.Context, doctorID string) error {
	query := `UPDATE` + tableName + ` SET deleted_at = $1 WHERE doctor_id = $2`

	_, err := s.DB.Exec(query, time.Now(), doctorID)
	if err != nil {
		ctx.Log("INFO", fmt.Sprintf("Error while deleting prescription for doctorID: %v, Error: %v", doctorID, err))
		return errors.Response{
			StatusCode: http.StatusInternalServerError,
			Code:       http.StatusText(http.StatusInternalServerError),
			Reason:     "Sorry an unexpected error occurred, please try again later",
		}
	}

	return nil
}

func generateSetClause(requestData *models.PatientPrescription, i int) (q string, p []interface{}) {

	var index = 1
	if i != 0 {
		index = i
	}

	if requestData.PrescriptionLocation != "" {
		q += "prescription_location = $" + strconv.Itoa(index) + "and "
		p = append(p, requestData.PrescriptionLocation)
		index += 1
	}

	if requestData.Notes != "" {
		q += "notes = $" + strconv.Itoa(index) + "and "
		p = append(p, requestData.Notes)
		index += 1
	}

	q += "updated_at = $" + strconv.Itoa(index)
	p = append(p, time.Now())
	index += 1

	return q, p
}

func generateWhereClause(filter *filters.Prescription, i int) (q string, p []interface{}) {

	var index = 1
	if i != 0 {
		index = i
	}

	if filter.PatientID != "" {
		q += "patient_id = $" + strconv.Itoa(index) + "and "
		p = append(p, filter.PatientID)
		index += 1
	}

	if filter.DoctorID != "" {
		q += "doctor_id = $" + strconv.Itoa(index) + "and "
		p = append(p, filter.DoctorID)
		index += 1
	}

	return q, p
}
