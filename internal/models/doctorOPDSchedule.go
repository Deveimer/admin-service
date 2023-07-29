package models

import (
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"time"
)

type DoctorOPDSchedule struct {
	ID           int    `json:"id"`
	DoctorID     string `json:"doctorId"`
	OPDStatus    string `json:"opdStatus"`
	OPDStartDate string `json:"opdStartDate"`
	OPDEndDate   string `json:"opdEndDte"`
	OPDStartTime string `json:"opdStartTime"`
	OPDEndTime   string `json:"opdEndTime"`
}

type DoctorOPDScheduleCreateRequest struct {
	DoctorID     string `json:"doctorId"`
	OPDStatus    string `json:"opdStatus"`
	OPDStartDate string `json:"opdStartDate"`
	OPDEndDate   string `json:"opdEndDate"`
	OPDStartTime string `json:"opdStartTime"`
	OPDEndTime   string `json:"opdEndTime"`
}

func (dos *DoctorOPDScheduleCreateRequest) Validate() error {
	var err error
	if dos.DoctorID == "" {
		return errors.MissingParam{Param: []string{"doctorId"}}
	}

	if dos.OPDStartDate == "" {
		return errors.MissingParam{Param: []string{"opdStartDate"}}
	}

	if dos.OPDEndDate == "" {
		return errors.MissingParam{Param: []string{"opdStartDate"}}
	}

	if dos.OPDStartTime == "" {
		return errors.MissingParam{Param: []string{"opdStartTime"}}
	}

	if dos.OPDEndTime == "" {
		return errors.MissingParam{Param: []string{"opdEndTime"}}
	}

	startDate, err := time.Parse("2006-01-02", dos.OPDStartDate)
	if err != nil {
		return errors.InvalidParam{Param: []string{"opdStartDate"}}
	}

	endDate, err := time.Parse("2006-01-02", dos.OPDEndDate)
	if err != nil {
		return errors.InvalidParam{Param: []string{"opdEndDate"}}
	}

	startTime, err := time.Parse("15:04:05", dos.OPDStartTime)
	if err != nil {
		return errors.InvalidParam{Param: []string{"opdStartTime"}}
	}

	endTime, err := time.Parse("15:04:05", dos.OPDEndTime)
	if err != nil {
		return errors.InvalidParam{Param: []string{"opdEndTime"}}
	}

	if startDate.After(endDate) {
		return &errors.Response{Reason: "start Date cannot after end Date"}
	}

	if startTime.After(endTime) {
		return &errors.Response{Reason: "start time cannot after end time"}
	}

	return nil
}
