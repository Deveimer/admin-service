package appointments

import (
	"database/sql"
	"fmt"
	"github.com/Deveimer/goofy/pkg/goofy"
	"main/internal/constants"
	"main/internal/filters"
	"main/internal/models"
	"strconv"
	"strings"
)

type appointmentStore struct {
	DB *sql.DB
}

func New(db *sql.DB) *appointmentStore {
	return &appointmentStore{DB: db}
}

func (a *appointmentStore) Create(ctx *goofy.Context, request *models.AppointmentCreateRequest) (*models.Appointment, error) {
	query := `INSERT INTO ` + constants.AppointmentTable +
		` (doctor_id, patient_id, status, description, start_time, end_time, appointment_date, created_by) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`

	var lastInsertedID int
	err := a.DB.QueryRow(
		query,
		request.DoctorID,
		request.PatientID,
		request.Status,
		request.Description,
		request.StartTime,
		request.EndTime,
		request.Date,
		request.CreatedBy,
	).Scan(&lastInsertedID)

	if err != nil {
		ctx.Logger.Errorf("INFO", fmt.Sprintf("Error While Inserting  to DB in appointments, Error: %v", err.Error()))
	}

	return a.GetByID(ctx, lastInsertedID)
}

func (a *appointmentStore) GetByID(ctx *goofy.Context, id int) (*models.Appointment, error) {
	query := `SELECT id, doctor_id, patient_id, status, description, start_time, end_time, appointment_date, reason FROM ` + constants.AppointmentTable + ` WHERE id = $1`

	var appointment models.Appointment

	err := a.DB.QueryRow(query, id).Scan(
		&appointment.ID,
		&appointment.DoctorID,
		&appointment.PatientID,
		&appointment.Status,
		&appointment.Description,
		&appointment.StartTime,
		&appointment.EndTime,
		&appointment.Date,
		&appointment.Reason,
	)
	if err != nil {
		ctx.Logger.Errorf("INFO", fmt.Sprintf("Error While fetching doctor opd schedule by id in appointments, Error: %v", err.Error()))
		return nil, err
	}

	return &appointment, nil
}

func (a *appointmentStore) GetAll(ctx *goofy.Context, filter *filters.Appointment) ([]*models.Appointment, error) {
	where, values := generateWhereClause(filter)

	if len(values) > 0 {
		where = ` WHERE ` + where
	}

	query := `SELECT id, doctor_id, patient_id,status, description, start_time, end_time, appointment_date, reason FROM ` + constants.AppointmentTable + where
	rows, err := a.DB.Query(query, values...)
	if err != nil {
		ctx.Logger.Errorf("INFO", fmt.Sprintf("Error While fetching doctor opd schedule getAll in appointment, Error: %v", err.Error()))
		return nil, err
	}

	var appointments []*models.Appointment
	for rows.Next() {
		var appointment models.Appointment

		err := rows.Scan(
			&appointment.ID,
			&appointment.DoctorID,
			&appointment.PatientID,
			&appointment.Status,
			&appointment.Description,
			&appointment.StartTime,
			&appointment.EndTime,
			&appointment.Date,
			&appointment.Reason,
		)
		if err != nil {
			ctx.Logger.Errorf("INFO", fmt.Sprintf("Error While Scaning appointment getAll in appointment, Error: %v", err.Error()))
			return nil, err
		}
		appointments = append(appointments, &appointment)
	}

	return appointments, err
}

func (a *appointmentStore) Update(ctx *goofy.Context, updateRequest *models.AppointmentUpdateRequest) (*models.Appointment, error) {
	setQuery, values, index := generateSetClause(updateRequest)

	if len(values) > 0 {
		setQuery = ` SET` + setQuery
	}

	values = append(values, updateRequest.ID)
	query := `UPDATE ` + constants.AppointmentTable + setQuery + ` WHERE id = $` + strconv.Itoa(index+1)

	_, err := a.DB.Exec(query, values...)
	if err != nil {
		ctx.Logger.Errorf("INFO", fmt.Sprintf("Error While updating status in appointmnet, Error: %v", err.Error()))
		return nil, err
	}

	return a.GetByID(ctx, updateRequest.ID)
}

func generateSetClause(appointmentUpdateRequest *models.AppointmentUpdateRequest) (setQuery string, values []interface{}, index int) {
	if appointmentUpdateRequest.Status != "" {
		setQuery += fmt.Sprintf(` status = $%v,`, strconv.Itoa(index+1))
		index += 1
		values = append(values, appointmentUpdateRequest.Status)
	}

	if appointmentUpdateRequest.Reason != "" && appointmentUpdateRequest.Status == "CANCELLED" {
		setQuery += fmt.Sprintf(` reason = $%v,`, strconv.Itoa(index+1))
		index += 1
		values = append(values, appointmentUpdateRequest.Reason)
	}

	if appointmentUpdateRequest.Description != "" {
		setQuery += fmt.Sprintf(` description = $%v,`, strconv.Itoa(index+1))
		index += 1
		values = append(values, appointmentUpdateRequest.Description)
	}

	setQuery = strings.TrimSuffix(setQuery, `,`)

	return
}

func generateWhereClause(filter *filters.Appointment) (where string, values []interface{}) {
	index := 0
	if filter.Status != "" {
		where += fmt.Sprintf(`status = $%v AND `, index+1)
		index += 1
		values = append(values, filter.Status)
	}

	if filter.DoctorID != "" {
		where += fmt.Sprintf(`doctor_id = $%v AND `, index+1)
		index += 1
		values = append(values, filter.DoctorID)
	}

	if filter.PatientID != "" {
		where += fmt.Sprintf(`patient_id = $%v AND `, index+1)
		index += 1
		values = append(values, filter.PatientID)
	}

	if filter.Date != "" {
		where += fmt.Sprintf(`appointment_date = CAST($%v AS DATE) AND `, index+1)
		index += 1
		values = append(values, filter.Date)
	}

	if filter.IncludeCancelled == false {
		where += fmt.Sprintf(`status != 'CANCELLED'`)
	}

	where = strings.TrimSuffix(where, `AND `)
	return
}
