package doctors

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/Deveimer/goofy/pkg/goofy"
	"github.com/Deveimer/goofy/pkg/goofy/errors"

	"main/internal/models"
)

type DoctorStore struct {
	DB *sql.DB
}

func New(db *sql.DB) *DoctorStore {
	return &DoctorStore{DB: db}
}

const tableName = "doctor"

func (p *DoctorStore) Create(ctx *goofy.Context, patient *models.DoctorDetails) (interface{}, error) {
	var id string

	query := `INSERT INTO ` + tableName + ` (id,name,gender,phone,email,licence_number,age,city,state,pincode,joined_on,status,created_at,updated_at,password,salt) 
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16) RETURNING id`

	err := p.DB.QueryRow(query, patient.Id, patient.Name, patient.Gender, patient.Phone, patient.Email, patient.LicenceNumber, patient.Age,
		patient.City, patient.State, patient.Pincode, time.Now().Format("2006-01-02 15:04:05"), patient.Status, time.Now(), time.Now(), patient.Password, patient.Salt).Scan(&id)

	if err != nil {
		ctx.Logger.Errorf("error while creating new patient, Error: %v", err)
		return nil, errors.DbError{Err: err}
	}

	return p.Get(ctx, id)
}

func (p *DoctorStore) Get(ctx *goofy.Context, id string) (*models.DoctorDetails, error) {

	var doctor models.DoctorDetails

	query := `SELECT id,name,gender,phone,email,licence_number,age,city,state,pincode,joined_on,status,created_at,updated_at FROM `
	query += tableName + ` WHERE id = $1`

	err := p.DB.QueryRow(query, id).Scan(&doctor.Id, &doctor.Name, &doctor.Gender, &doctor.Phone, &doctor.Email, &doctor.LicenceNumber, &doctor.Age,
		&doctor.City, &doctor.State, &doctor.Pincode, &doctor.JoinedOn, &doctor.Status, &doctor.CreatedAt, &doctor.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.EntityNotFound{
			Entity: "doctor",
			ID:     id,
		}
	}

	if err != nil {
		ctx.Logger.Errorf("error while fetching doctor from store, Error: %v", err)
		return nil, errors.DbError{Err: err}
	}

	return &doctor, nil
}

func (p *DoctorStore) CheckDoctorExist(ctx *goofy.Context, phone, email, licenceNumber string) (string, error) {

	var id string

	query := `SELECT id FROM ` + tableName + ` WHERE phone = $1 OR email = $2 OR licence_number = $3`

	err := p.DB.QueryRow(query, phone, email, licenceNumber).Scan(&id)
	if err == sql.ErrNoRows {
		return "", errors.EntityNotFound{
			Entity: "patient",
			ID:     id,
		}
	}

	if err != nil {
		ctx.Logger.Errorf("error while fetching doctor from store using email/phone, Error: %v", err)
		return "", errors.DbError{Err: err}
	}

	return id, nil
}

func (p *DoctorStore) Update(ctx *goofy.Context, doctorDetails *models.DoctorRequest, id string) (*models.DoctorDetails, error) {

	query, values, i := getUpdateQuery(doctorDetails)
	query = `UPDATE ` + tableName + ` SET ` + query + ` WHERE id = $` + strconv.Itoa(i+1)

	values = append(values, id)

	_, err := p.DB.Exec(query, values...)

	if err != nil {
		ctx.Logger.Errorf("error while updating doctor from store, Error: %v", err)
		return nil, errors.DbError{Err: err}
	}

	return p.Get(ctx, id)
}

func (p *DoctorStore) Delete(ctx *goofy.Context, id string) error {
	query := `DELETE FROM ` + tableName + ` WHERE id = $1`

	_, err := p.DB.Exec(query, id)
	if err != nil {
		ctx.Logger.Errorf("error while deleting doctor from store, Error: %v", err)
		return errors.DbError{Err: err}
	}

	return nil
}

func getUpdateQuery(doctorDetails *models.DoctorRequest) (query string, values []interface{}, i int) {

	i = 1

	if doctorDetails.Name != "" {
		query += "name = $" + strconv.Itoa(i)
		values = append(values, doctorDetails.Name)
		i += 1
	}

	if doctorDetails.Gender != "" {
		query += ",gender = $" + strconv.Itoa(i)
		values = append(values, doctorDetails.Gender)
		i += 1
	}

	if doctorDetails.Phone != "" {
		query += ",phone = $" + strconv.Itoa(i)
		values = append(values, doctorDetails.Phone)
		i += 1
	}

	if doctorDetails.Email != "" {
		query += ",email = $" + strconv.Itoa(i)
		values = append(values, doctorDetails.Email)
		i += 1
	}

	if doctorDetails.Age > 0 {
		query += ",age = $" + strconv.Itoa(i)
		values = append(values, doctorDetails.Age)
		i += 1
	}

	if doctorDetails.City != "" {
		query += ",city = $" + strconv.Itoa(i)
		values = append(values, doctorDetails.City)
		i += 1
	}

	if doctorDetails.State != "" {
		query += ",state = $" + strconv.Itoa(i)
		values = append(values, doctorDetails.State)
		i += 1
	}

	if doctorDetails.Status != "" {
		query += ",status = $" + strconv.Itoa(i)
		values = append(values, doctorDetails.Status)
		i += 1
	}

	if doctorDetails.Pincode != "" {
		query += ",pincode = $" + strconv.Itoa(i)
		values = append(values, doctorDetails.Pincode)
		i += 1
	}

	if doctorDetails.LicenceNumber != "" {
		query += ",licence_number = $" + strconv.Itoa(i)
		values = append(values, doctorDetails.LicenceNumber)
		i += 1
	}

	query += ",updated_at = $" + strconv.Itoa(i)
	values = append(values, time.Now())

	query = strings.TrimPrefix(query, ",")

	return
}
