package filters

import (
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"main/internal/utils"
	"time"
)

type DoctorOPDSchedule struct {
	DoctorID  string
	StartDate string
	EndDate   string
	Date      string
	Status    string
}

func (dos *DoctorOPDSchedule) Validate() error {
	var (
		err                error
		startDate, endDate time.Time
	)
	if dos.StartDate != "" && dos.EndDate == "" {
		return errors.MissingParam{Param: []string{"EndDate"}}
	}

	if dos.StartDate == "" && dos.EndDate != "" {
		return errors.MissingParam{Param: []string{"StartDate"}}
	}

	if dos.StartDate != "" {
		startDate, err = time.Parse("2006-01-02", dos.StartDate)
	}

	if err != nil {
		return errors.InvalidParam{Param: []string{"startDate"}}
	}

	if dos.EndDate != "" {
		endDate, err = time.Parse("2006-01-02", dos.EndDate)
	}
	if err != nil {
		return errors.InvalidParam{Param: []string{"endDate"}}
	}

	if dos.StartDate != "" && dos.EndDate != "" && startDate.After(endDate) {
		return &errors.Response{Reason: "start date cannot be greater than end date"}
	}

	if dos.Date != "" {
		_, err = time.Parse("2006-01-02", dos.Date)
	}

	if err != nil {
		return errors.InvalidParam{Param: []string{"date"}}
	}

	if dos.Status != "" {
		if ok := utils.Validate(dos.Status, utils.Enum{Values: []interface{}{"SCHEDULED", "OPEN", "CANCELLED"}}); !ok {
			return errors.InvalidParam{Param: []string{"status"}}
		}
	}

	return nil
}
