package doctorOPDSchedule

import (
	"database/sql"
	"github.com/Deveimer/goofy/pkg/goofy"
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"main/internal/filters"
	"main/internal/models"
	"main/internal/services"
	"main/internal/stores"
	"main/internal/utils"
	"strconv"
	"time"
)

type doctorOPDScheduler struct {
	doctorOPDStore stores.DoctorOPDSchedule
}

func New(doctorOPDStore stores.DoctorOPDSchedule) services.DoctorOPDScheduler {
	return &doctorOPDScheduler{doctorOPDStore: doctorOPDStore}
}

func (dos *doctorOPDScheduler) Create(ctx *goofy.Context, request *models.OPDScheduleCreateRequest) (*models.OPDSchedule, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	request.OPDStatus = "SCHEDULED"
	doctorOPDSchedule, err := dos.doctorOPDStore.Create(ctx, request)
	if err != nil {
		ctx.Logger.Errorf("Error while creating doctorOpd schedule from store : err %v", err)
		return nil, errors.InternalServerError{}
	}

	dos.changeDateAndTimeToRequiredFormat(ctx, doctorOPDSchedule)

	return doctorOPDSchedule, nil
}

func (dos *doctorOPDScheduler) GetById(ctx *goofy.Context, id string) (*models.OPDSchedule, error) {
	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	doctorOPDScheduleID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	doctorOPDSchedule, err := dos.doctorOPDStore.GetByID(ctx, doctorOPDScheduleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.EntityNotFound{Entity: "doctor_opd_schedule", ID: id}
		}
		ctx.Logger.Errorf("Error while get by ID call doctorOpd schedule from store : err %v", err)
		return nil, errors.InternalServerError{}
	}

	dos.changeDateAndTimeToRequiredFormat(ctx, doctorOPDSchedule)

	return doctorOPDSchedule, nil
}

func (dos *doctorOPDScheduler) GetAll(ctx *goofy.Context, filter *filters.DoctorOPDSchedule) ([]*models.OPDSchedule, error) {
	err := filter.Validate()
	if err != nil {
		return nil, err
	}

	doctorOPDSchedules, err := dos.doctorOPDStore.GetAll(ctx, filter)
	if err != nil {
		ctx.Logger.Errorf("Error while get All call doctorOpd schedule from store : err %v", err)
		return nil, errors.InternalServerError{}
	}

	for _, doctorOPDSchedule := range doctorOPDSchedules {
		dos.changeDateAndTimeToRequiredFormat(ctx, doctorOPDSchedule)
	}

	return doctorOPDSchedules, nil
}

func (dos *doctorOPDScheduler) Update(ctx *goofy.Context, id string, status string, reason string) (*models.OPDSchedule, error) {
	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	doctorOPDScheduleID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	if status == "" {
		return nil, errors.MissingParam{Param: []string{"status"}}
	}

	if status != "" {
		if ok := utils.Validate(status, utils.Enum{Values: []interface{}{"SCHEDULED", "OPEN", "CANCELLED"}}); !ok {
			return nil, errors.InvalidParam{Param: []string{"status"}}
		}
	}

	if status == "CANCELLED" && reason == "" {
		return nil, errors.MissingParam{Param: []string{"reason"}}
	}

	_, err = dos.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	doctorOPDSchedule, err := dos.doctorOPDStore.Update(ctx, doctorOPDScheduleID, status, reason)
	if err != nil {
		ctx.Logger.Errorf("Error while update call doctorOpd schedule from store : err %v", err)
		return nil, errors.InternalServerError{}
	}

	dos.changeDateAndTimeToRequiredFormat(ctx, doctorOPDSchedule)

	return doctorOPDSchedule, nil
}

// ChangeDateAndTimeToRequiredFormat required format yyyy-mm-dd and HH-mm-ss
func (dos *doctorOPDScheduler) changeDateAndTimeToRequiredFormat(ctx *goofy.Context, doctorOPDSchedule *models.OPDSchedule) {
	startDate, err := time.Parse("2006-01-02T15:04:05Z", doctorOPDSchedule.OPDStartDate)
	endDate, err := time.Parse("2006-01-02T15:04:05Z", doctorOPDSchedule.OPDEndDate)
	startTime, err := time.Parse("2006-01-02T15:04:05Z", doctorOPDSchedule.OPDStartTime)
	endTime, err := time.Parse("2006-01-02T15:04:05Z", doctorOPDSchedule.OPDEndTime)

	if err != nil {
		ctx.Logger.Errorf("Error in doctorOPDService changeDateAndTimeToRequiredFormat %v", err.Error())
	}

	doctorOPDSchedule.OPDStartDate = startDate.Format("2006-01-02")
	doctorOPDSchedule.OPDEndDate = endDate.Format("2006-01-02")
	doctorOPDSchedule.OPDStartTime = startTime.Format("15:04:05")
	doctorOPDSchedule.OPDEndTime = endTime.Format("15:04:05")

	return
}
