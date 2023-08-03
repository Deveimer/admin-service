package models

import (
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"main/internal/utils"
	"time"
)

type Appointment struct {
	ID          int     `json:"ID"`
	PatientID   string  `json:"patient_id"`
	DoctorID    string  `json:"doctor_id"`
	StartTime   string  `json:"start_time"`
	EndTime     string  `json:"end_time"`
	Description *string `json:"description"`
	Status      string  `json:"status"`
	Date        string  `json:"date"`
	Reason      *string `json:"reason"`
}

type AppointmentCreateRequest struct {
	PatientID   string  `json:"patient_id"`
	DoctorID    string  `json:"doctor_id"`
	StartTime   string  `json:"start_time"`
	Status      string  `json:"status"`
	EndTime     string  `json:"end_time"`
	Description *string `json:"description"`
	Date        string  `json:"date"`
	CreatedBy   string  `json:"created_by"`
}

type AppointmentUpdateRequest struct {
	ID          int    `json:"-"`
	Reason      string `json:"reason"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func (a *AppointmentCreateRequest) Validate() error {
	var err error
	if a.DoctorID == "" {
		return errors.MissingParam{Param: []string{"doctorId"}}
	}

	if a.PatientID == "" {
		return errors.MissingParam{Param: []string{"patientId"}}
	}

	if a.StartTime == "" {
		return errors.MissingParam{Param: []string{"startTime"}}
	}

	if a.EndTime == "" {
		return errors.MissingParam{Param: []string{"endTime"}}
	}

	if a.Date == "" {
		return errors.MissingParam{Param: []string{"date"}}
	}

	if a.Status != "" {
		if ok := utils.Validate(a.Status, utils.Enum{Values: []interface{}{"SCHEDULED", "PENDING"}}); !ok {
			return errors.InvalidParam{Param: []string{"status"}}
		}
	}

	_, err = time.Parse("2006-01-02", a.Date)
	if err != nil {
		return errors.InvalidParam{Param: []string{"date"}}
	}

	startTime, err := time.Parse("15:04:05", a.StartTime)
	if err != nil {
		return errors.InvalidParam{Param: []string{"startTime"}}
	}

	endTime, err := time.Parse("15:04:05", a.EndTime)
	if err != nil {
		return errors.InvalidParam{Param: []string{"endTime"}}
	}

	if startTime.After(endTime) {
		return &errors.Response{Reason: "start time cannot after end time"}
	}

	return nil
}

func (a *AppointmentUpdateRequest) Validate() error {
	if a.Status != "" {
		if ok := utils.Validate(a.Status, utils.Enum{Values: []interface{}{"SCHEDULED", "PENDING", "CANCELLED", "COMPLETED"}}); !ok {
			return errors.InvalidParam{Param: []string{"status"}}
		}
	}

	if a.Status == "CANCELLED" && a.Reason == "" {
		return errors.MissingParam{Param: []string{"reason"}}
	}

	return nil
}
