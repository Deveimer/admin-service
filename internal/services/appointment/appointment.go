package appointment

import (
	"database/sql"
	"github.com/Deveimer/goofy/pkg/goofy"
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"main/internal/filters"
	"main/internal/models"
	"main/internal/services"
	"main/internal/stores"
	"net/http"
	"strconv"
	"time"
)

type appointmentScheduler struct {
	appointmentStore         stores.Appointment
	doctorOPDScheduleService services.DoctorOPDScheduler
}

func New(appointmentStore stores.Appointment, doctorOPDScheduleServices services.DoctorOPDScheduler) services.AppointmentScheduler {
	return &appointmentScheduler{appointmentStore: appointmentStore, doctorOPDScheduleService: doctorOPDScheduleServices}
}

func (a *appointmentScheduler) Create(ctx *goofy.Context, request *models.AppointmentCreateRequest) (*models.Appointment, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	doctorOPDSchedules, err := a.doctorOPDScheduleService.GetAll(ctx, &filters.DoctorOPDSchedule{Date: request.Date})
	if err != nil {
		ctx.Logger.Errorf("Error while calling doctorOPDSchedule service get All Error %v", err.Error())
		return nil, err
	}

	appointmentStartTime, _ := time.Parse("15:04:05", request.StartTime)
	appointmentEndTime, _ := time.Parse("15:04:05", request.EndTime)

	var isSlotBooked = false
	for _, doctorOPDSchedule := range doctorOPDSchedules {
		opdStartTime, _ := time.Parse("15:04:05", doctorOPDSchedule.OPDStartTime)
		opdEndTime, _ := time.Parse("15:04:05", doctorOPDSchedule.OPDEndTime)

		if ((appointmentStartTime.After(opdStartTime) || appointmentStartTime.Equal(opdStartTime)) &&
			(appointmentStartTime.Before(opdEndTime))) || (appointmentEndTime.After(opdStartTime) &&
			appointmentEndTime.Before(opdEndTime)) {
			isSlotBooked = true
			break
		}
	}

	appointments, err := a.GetAll(ctx, &filters.Appointment{Date: request.Date})
	if err != nil {
		return nil, err
	}

	for _, appointment := range appointments {
		startTime, _ := time.Parse("15:04:05", appointment.StartTime)
		endTime, _ := time.Parse("15:04:05", appointment.EndTime)

		if ((appointmentStartTime.After(startTime) || appointmentStartTime.Equal(startTime)) &&
			(appointmentStartTime.Before(endTime))) || (appointmentEndTime.After(startTime) &&
			appointmentEndTime.Before(endTime)) {
			isSlotBooked = true
			break
		}
	}

	if isSlotBooked {
		return nil, errors.Response{Reason: "slot already booked for this time and date", StatusCode: http.StatusBadRequest}
	}

	appointment, err := a.appointmentStore.Create(ctx, request)
	if err != nil {
		return nil, err
	}

	a.changeDateAndTimeToRequiredFormat(ctx, appointment)

	return appointment, nil
}

func (a *appointmentScheduler) GetAll(ctx *goofy.Context, filter *filters.Appointment) ([]*models.Appointment, error) {
	err := filter.Validate()
	if err != nil {
		return nil, err
	}

	appointments, err := a.appointmentStore.GetAll(ctx, filter)
	if err != nil {
		ctx.Logger.Errorf("Error for appointment store getAll  Error:%v", err.Error())
		return nil, err
	}

	for _, appointment := range appointments {
		a.changeDateAndTimeToRequiredFormat(ctx, appointment)
	}

	return appointments, nil
}

func (a *appointmentScheduler) GetById(ctx *goofy.Context, id string) (*models.Appointment, error) {
	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	appointmentID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	appointment, err := a.appointmentStore.GetByID(ctx, appointmentID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.EntityNotFound{Entity: "appointment", ID: id}
		}
		ctx.Logger.Errorf("error while calling getByID appointment store, ERROR %v", err.Error())
		return nil, err
	}

	a.changeDateAndTimeToRequiredFormat(ctx, appointment)

	return appointment, nil
}

func (a *appointmentScheduler) Update(ctx *goofy.Context, request *models.AppointmentUpdateRequest) (*models.Appointment, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	_, err = a.GetById(ctx, strconv.Itoa(request.ID))
	if err != nil {
		return nil, err
	}

	appointment, err := a.appointmentStore.Update(ctx, request)
	if err != nil {
		ctx.Logger.Errorf("error while calling update appointment store, ERROR %v", err.Error())
		return nil, err
	}

	a.changeDateAndTimeToRequiredFormat(ctx, appointment)

	return appointment, nil
}

func (dos *appointmentScheduler) changeDateAndTimeToRequiredFormat(ctx *goofy.Context, appointment *models.Appointment) {
	date, err := time.Parse("2006-01-02T15:04:05Z", appointment.Date)
	startTime, err := time.Parse("2006-01-02T15:04:05Z", appointment.StartTime)
	endTime, err := time.Parse("2006-01-02T15:04:05Z", appointment.EndTime)

	if err != nil {
		ctx.Logger.Errorf("Error in appointment changeDateAndTimeToRequiredFormat %v", err.Error())
	}

	appointment.Date = date.Format("2006-01-02")
	appointment.StartTime = startTime.Format("15:04:05")
	appointment.EndTime = endTime.Format("15:04:05")

	return
}
