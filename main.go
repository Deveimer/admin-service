package main

import (
	"github.com/Deveimer/goofy/pkg/goofy"
	patientsHandler "main/internal/http/patients"
	patientsSvc "main/internal/services/patients"
	patientsStore "main/internal/stores/patients"
)

func main() {

	app := goofy.New()

	patientStore := patientsStore.New(app.Database)
	patientSvc := patientsSvc.New(patientStore)
	patientHandler := patientsHandler.New(patientSvc)

	app.POST("/patient", patientHandler.Create)

	// To pass database instance just pass app.Database and make sure to have DB configs in .local.env (DAB_PORT should be used instead of DB_PORT)

	app.Start()
}
