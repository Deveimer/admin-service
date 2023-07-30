package prescription

import (
	"github.com/Deveimer/goofy/pkg/goofy"
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"github.com/Deveimer/goofy/pkg/goofy/types"
	"main/internal/models"
	"main/internal/services"
	"net/http"
)

type Handler struct {
	svc services.Prescription
}

func New(svc services.Prescription) *Handler {
	return &Handler{svc}
}

func (h *Handler) Create(ctx *goofy.Context) (interface{}, error) {
	file, _, err := ctx.Request().FormFile("fileUpload")
	if err != nil {
		return nil, &errors.Response{
			StatusCode: http.StatusBadRequest,
			Code:       http.StatusText(http.StatusBadRequest),
			Reason:     "Please provide the file data",
		}
	}

	patientID := ctx.Request().FormValue("patientId")
	doctorID := ctx.Request().FormValue("doctorId")
	notes := ctx.Request().FormValue("notes")

	resp, err := h.svc.Create(ctx, patientID, doctorID, notes, file)
	if err != nil {
		return nil, err
	}

	response := types.Response{
		Data: resp,
	}

	return response, nil
}

func (h *Handler) Get(ctx *goofy.Context) (interface{}, error) {
	patientID := ctx.PathParam("patientId")
	doctorID := ctx.PathParam("doctorId")

	if patientID == "" && doctorID == "" {
		return nil, errors.MissingParam{Param: []string{"patientId or doctorId"}}
	}

	patientDetails, err := h.svc.Get(ctx, patientID, doctorID)
	if err != nil {
		return nil, err
	}

	return patientDetails, nil
}

func (h *Handler) Update(ctx *goofy.Context) (interface{}, error) {
	var prescription models.PatientPrescription

	err := ctx.Bind(&prescription)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"json request body"}}
	}

	patientID := ctx.PathParam("patientId")
	doctorID := ctx.PathParam("doctorId")

	if patientID == "" && doctorID == "" {
		return nil, errors.MissingParam{Param: []string{"patientId or doctorId"}}
	}

	patientDetails, err := h.svc.Update(ctx, &prescription, patientID, doctorID)
	if err != nil {
		return nil, err
	}

	return patientDetails, nil
}

func (h *Handler) Delete(ctx *goofy.Context) (interface{}, error) {
	doctorID := ctx.PathParam("doctorId")

	if doctorID == "" {
		return nil, errors.MissingParam{Param: []string{"doctorId"}}
	}

	err := h.svc.Delete(ctx, doctorID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
