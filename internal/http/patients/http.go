package patients

import (
	"github.com/Deveimer/goofy/pkg/goofy"
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"main/internal/models"
	"main/internal/services"
	"net/http"
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
	return nil, nil
}

func (h *PatientHandler) Update(ctx *goofy.Context) (interface{}, error) {
	return nil, nil
}

func (h *PatientHandler) Delete(ctx *goofy.Context) error {
	return nil
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

	if patient.City == "" && patient.State == "" && patient.Pincode < 1 {
		return errors.Response{
			StatusCode: http.StatusBadRequest,
			Status:     http.StatusText(http.StatusBadRequest),
			Reason:     "Invalid city/state/pincode",
		}
	}

	return nil
}
