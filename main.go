package main

import (
	"github.com/Deveimer/goofy/pkg/goofy"

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

	// Service Layer
	patientSvc := patientsSvc.New(patientStore)
	doctorService := doctorsSvc.New(doctorStore)

	// HTTP Handler
	patientHandler := patientsHandler.New(patientSvc)
	doctorHandler := doctorsHandler.New(doctorService)

	// Patient Endpoints
	app.GET("/patient/{id}", patientHandler.Get)
	app.POST("/patient", patientHandler.Create)
	app.PUT("/patient/{id}", patientHandler.Update)
	app.DELETE("/patient/{id}", patientHandler.Delete)

	// Doctor Eeandpoints
	app.GET("/doctor/{id}", doctorHandler.Get)
	app.POST("/doctor", doctorHandler.Create)
	app.PUT("/doctor/{id}", doctorHandler.Update)
	app.DELETE("/doctor/{id}", doctorHandler.Delete)

	app.Start()
}
