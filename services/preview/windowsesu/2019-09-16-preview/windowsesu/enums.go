package windowsesu

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

// OsType enumerates the values for os type.
type OsType string

const (
	// Windows7 ...
	Windows7 OsType = "Windows7"
	// WindowsServer2008 ...
	WindowsServer2008 OsType = "WindowsServer2008"
	// WindowsServer2008R2 ...
	WindowsServer2008R2 OsType = "WindowsServer2008R2"
)

// PossibleOsTypeValues returns an array of possible values for the OsType const type.
func PossibleOsTypeValues() []OsType {
	return []OsType{Windows7, WindowsServer2008, WindowsServer2008R2}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Accepted ...
	Accepted ProvisioningState = "Accepted"
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Provisioning ...
	Provisioning ProvisioningState = "Provisioning"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Accepted, Canceled, Failed, Provisioning, Succeeded}
}

// SupportType enumerates the values for support type.
type SupportType string

const (
	// PremiumAssurance ...
	PremiumAssurance SupportType = "PremiumAssurance"
	// SupplementalServicing ...
	SupplementalServicing SupportType = "SupplementalServicing"
)

// PossibleSupportTypeValues returns an array of possible values for the SupportType const type.
func PossibleSupportTypeValues() []SupportType {
	return []SupportType{PremiumAssurance, SupplementalServicing}
}
