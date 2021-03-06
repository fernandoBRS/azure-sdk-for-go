package migrate

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Bitness enumerates the values for bitness.
type Bitness string

const (
	// SixFourbit ...
	SixFourbit Bitness = "64bit"
	// ThreeTwobit ...
	ThreeTwobit Bitness = "32bit"
)

// PossibleBitnessValues returns an array of possible values for the Bitness const type.
func PossibleBitnessValues() []Bitness {
	return []Bitness{SixFourbit, ThreeTwobit}
}

// CPUSpeedAccuracy enumerates the values for cpu speed accuracy.
type CPUSpeedAccuracy string

const (
	// Actual ...
	Actual CPUSpeedAccuracy = "actual"
	// Estimated ...
	Estimated CPUSpeedAccuracy = "estimated"
)

// PossibleCPUSpeedAccuracyValues returns an array of possible values for the CPUSpeedAccuracy const type.
func PossibleCPUSpeedAccuracyValues() []CPUSpeedAccuracy {
	return []CPUSpeedAccuracy{Actual, Estimated}
}

// CredentialType enumerates the values for credential type.
type CredentialType string

const (
	// HyperVFabric ...
	HyperVFabric CredentialType = "HyperVFabric"
	// LinuxGuest ...
	LinuxGuest CredentialType = "LinuxGuest"
	// LinuxServer ...
	LinuxServer CredentialType = "LinuxServer"
	// VMwareFabric ...
	VMwareFabric CredentialType = "VMwareFabric"
	// WindowsGuest ...
	WindowsGuest CredentialType = "WindowsGuest"
	// WindowsServer ...
	WindowsServer CredentialType = "WindowsServer"
)

// PossibleCredentialTypeValues returns an array of possible values for the CredentialType const type.
func PossibleCredentialTypeValues() []CredentialType {
	return []CredentialType{HyperVFabric, LinuxGuest, LinuxServer, VMwareFabric, WindowsGuest, WindowsServer}
}

// Family enumerates the values for family.
type Family string

const (
	// Aix ...
	Aix Family = "aix"
	// Linux ...
	Linux Family = "linux"
	// Solaris ...
	Solaris Family = "solaris"
	// Unknown ...
	Unknown Family = "unknown"
	// Windows ...
	Windows Family = "windows"
)

// PossibleFamilyValues returns an array of possible values for the Family const type.
func PossibleFamilyValues() []Family {
	return []Family{Aix, Linux, Solaris, Unknown, Windows}
}

// HighlyAvailable enumerates the values for highly available.
type HighlyAvailable string

const (
	// HighlyAvailableNo ...
	HighlyAvailableNo HighlyAvailable = "No"
	// HighlyAvailableUnknown ...
	HighlyAvailableUnknown HighlyAvailable = "Unknown"
	// HighlyAvailableYes ...
	HighlyAvailableYes HighlyAvailable = "Yes"
)

// PossibleHighlyAvailableValues returns an array of possible values for the HighlyAvailable const type.
func PossibleHighlyAvailableValues() []HighlyAvailable {
	return []HighlyAvailable{HighlyAvailableNo, HighlyAvailableUnknown, HighlyAvailableYes}
}

// HypervisorType enumerates the values for hypervisor type.
type HypervisorType string

const (
	// HypervisorTypeHyperv ...
	HypervisorTypeHyperv HypervisorType = "hyperv"
	// HypervisorTypeUnknown ...
	HypervisorTypeUnknown HypervisorType = "unknown"
)

// PossibleHypervisorTypeValues returns an array of possible values for the HypervisorType const type.
func PossibleHypervisorTypeValues() []HypervisorType {
	return []HypervisorType{HypervisorTypeHyperv, HypervisorTypeUnknown}
}

// PropertiesMonitoringState enumerates the values for properties monitoring state.
type PropertiesMonitoringState string

const (
	// Discovered ...
	Discovered PropertiesMonitoringState = "discovered"
	// Monitored ...
	Monitored PropertiesMonitoringState = "monitored"
)

// PossiblePropertiesMonitoringStateValues returns an array of possible values for the PropertiesMonitoringState const type.
func PossiblePropertiesMonitoringStateValues() []PropertiesMonitoringState {
	return []PropertiesMonitoringState{Discovered, Monitored}
}

// PropertiesVirtualizationState enumerates the values for properties virtualization state.
type PropertiesVirtualizationState string

const (
	// PropertiesVirtualizationStateHypervisor ...
	PropertiesVirtualizationStateHypervisor PropertiesVirtualizationState = "hypervisor"
	// PropertiesVirtualizationStatePhysical ...
	PropertiesVirtualizationStatePhysical PropertiesVirtualizationState = "physical"
	// PropertiesVirtualizationStateUnknown ...
	PropertiesVirtualizationStateUnknown PropertiesVirtualizationState = "unknown"
	// PropertiesVirtualizationStateVirtual ...
	PropertiesVirtualizationStateVirtual PropertiesVirtualizationState = "virtual"
)

// PossiblePropertiesVirtualizationStateValues returns an array of possible values for the PropertiesVirtualizationState const type.
func PossiblePropertiesVirtualizationStateValues() []PropertiesVirtualizationState {
	return []PropertiesVirtualizationState{PropertiesVirtualizationStateHypervisor, PropertiesVirtualizationStatePhysical, PropertiesVirtualizationStateUnknown, PropertiesVirtualizationStateVirtual}
}

// RebootStatus enumerates the values for reboot status.
type RebootStatus string

const (
	// RebootStatusNotRebooted ...
	RebootStatusNotRebooted RebootStatus = "notRebooted"
	// RebootStatusRebooted ...
	RebootStatusRebooted RebootStatus = "rebooted"
	// RebootStatusUnknown ...
	RebootStatusUnknown RebootStatus = "unknown"
)

// PossibleRebootStatusValues returns an array of possible values for the RebootStatus const type.
func PossibleRebootStatusValues() []RebootStatus {
	return []RebootStatus{RebootStatusNotRebooted, RebootStatusRebooted, RebootStatusUnknown}
}

// VirtualDiskMode enumerates the values for virtual disk mode.
type VirtualDiskMode string

const (
	// Append ...
	Append VirtualDiskMode = "append"
	// IndependentNonpersistent ...
	IndependentNonpersistent VirtualDiskMode = "independent_nonpersistent"
	// IndependentPersistent ...
	IndependentPersistent VirtualDiskMode = "independent_persistent"
	// Nonpersistent ...
	Nonpersistent VirtualDiskMode = "nonpersistent"
	// Persistent ...
	Persistent VirtualDiskMode = "persistent"
	// Undoable ...
	Undoable VirtualDiskMode = "undoable"
)

// PossibleVirtualDiskModeValues returns an array of possible values for the VirtualDiskMode const type.
func PossibleVirtualDiskModeValues() []VirtualDiskMode {
	return []VirtualDiskMode{Append, IndependentNonpersistent, IndependentPersistent, Nonpersistent, Persistent, Undoable}
}

// VirtualMachineType enumerates the values for virtual machine type.
type VirtualMachineType string

const (
	// VirtualMachineTypeHyperv ...
	VirtualMachineTypeHyperv VirtualMachineType = "hyperv"
	// VirtualMachineTypeLdom ...
	VirtualMachineTypeLdom VirtualMachineType = "ldom"
	// VirtualMachineTypeLpar ...
	VirtualMachineTypeLpar VirtualMachineType = "lpar"
	// VirtualMachineTypeUnknown ...
	VirtualMachineTypeUnknown VirtualMachineType = "unknown"
	// VirtualMachineTypeVirtualPc ...
	VirtualMachineTypeVirtualPc VirtualMachineType = "virtualPc"
	// VirtualMachineTypeVmware ...
	VirtualMachineTypeVmware VirtualMachineType = "vmware"
	// VirtualMachineTypeXen ...
	VirtualMachineTypeXen VirtualMachineType = "xen"
)

// PossibleVirtualMachineTypeValues returns an array of possible values for the VirtualMachineType const type.
func PossibleVirtualMachineTypeValues() []VirtualMachineType {
	return []VirtualMachineType{VirtualMachineTypeHyperv, VirtualMachineTypeLdom, VirtualMachineTypeLpar, VirtualMachineTypeUnknown, VirtualMachineTypeVirtualPc, VirtualMachineTypeVmware, VirtualMachineTypeXen}
}
