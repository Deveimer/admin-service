package patients

import (
	"database/sql"
	"strconv"
	"strings"
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
	var id string

	query := `INSERT INTO ` + tableName + ` (id,name,gender,phone,email,age,city,state,pincode,joined_on,status,created_at,updated_at,password,salt) 
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15) RETURNING id`

	err := p.DB.QueryRow(query, patient.Id, patient.Name, patient.Gender, patient.Phone, patient.Email, patient.Age,
		patient.City, patient.State, patient.Pincode, time.Now(), patient.Status, time.Now(), time.Now(), patient.Password, patient.Salt).Scan(&id)
	if err != nil {
		ctx.Logger.Errorf("error while creating new patient, Error: %v", err)
		return nil, errors.DbError{Err: err}
	}

	return p.Get(ctx, id)
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
	if err == sql.ErrNoRows {
		return "", errors.EntityNotFound{
			Entity: "patient",
			ID:     id,
		}
	}

	if err != nil {
		ctx.Logger.Errorf("error while fetching patient from store, Error: %v", err)
		return "", errors.DbError{Err: err}
	}

	return id, nil
}

func (p *PatientStore) Update(ctx *goofy.Context, patientDetails *models.PatientRequest, id string) (*models.PatientDetails, error) {

	query, values, i := getUpdateQuery(patientDetails)
	query = `UPDATE ` + tableName + ` SET ` + query + ` WHERE id = $` + strconv.Itoa(i+1)

	values = append(values, id)

	_, err := p.DB.Exec(query, values...)
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

func getUpdateQuery(patientDetails *models.PatientRequest) (query string, values []interface{}, i int) {

	i = 1

	if patientDetails.Name != "" {
		query += "name = $" + strconv.Itoa(i)
		values = append(values, patientDetails.Name)
		i += 1
	}

	if patientDetails.Gender != "" {
		query += ",gender = $" + strconv.Itoa(i)
		values = append(values, patientDetails.Gender)
		i += 1
	}

	if patientDetails.Phone != "" {
		query += ",phone = $" + strconv.Itoa(i)
		values = append(values, patientDetails.Phone)
		i += 1
	}

	if patientDetails.Email != "" {
		query += ",email = $" + strconv.Itoa(i)
		values = append(values, patientDetails.Email)
		i += 1
	}

	if patientDetails.Age > 0 {
		query += ",age = $" + strconv.Itoa(i)
		values = append(values, patientDetails.Age)
		i += 1
	}

	if patientDetails.City != "" {
		query += ",city = $" + strconv.Itoa(i)
		values = append(values, patientDetails.City)
		i += 1
	}

	if patientDetails.State != "" {
		query += ",state = $" + strconv.Itoa(i)
		values = append(values, patientDetails.State)
		i += 1
	}

	if patientDetails.Status != "" {
		query += ",status = $" + strconv.Itoa(i)
		values = append(values, patientDetails.Status)
		i += 1
	}

	if patientDetails.Pincode != "" {
		query += ",pincode = $" + strconv.Itoa(i)
		values = append(values, patientDetails.Pincode)
		i += 1
	}

	query += ",updated_at = $" + strconv.Itoa(i)
	values = append(values, time.Now())

	query = strings.TrimPrefix(query, ",")

	return
}
