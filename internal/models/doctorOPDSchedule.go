package models

import (
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"time"
)

type OPDSchedule struct {
	ID              int     `json:"id"`
	DoctorID        string  `json:"doctor_id"`
	OPDStatus       string  `json:"status"`
	OPDStartDate    string  `json:"start_date"`
	OPDEndDate      string  `json:"end_date"`
	OPDStartTime    string  `json:"start_time"`
	OPDCancelReason *string `json:"cancel_reason"`
	OPDEndTime      string  `json:"end_time"`
}

type OPDScheduleCreateRequest struct {
	ID           string `json:"id"`
	OPDStatus    string `json:"status"`
	OPDStartDate string `json:"start_date"`
	OPDEndDate   string `json:"end_date"`
	OPDStartTime string `json:"start_time"`
	OPDEndTime   string `json:"end_time"`
}

func (dos *OPDScheduleCreateRequest) Validate() error {
	var err error
	if dos.ID == "" {
		return errors.MissingParam{Param: []string{"doctor_id"}}
	}

	if dos.OPDStartDate == "" {
		return errors.MissingParam{Param: []string{"start_date"}}
	}

	if dos.OPDEndDate == "" {
		return errors.MissingParam{Param: []string{"start_date"}}
	}

	if dos.OPDStartTime == "" {
		return errors.MissingParam{Param: []string{"start_time"}}
	}

	if dos.OPDEndTime == "" {
		return errors.MissingParam{Param: []string{"end_time"}}
	}

	startDate, err := time.Parse("2006-01-02", dos.OPDStartDate)
	if err != nil {
		return errors.InvalidParam{Param: []string{"start_date"}}
	}

	endDate, err := time.Parse("2006-01-02", dos.OPDEndDate)
	if err != nil {
		return errors.InvalidParam{Param: []string{"end_date"}}
	}

	startTime, err := time.Parse("15:04:05", dos.OPDStartTime)
	if err != nil {
		return errors.InvalidParam{Param: []string{"start_time"}}
	}

	endTime, err := time.Parse("15:04:05", dos.OPDEndTime)
	if err != nil {
		return errors.InvalidParam{Param: []string{"end_time"}}
	}

	if startDate.After(endDate) {
		return &errors.Response{Reason: "start Date cannot after end Date"}
	}

	if startTime.After(endTime) {
		return &errors.Response{Reason: "start time cannot after end time"}
	}

	return nil
}
