package main

import (
	"github.com/Deveimer/goofy/pkg/goofy"
)

func main() {

	app := goofy.New()

	// To pass database instance just pass app.Database and make sure to have DB configs in .local.env (DAB_PORT should be used instead of DB_PORT)

	app.Start()
}
