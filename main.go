package main

import (
	"github.com/Deveimer/goofy/pkg/goofy"

	appointmentHandler "main/internal/http/appointment"
	appointmentService "main/internal/services/appointment"
	appointmentStore "main/internal/stores/appointments"

	doctorOPDScheduleHandler "main/internal/http/doctorOPDSchedule"
	doctorOPDScheduleService "main/internal/services/doctorOPDSchedule"
	doctorOPDScheduleStore "main/internal/stores/doctorOPDSchedule"

	patientsHandler "main/internal/http/patients"
	patientsSvc "main/internal/services/patients"
	patientsStore "main/internal/stores/patients"
)

func main() {

	app := goofy.New()

	// Store Layer
	patientStore := patientsStore.New(app.Database)
	doctorOPDScheduleStore := doctorOPDScheduleStore.New(app.Database)
	appointmentStore := appointmentStore.New(app.Database)

	// Service Layer
	patientSvc := patientsSvc.New(patientStore)
	doctorOPDScheduleSvc := doctorOPDScheduleService.New(doctorOPDScheduleStore)
	appointmentService := appointmentService.New(appointmentStore, doctorOPDScheduleSvc)

	// HTTP Handler
	patientHandler := patientsHandler.New(patientSvc)
	doctorOPDScheduleHandler := doctorOPDScheduleHandler.New(doctorOPDScheduleSvc)
	appointmentHandler := appointmentHandler.New(appointmentService)

	app.GET("/patient/{id}", patientHandler.Get)
	app.POST("/patient", patientHandler.Create)
	app.PUT("/patient/{id}", patientHandler.Update)
	app.DELETE("/patient/{id}", patientHandler.Delete)

	app.POST("/doctor-opd-schedule", doctorOPDScheduleHandler.Create)
	app.GET("/doctor-opd-schedule", doctorOPDScheduleHandler.Index)
	app.GET("/doctor-opd-schedule/{id}", doctorOPDScheduleHandler.Read)
	app.PUT("/doctor-opd-schedule/{id}", doctorOPDScheduleHandler.Update)

	app.POST("/appointment", appointmentHandler.Create)
	app.GET("/appointment", appointmentHandler.Index)
	app.GET("/appointment/{id}", appointmentHandler.Read)
	app.PUT("/appointment/{id}", appointmentHandler.Update)

	app.Start()
}
