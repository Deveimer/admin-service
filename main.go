package main

import (
	"github.com/Deveimer/goofy/pkg/goofy"

	patientsHandler "main/internal/http/patients"
	prescriptionHandler "main/internal/http/patients/prescription"
	patientsSvc "main/internal/services/patients"
	prescriptionSvc "main/internal/services/patients/prescription"
	patientsStore "main/internal/stores/patients"
	prescriptionStore "main/internal/stores/patients/prescription"
)

func main() {

	app := goofy.New()

	// Store Layer
	patientStore := patientsStore.New(app.Database)
	prescriptionStore := prescriptionStore.New(app.Database)

	// Service Layer
	patientSvc := patientsSvc.New(patientStore)
	prescriptionSvc := prescriptionSvc.New(prescriptionStore)

	// HTTP Handler
	patientHandler := patientsHandler.New(patientSvc)
	prescriptionHandler := prescriptionHandler.New(prescriptionSvc)

	app.GET("/patient/{id}", patientHandler.Get)
	app.POST("/patient", patientHandler.Create)
	app.PUT("/patient/{id}", patientHandler.Update)
	app.DELETE("/patient/{id}", patientHandler.Delete)
	app.POST("/prescription", prescriptionHandler.Create)
	app.GET("/prescription", prescriptionHandler.Get)
	app.PUT("/prescription", prescriptionHandler.Update)
	app.DELETE("/prescription", prescriptionHandler.Delete)

	app.Start()
}
