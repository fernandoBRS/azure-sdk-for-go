package media

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

// EntityNameUnavailabilityReason enumerates the values for entity name unavailability reason.
type EntityNameUnavailabilityReason string

const (
	// AlreadyExists ...
	AlreadyExists EntityNameUnavailabilityReason = "AlreadyExists"
	// Invalid ...
	Invalid EntityNameUnavailabilityReason = "Invalid"
	// None ...
	None EntityNameUnavailabilityReason = "None"
)

// PossibleEntityNameUnavailabilityReasonValues returns an array of possible values for the EntityNameUnavailabilityReason const type.
func PossibleEntityNameUnavailabilityReasonValues() []EntityNameUnavailabilityReason {
	return []EntityNameUnavailabilityReason{AlreadyExists, Invalid, None}
}

// KeyType enumerates the values for key type.
type KeyType string

const (
	// Primary ...
	Primary KeyType = "Primary"
	// Secondary ...
	Secondary KeyType = "Secondary"
)

// PossibleKeyTypeValues returns an array of possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{Primary, Secondary}
}

// ResourceType enumerates the values for resource type.
type ResourceType string

const (
	// Mediaservices ...
	Mediaservices ResourceType = "mediaservices"
)

// PossibleResourceTypeValues returns an array of possible values for the ResourceType const type.
func PossibleResourceTypeValues() []ResourceType {
	return []ResourceType{Mediaservices}
}
