package patients

import (
	"database/sql"
	"time"

	"github.com/Deveimer/goofy/pkg/goofy"
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"main/internal/models"
)

type PatientStore struct {
	DB *sql.DB
}

func New(db *sql.DB) *PatientStore {
	return &PatientStore{DB: db}
}

const tableName = "patient"

func (p *PatientStore) Create(ctx *goofy.Context, patient *models.PatientDetails) (interface{}, error) {
	query := `INSERT INTO ` + tableName + ` (id,name,gender,phone,email,age,city,state,pincode,joined_on,status,created_at,updated_at,password,salt) 
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)`

	_, err := p.DB.Exec(query, patient.Id, patient.Name, patient.Gender, patient.Phone, patient.Email, patient.Age,
		patient.City, patient.State, patient.Pincode, time.Now(), patient.Status, time.Now(), time.Now(), patient.Password, patient.Salt)
	if err != nil {
		ctx.Logger.Errorf("error while creating new patient, Error: %v", err)
		return nil, errors.DbError{Err: err}
	}

	return nil, nil
}

func (p *PatientStore) Get(ctx *goofy.Context, id string) (*models.PatientDetails, error) {

	var patient models.PatientDetails

	query := `SELECT id,name,gender,phone,email,age,city,state,pincode,joined_on,status,created_at,updated_at FROM patient WHERE id = $1`

	err := p.DB.QueryRow(query, id).Scan(&patient.Id, &patient.Name, &patient.Gender, &patient.Phone, &patient.Email, &patient.Age,
		&patient.City, &patient.State, &patient.Pincode, &patient.JoinedOn, &patient.Status, &patient.CreatedAt, &patient.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.EntityNotFound{
			Entity: "patient",
			ID:     id,
		}
	}

	if err != nil {
		ctx.Logger.Errorf("error while fetching patient from store, Error: %v", err)
		return nil, errors.DbError{Err: err}
	}

	return &patient, nil
}

func (p *PatientStore) GetPatientByPhoneAndEmail(ctx *goofy.Context, phone, email string) (string, error) {

	var id string

	query := `SELECT id FROM patient WHERE phone = $1 OR email = $2`

	err := p.DB.QueryRow(query, phone, email).Scan(&id)
	if err != nil && err == sql.ErrNoRows {
		ctx.Logger.Errorf("error while fetching patient from store, Error: %v", err)
		return "", errors.DbError{Err: err}
	}

	return id, nil
}

func (p *PatientStore) Update(ctx *goofy.Context, patientDetails *models.PatientRequest, id string) (*models.PatientDetails, error) {

	query := `UPDATE ` + tableName + ` SET name = $1, gender = $2, phone = $3, email = $4, age = $5, city = $6, state = $7, pincode = $8, status = $9, updated_at = $10 WHERE id = $11`

	_, err := p.DB.Exec(query, patientDetails.Name, patientDetails.Gender, patientDetails.Phone, patientDetails.Email, patientDetails.Age,
		patientDetails.City, patientDetails.State, patientDetails.Pincode, patientDetails.Status, time.Now(), id)
	if err != nil {
		ctx.Logger.Errorf("error while updating patient from store, Error: %v", err)
		return nil, errors.DbError{Err: err}
	}

	return p.Get(ctx, id)
}

func (p *PatientStore) Delete(ctx *goofy.Context, id string) error {
	query := `DELETE FROM ` + tableName + ` WHERE id = $1`

	_, err := p.DB.Exec(query, id)
	if err != nil {
		ctx.Logger.Errorf("error while updating patient from store, Error: %v", err)
		return errors.DbError{Err: err}
	}

	return nil
}
