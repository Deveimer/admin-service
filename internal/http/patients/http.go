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

func New() *PatientHandler {
	return &PatientHandler{}
}
