package filters

import (
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"main/utils"
	"time"
)

type DoctorOPDSchedule struct {
	DoctorID  string
	StartDate string
	EndDate   string
	Date      string
	Status    string
}

func (d *DoctorOPDSchedule) Validate() error {
	var (
		err                error
		startDate, endDate time.Time
	)
	if d.StartDate != "" && d.EndDate == "" {
		return errors.MissingParam{Param: []string{"EndDate"}}
	}

	if d.StartDate == "" && d.EndDate != "" {
		return errors.MissingParam{Param: []string{"StartDate"}}
	}

	if d.StartDate != "" {
		startDate, err = time.Parse("2006-01-02", d.StartDate)
	}

	if err != nil {
		return errors.InvalidParam{Param: []string{"startDate"}}
	}

	if d.EndDate != "" {
		endDate, err = time.Parse("2006-01-02", d.EndDate)
	}
	if err != nil {
		return errors.InvalidParam{Param: []string{"endDate"}}
	}

	if d.StartDate != "" && d.EndDate != "" && startDate.After(endDate) {
		return &errors.Response{Reason: "start date cannot be greater than end date"}
	}

	if d.Date != "" {
		_, err = time.Parse("2006-01-02", d.Date)
	}

	if err != nil {
		return errors.InvalidParam{Param: []string{"date"}}
	}

	if d.Status != "" {
		if ok := utils.Validate(d.Status, utils.Enum{Values: []interface{}{"SCHEDULED", "OPEN", "CANCELLED"}}); !ok {
			return errors.InvalidParam{Param: []string{"status"}}
		}
	}

	return nil
}
