package main

import (
	"github.com/Deveimer/goofy/pkg/goofy"

	patientsHandler "main/internal/http/patients"
	patientsSvc "main/internal/services/patients"
	patientsStore "main/internal/stores/patients"
)

func main() {

	app := goofy.New()

	// Store Layer
	patientStore := patientsStore.New(app.Database)

	// Service Layer
	patientSvc := patientsSvc.New(patientStore)

	// HTTP Handler
	patientHandler := patientsHandler.New(patientSvc)

	app.POST("/patient", patientHandler.Create)

	app.Start()
}
