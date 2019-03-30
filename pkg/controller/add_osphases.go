package controller

import (
	"github.com/kubedge/kubedge-operator-base/pkg/controller/basecontroller"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, basecontroller.AddArpscanController)
	AddToManagerFuncs = append(AddToManagerFuncs, basecontroller.AddECDSClusterController)
	AddToManagerFuncs = append(AddToManagerFuncs, basecontroller.AddMMESimController)
}
