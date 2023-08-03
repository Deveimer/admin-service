package appointment

import (
	"github.com/Deveimer/goofy/pkg/goofy"
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"main/internal/filters"
	"main/internal/models"
	"main/internal/services"
	"net/http"
	"path"
	"strconv"
)

type Handler struct {
	appointmentService services.AppointmentScheduler
}

func New(appointmentService services.AppointmentScheduler) *Handler {
	return &Handler{appointmentService: appointmentService}
}

func (h *Handler) Create(ctx *goofy.Context) (interface{}, error) {
	var appointmentCreateRequest models.AppointmentCreateRequest

	err := ctx.Bind(&appointmentCreateRequest)
	if err != nil {
		ctx.Logger.Errorf("error while unmarshalling json in Create [error message]: %v ", err.Error())
		return nil, &errors.Response{StatusCode: http.StatusInternalServerError, Reason: "Unable to unmarshall user data"}
	}

	appointment, err := h.appointmentService.Create(ctx, &appointmentCreateRequest)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}

func (h *Handler) Index(ctx *goofy.Context) (interface{}, error) {
	queryParam := ctx.Request().URL.Query()

	var appointmentFilter filters.Appointment
	appointmentFilter.DoctorID = queryParam.Get("doctorId")
	appointmentFilter.Status = queryParam.Get("status")
	appointmentFilter.Date = queryParam.Get("date")
	appointmentFilter.PatientID = queryParam.Get("patient")
	includeCancelled, _ := strconv.ParseBool(queryParam.Get("includeCancelled"))
	appointmentFilter.IncludeCancelled = includeCancelled

	doctorOPDSchedules, err := h.appointmentService.GetAll(ctx, &appointmentFilter)
	if err != nil {
		return nil, err
	}

	return doctorOPDSchedules, nil
}

func (h *Handler) Read(ctx *goofy.Context) (interface{}, error) {
	id := path.Base(ctx.Request().URL.Path)

	appointment, err := h.appointmentService.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return appointment, err
}

func (h *Handler) Update(ctx *goofy.Context) (interface{}, error) {
	id := path.Base(ctx.Request().URL.Path)

	var appointmentUpdateRequest *models.AppointmentUpdateRequest
	err := ctx.Bind(&appointmentUpdateRequest)
	if err != nil {
		ctx.Logger.Errorf("error while unmarshalling json in orderPayment create [error message]: %v ", err.Error())
		return nil, &errors.Response{StatusCode: http.StatusInternalServerError, Reason: "Unable to unmarshall user data"}
	}

	appointmentID, _ := strconv.Atoi(id)
	appointmentUpdateRequest.ID = appointmentID
	doctorOPDSchedule, err := h.appointmentService.Update(ctx, appointmentUpdateRequest)
	if err != nil {
		return nil, err
	}

	return doctorOPDSchedule, err
}
