package doctorOPDSchedule

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

type doctorOPDScheduleStore struct {
	DB *sql.DB
}

func New(db *sql.DB) *doctorOPDScheduleStore {
	return &doctorOPDScheduleStore{DB: db}
}

func (dos *doctorOPDScheduleStore) Create(ctx *goofy.Context, request *models.DoctorOPDScheduleCreateRequest) (*models.DoctorOPDSchedule, error) {
	query := `INSERT INTO ` + constants.DoctorOPDScheduleTable +
		` (doctor_id, opd_status, opd_start_date, opd_end_date, opd_start_time, opd_end_time) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`

	var lastInsertedID int
	err := dos.DB.QueryRow(
		query,
		request.DoctorID,
		request.OPDStatus,
		request.OPDStartDate,
		request.OPDEndDate,
		request.OPDStartTime,
		request.OPDEndTime,
	).Scan(&lastInsertedID)

	if err != nil {
		ctx.Logger.Errorf("INFO", fmt.Sprintf("Error While Inserting  to DB in doctorOPDScheduleStore, Error: %v", err.Error()))
	}

	return dos.GetByID(ctx, lastInsertedID)
}

func (dos *doctorOPDScheduleStore) GetByID(ctx *goofy.Context, id int) (*models.DoctorOPDSchedule, error) {
	var doctorOPDSchedule models.DoctorOPDSchedule

	query := `SELECT * FROM ` + constants.DoctorOPDScheduleTable + ` WHERE id = $1`

	err := dos.DB.QueryRow(query, id).Scan(
		&doctorOPDSchedule.ID,
		&doctorOPDSchedule.DoctorID,
		&doctorOPDSchedule.OPDStatus,
		&doctorOPDSchedule.OPDStartDate,
		&doctorOPDSchedule.OPDEndDate,
		&doctorOPDSchedule.OPDStartTime,
		&doctorOPDSchedule.OPDEndTime,
		&doctorOPDSchedule.OPDCancelReason,
	)
	if err != nil {
		ctx.Logger.Errorf("INFO", fmt.Sprintf("Error While fetching doctor opd schedule by id in doctorOPDScheduleStore, Error: %v", err.Error()))
		return nil, err
	}

	return &doctorOPDSchedule, nil
}

func (dos *doctorOPDScheduleStore) GetAll(ctx *goofy.Context, filter *filters.DoctorOPDSchedule) ([]*models.DoctorOPDSchedule, error) {
	where, values := generateWhereClause(filter)

	if len(values) > 0 {
		where = ` WHERE ` + where
	}

	query := `SELECT * FROM ` + constants.DoctorOPDScheduleTable + where

	rows, err := dos.DB.Query(query, values...)
	if err != nil {
		ctx.Logger.Errorf("INFO", fmt.Sprintf("Error While fetching doctor opd schedule getAll in doctorOPDScheduleStore, Error: %v", err.Error()))
		return nil, err
	}

	var doctorOPDSchedules []*models.DoctorOPDSchedule

	for rows.Next() {
		var doctorOPDSchedule models.DoctorOPDSchedule
		err := rows.Scan(&doctorOPDSchedule.ID,
			&doctorOPDSchedule.DoctorID,
			&doctorOPDSchedule.OPDStatus,
			&doctorOPDSchedule.OPDStartDate,
			&doctorOPDSchedule.OPDEndDate,
			&doctorOPDSchedule.OPDStartTime,
			&doctorOPDSchedule.OPDEndTime,
			&doctorOPDSchedule.OPDCancelReason,
		)
		if err != nil {
			ctx.Logger.Errorf("INFO", fmt.Sprintf("Error While Scaning doctor opd schedule getAll in doctorOPDScheduleStore, Error: %v", err.Error()))
			return nil, err
		}
		doctorOPDSchedules = append(doctorOPDSchedules, &doctorOPDSchedule)
	}

	return doctorOPDSchedules, nil
}

func (dos *doctorOPDScheduleStore) Update(ctx *goofy.Context, id int, status string, reason string) (*models.DoctorOPDSchedule, error) {
	setQuery, values, index := generateSetClause(status, reason)

	if len(values) > 0 {
		setQuery = ` SET` + setQuery
	}

	values = append(values, id)
	query := `UPDATE ` + constants.DoctorOPDScheduleTable + setQuery + ` WHERE id = $` + strconv.Itoa(index+1)

	_, err := dos.DB.Exec(query, values...)
	if err != nil {
		ctx.Logger.Errorf("INFO", fmt.Sprintf("Error While updating status in doctorOPDScheduleStore, Error: %v", err.Error()))
		return nil, err
	}

	return dos.GetByID(ctx, id)
}

func generateSetClause(status string, reason string) (setQuery string, values []interface{}, index int) {
	if status != "" {
		setQuery += fmt.Sprintf(` opd_status = $%v,`, strconv.Itoa(index+1))
		index += 1
		values = append(values, status)
	}

	if reason != "" && status == "CANCELLED" {
		setQuery += fmt.Sprintf(` opd_cancel_reason = $%v,`, strconv.Itoa(index+1))
		index += 1
		values = append(values, reason)
	}

	setQuery = strings.TrimSuffix(setQuery, `,`)

	return
}

func generateWhereClause(filter *filters.DoctorOPDSchedule) (where string, values []interface{}) {
	index := 0

	if filter.DoctorID != "" {
		index += 1
		where += `doctor_id = $` + strconv.Itoa(index) + ` AND `
		values = append(values, filter.DoctorID)
	}

	if filter.StartDate != "" && filter.EndDate != "" {
		//where += `(opd_start_date, opd_end_date) OVERLAPS (CAST($` + strconv.Itoa(index+1) + ` AS DATE), CAST($` + strconv.Itoa(index+2) + ` AS DATE)) AND `
		where += `opd_start_date  >= CAST($` + strconv.Itoa(index+1) + ` AS DATE) AND CAST($` + strconv.Itoa(index+2) + ` AS DATE) <= opd_end_date AND ` + `opd_start_date  <= CAST($` + strconv.Itoa(index+3) + ` AS DATE) AND CAST($` + strconv.Itoa(index+4) + ` AS DATE) >= opd_end_date AND `
		index += 4
		values = append(values, filter.StartDate, filter.StartDate, filter.EndDate, filter.EndDate)
	}

	if filter.Status != "" {
		index += 1
		where += `opd_status = $` + strconv.Itoa(index) + ` AND `
		values = append(values, filter.Status)
	}

	if filter.Date != "" {
		where += `$` + strconv.Itoa(index+1) + ` >= opd_start_date AND $` + strconv.Itoa(index+2) + ` >= opd_end_date AND `
		index += 2
		values = append(values, filter.Date, filter.Date)
	}

	where = strings.TrimSuffix(where, `AND `)

	return
}
