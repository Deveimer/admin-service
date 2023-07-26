package patients

import (
	"database/sql"
	"github.com/Deveimer/goofy/pkg/goofy"
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"github.com/Deveimer/goofy/pkg/goofy/types"
	"main/internal/models"
	"time"
)

type PatientStore struct {
	DB *sql.DB
}

func New(db *sql.DB) *PatientStore {
	return &PatientStore{DB: db}
}

const tableName = "patient"

func (p *PatientStore) Create(ctx *goofy.Context, patient *models.PatientDetails) (interface{}, error) {
	query := `INSERT INTO ` + tableName + ` (id,name,gender,phone,email,age,city,state,pincode,joined_on,status,created_at,password,salt) 
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)`

	_, err := p.DB.Exec(query, patient.Id, patient.Name, patient.Gender, patient.Phone, patient.Email, patient.Age,
		patient.City, patient.State, patient.Pincode, time.Now(), patient.Status, time.Now(), patient.Password, patient.Salt)
	if err != nil {
		ctx.Logger.Errorf("error while creating new patient, Error: %v", err)
		return nil, errors.DbError{Err: err}
	}

	return types.Response{
		Data: "patient created successfully",
	}, nil
}

func (p *PatientStore) Get(ctx *goofy.Context, patient *models.PatientDetails) (*models.PatientDetails, error) {
	return nil, nil
}

func (p *PatientStore) Update(ctx *goofy.Context, patient *models.PatientDetails) (*models.PatientDetails, error) {
	return nil, nil
}

func (p *PatientStore) Delete(ctx *goofy.Context, patient *models.PatientDetails) error {
	return nil
}
