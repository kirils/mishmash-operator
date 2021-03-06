package controller

import (
	"github.com/kirils/mishmash-operator/pkg/controller/mishmash"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, mishmash.Add)
}
