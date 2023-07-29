package doctorOPDSchedule

import (
	"github.com/Deveimer/goofy/pkg/goofy"
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"main/internal/filters"
	"main/internal/models"
	"main/internal/services"
	"net/http"
	"path"
)

type Handler struct {
	DoctorOPDService services.DoctorOPDScheduler
}

func New(DoctorOPDService services.DoctorOPDScheduler) *Handler {
	return &Handler{DoctorOPDService: DoctorOPDService}
}

func (h *Handler) Create(ctx *goofy.Context) (interface{}, error) {
	var doctorOPDCreateRequest models.DoctorOPDScheduleCreateRequest

	err := ctx.Bind(&doctorOPDCreateRequest)
	if err != nil {
		ctx.Logger.Errorf("error while unmarshalling json in Create [error message]: %v ", err.Error())
		return nil, &errors.Response{StatusCode: http.StatusInternalServerError, Reason: "Unable to unmarshall user data"}
	}

	doctorOPDSchedule, err := h.DoctorOPDService.Create(ctx, &doctorOPDCreateRequest)
	if err != nil {
		return nil, err
	}

	return doctorOPDSchedule, nil
}

func (h *Handler) Index(ctx *goofy.Context) (interface{}, error) {
	queryParam := ctx.Request().URL.Query()

	var doctorOPDScheduleFilter filters.DoctorOPDSchedule
	doctorOPDScheduleFilter.DoctorID = queryParam.Get("doctorId")
	doctorOPDScheduleFilter.StartDate = queryParam.Get("startDate")
	doctorOPDScheduleFilter.EndDate = queryParam.Get("endDate")
	doctorOPDScheduleFilter.Status = queryParam.Get("status")
	doctorOPDScheduleFilter.Date = queryParam.Get("date")

	doctorOPDSchedules, err := h.DoctorOPDService.GetAll(ctx, &doctorOPDScheduleFilter)
	if err != nil {
		return nil, err
	}

	return doctorOPDSchedules, nil
}

func (h *Handler) Read(ctx *goofy.Context) (interface{}, error) {
	id := path.Base(ctx.Request().URL.Path)

	doctorOPDSchedule, err := h.DoctorOPDService.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return doctorOPDSchedule, err
}

func (h *Handler) Update(ctx *goofy.Context) (interface{}, error) {
	id := path.Base(ctx.Request().URL.Path)

	doctorOPDScheduleUpdateRequest := struct {
		Status string `json:"status"`
	}{}
	err := ctx.Bind(&doctorOPDScheduleUpdateRequest)
	if err != nil {
		ctx.Logger.Errorf("error while unmarshalling json in orderPayment create [error message]: %v ", err.Error())
		return nil, &errors.Response{StatusCode: http.StatusInternalServerError, Reason: "Unable to unmarshall user data"}
	}

	doctorOPDSchedule, err := h.DoctorOPDService.Update(ctx, id, doctorOPDScheduleUpdateRequest.Status)
	if err != nil {
		return nil, err
	}

	return doctorOPDSchedule, err
}
