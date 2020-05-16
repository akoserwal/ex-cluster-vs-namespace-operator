package controller

import (
	"github.com/akoserwal/ex-cluster-vs-namespace-operatorr/cluster-scope-op/pkg/controller/foo"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, foo.Add)
}
