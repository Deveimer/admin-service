package patients

import (
	"net/http"

	"github.com/Deveimer/goofy/pkg/goofy"
	"github.com/Deveimer/goofy/pkg/goofy/errors"

	"main/internal/models"
	"main/internal/services"
)

type PatientHandler struct {
	service services.Patient
}

func New(service services.Patient) *PatientHandler {
	return &PatientHandler{service}
}

func (h *PatientHandler) Create(ctx *goofy.Context) (interface{}, error) {
	var patient models.PatientRequest

	err := ctx.Bind(&patient)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"json request body"}}
	}

	err = validatePatientRequest(&patient)
	if err != nil {
		return nil, err
	}

	resp, err := h.service.Create(ctx, &patient)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *PatientHandler) Get(ctx *goofy.Context) (interface{}, error) {

	patientId := ctx.PathParam("id")

	if patientId == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	patientDetails, err := h.service.Get(ctx, patientId)
	if err != nil {
		return nil, err
	}

	return patientDetails, nil
}

func (h *PatientHandler) Update(ctx *goofy.Context) (interface{}, error) {
	var patient models.PatientRequest

	err := ctx.Bind(&patient)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"json request body"}}
	}

	patientId := ctx.PathParam("id")

	if patientId == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	patientDetails, err := h.service.Update(ctx, &patient, patientId)
	if err != nil {
		return nil, err
	}

	return patientDetails, nil
}

func (h *PatientHandler) Delete(ctx *goofy.Context) (interface{}, error) {
	patientId := ctx.PathParam("id")

	if patientId == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	err := h.service.Delete(ctx, patientId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func validatePatientRequest(patient *models.PatientRequest) error {
	if patient.Name == "" {
		return errors.InvalidParam{Param: []string{"name"}}
	}

	if patient.Phone == "" || len(patient.Phone) != 10 {
		return errors.InvalidParam{Param: []string{"phone"}}
	}

	if patient.Age < 1 {
		return errors.InvalidParam{Param: []string{"age"}}
	}

	if patient.City == "" && patient.State == "" && patient.Pincode != "" {
		return errors.Response{
			StatusCode: http.StatusBadRequest,
			Status:     http.StatusText(http.StatusBadRequest),
			Reason:     "Invalid city/state/pincode",
		}
	}

	return nil
}
