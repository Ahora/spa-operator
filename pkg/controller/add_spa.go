package controller

import (
	"github.com/ahora/spa-operator/pkg/controller/spa"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, spa.Add)
}
