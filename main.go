package main

import (
	"github.com/Deveimer/goofy/pkg/goofy"

	doctorOPDScheduleHandler "main/internal/http/opdScheduler"
	doctorOPDScheduleService "main/internal/services/opdScheduler"
	doctorOPDScheduleStore "main/internal/stores/opdScheduler"

	patientsHandler "main/internal/http/patients"
	patientsSvc "main/internal/services/patients"
	patientsStore "main/internal/stores/patients"

	doctorsHandler "main/internal/http/doctors"
	doctorsSvc "main/internal/services/doctors"
	doctorsStore "main/internal/stores/doctors"
)

func main() {

	app := goofy.New()

	// Store Layer
	patientStore := patientsStore.New(app.Database)
	doctorStore := doctorsStore.New(app.Database)
	opdScheduleStore := doctorOPDScheduleStore.New(app.Database)

	// Service Layer
	patientSvc := patientsSvc.New(patientStore)
	doctorService := doctorsSvc.New(doctorStore)
	opdScheduleHandlerService := doctorOPDScheduleService.New(opdScheduleStore)

	// HTTP Handler
	patientHandler := patientsHandler.New(patientSvc)
	doctorHandler := doctorsHandler.New(doctorService)
	opdScheduleHandler := doctorOPDScheduleHandler.New(opdScheduleHandlerService)

	// Patient Endpoints
	app.GET("/patient/{id}", patientHandler.Get)
	app.POST("/patient", patientHandler.Create)
	app.PUT("/patient/{id}", patientHandler.Update)
	app.DELETE("/patient/{id}", patientHandler.Delete)

	// Doctor Endpoints
	app.GET("/doctor/{id}", doctorHandler.Get)
	app.POST("/doctor", doctorHandler.Create)
	app.PUT("/doctor/{id}", doctorHandler.Update)
	app.DELETE("/doctor/{id}", doctorHandler.Delete)

	app.POST("/doctor-opd-schedule", opdScheduleHandler.Create)
	app.GET("/doctor-opd-schedule", opdScheduleHandler.Index)
	app.GET("/doctor-opd-schedule/{id}", opdScheduleHandler.Read)
	app.PUT("/doctor-opd-schedule/{id}", opdScheduleHandler.Update)

	app.Start()
}
