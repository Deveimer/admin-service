package doctors

import (
	"net/http"

	"github.com/Deveimer/goofy/pkg/goofy"
	"github.com/Deveimer/goofy/pkg/goofy/errors"

	"main/internal/models"
	"main/internal/services"
)

type DoctorHandler struct {
	service services.Doctor
}

func New(service services.Doctor) *DoctorHandler {
	return &DoctorHandler{service: service}
}

func (h *DoctorHandler) Create(ctx *goofy.Context) (interface{}, error) {
	var doctor models.DoctorRequest

	err := ctx.Bind(&doctor)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"json request body"}}
	}

	err = validateDoctorRequest(&doctor)
	if err != nil {
		return nil, err
	}

	resp, err := h.service.Create(ctx, &doctor)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *DoctorHandler) Get(ctx *goofy.Context) (interface{}, error) {

	doctorId := ctx.PathParam("id")

	if doctorId == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	doctorDetails, err := h.service.Get(ctx, doctorId)
	if err != nil {
		return nil, err
	}

	return doctorDetails, nil
}

func (h *DoctorHandler) Update(ctx *goofy.Context) (interface{}, error) {
	var doctor models.DoctorRequest

	err := ctx.Bind(&doctor)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"json request body"}}
	}

	doctorId := ctx.PathParam("id")

	if doctorId == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	doctorDetails, err := h.service.Update(ctx, &doctor, doctorId)
	if err != nil {
		return nil, err
	}

	return doctorDetails, nil
}

func (h *DoctorHandler) Delete(ctx *goofy.Context) (interface{}, error) {
	doctorId := ctx.PathParam("id")

	if doctorId == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	err := h.service.Delete(ctx, doctorId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func validateDoctorRequest(doctor *models.DoctorRequest) error {
	if doctor.Name == "" {
		return errors.InvalidParam{Param: []string{"name"}}
	}

	if doctor.Phone == "" || len(doctor.Phone) != 10 {
		return errors.InvalidParam{Param: []string{"phone"}}
	}

	if doctor.Age < 1 {
		return errors.InvalidParam{Param: []string{"age"}}
	}

	if doctor.LicenceNumber == "" || len(doctor.LicenceNumber) > 9 {
		return errors.InvalidParam{Param: []string{"Licence Number"}}
	}

	if doctor.City == "" && doctor.State == "" && doctor.Pincode == "" {
		return errors.Response{
			StatusCode: http.StatusBadRequest,
			Status:     http.StatusText(http.StatusBadRequest),
			Reason:     "Invalid city/state/pincode",
		}
	}

	return nil
}
