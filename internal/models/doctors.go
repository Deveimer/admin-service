package models

import (
	"encoding/json"
	"time"
)

type Doctor struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Password  string    `json:"-"`
	Salt      string    `json:"-"`
}

type DoctorDetails struct {
	Id            string           `json:"id"`
	Name          string           `json:"name"`
	Gender        string           `json:"gender"`
	LicenceNumber string           `json:"licence_number"`
	Phone         string           `json:"phone"`
	Email         string           `json:"email"`
	Age           int              `json:"age"`
	City          string           `json:"city"`
	State         string           `json:"state"`
	Pincode       string           `json:"pincode"`
	JoinedTime    time.Time        `json:"joined_time"`
	JoinedOn      time.Time        `json:"joined_on"`
	LastLoginTime time.Time        `json:"lastLoginTime"`
	MetaData      *json.RawMessage `json:"metaData"`
	Status        string           `json:"status"`
	CreatedAt     time.Time        `json:"createdAt,omitempty"`
	UpdatedAt     time.Time        `json:"updatedAt,omitempty"`
	Password      string           `json:"-"`
	Salt          string           `json:"-"`
}

type DoctorAppointments struct {
	AppointmentID       string    `json:"appointment_id"`
	PatientID           string    `json:"patient_id"`
	DoctorID            string    `json:"doctor_id"`
	AppointmentLocation string    `json:"appointment_location"`
	AppointmentDate     string    `json:"appointment_date"`
	AppointmentStatus   string    `json:"appointment_status"`
	UpdatedAt           time.Time `json:"updatedAt,omitempty"`
	CreatedAt           time.Time `json:"createdAt,omitempty"`
}

type DoctorOPDSchedule struct {
	DoctorID     string    `json:"doctor_id"`
	OPDStatus    string    `json:"opd_status"`
	OPDStartDate time.Time `json:"opd_start_date"`
	OPDEndDate   time.Time `json:"opd_end_date"`
	OPDStartTime time.Time `json:"opd_start_time"`
	OPDEndTime   time.Time `json:"opd_end_time"`
}
