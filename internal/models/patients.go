package models

import (
	"encoding/json"
	"time"
)

type PatientRequest struct {
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode int    `json:"pincode"`
	Status  string `json:"status"`
}

type Patient struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Password  string    `json:"-"`
	Salt      string    `json:"-"`
}

type PatientDetails struct {
	Id            string           `json:"id"`
	Name          string           `json:"name"`
	Gender        string           `json:"gender"`
	Phone         string           `json:"phone"`
	Email         string           `json:"email"`
	Age           int              `json:"age"`
	City          string           `json:"city"`
	State         string           `json:"state"`
	Pincode       int              `json:"pincode"`
	JoinedOn      time.Time        `json:"joined_on"`
	LastLoginTime time.Time        `json:"lastLoginTime"`
	MetaData      *json.RawMessage `json:"metaData"`
	Status        string           `json:"status"`
	CreatedAt     time.Time        `json:"createdAt,omitempty"`
	UpdatedAt     time.Time        `json:"updatedAt,omitempty"`
	Password      string           `json:"-"`
	Salt          string           `json:"-"`
}

type PatientPrescription struct {
	ID                   string    `json:"id,omitempty"`
	PatientID            string    `json:"patient_id,omitempty"`
	DoctorID             string    `json:"doctor_id,omitempty"`
	PrescriptionLocation string    `json:"prescription_location,omitempty"`
	Notes                string    `json:"notes,omitempty"`
	CreatedAt            time.Time `json:"createdAt,omitempty"`
	UpdatedAt            time.Time `json:"updatedAt,omitempty"`
	DeletedAt            time.Time `json:"deletedAt,omitempty"`
}

type PatientAppointments struct {
	AppointmentID       string    `json:"appointment_id"`
	PatientID           string    `json:"patient_id"`
	DoctorID            string    `json:"doctor_id"`
	AppointmentLocation string    `json:"appointment_location"`
	AppointmentDate     string    `json:"appointment_date"`
	AppointmentStatus   string    `json:"appointment_status"`
	UpdatedAt           time.Time `json:"updatedAt,omitempty"`
	CreatedAt           time.Time `json:"createdAt,omitempty"`
}

type PatientCaseRecordings struct {
	PatientID string      `json:"patient_id"`
	CaseForms interface{} `json:"case_forms"`
}

type PatientFollowups struct {
	PatientID     string      `json:"patient_id"`
	FollowupForms interface{} `json:"followup_forms"`
}
