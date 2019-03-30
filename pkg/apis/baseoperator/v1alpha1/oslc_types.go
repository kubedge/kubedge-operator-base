package v1alpha1

import ()

// Phase of the Openstack Service Life Cyle
type OslcPhase string

// Describe the Phase of the Openstack Service Life Cycle
const (
	PhaseTest     OslcPhase = "test"
	PhaseUpgrade  OslcPhase = "upgrade"
	PhaseRollback OslcPhase = "roolback"
)

// String converts a OslcPhase to a printable string
func (x OslcPhase) String() string { return string(x) }
