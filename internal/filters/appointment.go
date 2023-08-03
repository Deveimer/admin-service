package filters

import (
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"main/utils"
	"time"
)

type Appointment struct {
	Date             string
	Status           string
	DoctorID         string
	PatientID        string
	IncludeCancelled bool
}

func (a *Appointment) Validate() error {
	var err error
	if a.Date != "" {
		_, err = time.Parse("2006-01-02", a.Date)
	}
	if err != nil {
		return errors.InvalidParam{Param: []string{"date"}}
	}

	if a.Status != "" {
		if ok := utils.Validate(a.Status, utils.Enum{Values: []interface{}{"COMPLETED", "PENDING", "CANCELLED", "SCHEDULED"}}); !ok {
			return errors.InvalidParam{Param: []string{"status"}}
		}
	}
	return nil
}
