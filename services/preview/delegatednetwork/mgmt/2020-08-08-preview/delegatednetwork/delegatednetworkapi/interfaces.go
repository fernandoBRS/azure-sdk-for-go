package delegatednetworkapi

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

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/delegatednetwork/mgmt/2020-08-08-preview/delegatednetwork"
)

// ControllerClientAPI contains the set of methods on the ControllerClient type.
type ControllerClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, resourceName string, parameters delegatednetwork.DelegatedController) (result delegatednetwork.ControllerCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result delegatednetwork.ControllerDeleteFuture, err error)
	GetDetails(ctx context.Context, resourceGroupName string, resourceName string) (result delegatednetwork.DelegatedController, err error)
	Patch(ctx context.Context, resourceGroupName string, resourceName string, parameters delegatednetwork.ControllerResourceUpdateParameters) (result delegatednetwork.DelegatedController, err error)
}

var _ ControllerClientAPI = (*delegatednetwork.ControllerClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result delegatednetwork.DelegatedControllersPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result delegatednetwork.DelegatedControllersIterator, err error)
	ListBySubscription(ctx context.Context) (result delegatednetwork.DelegatedControllersPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result delegatednetwork.DelegatedControllersIterator, err error)
}

var _ ClientAPI = (*delegatednetwork.Client)(nil)

// OrchestratorInstanceServiceClientAPI contains the set of methods on the OrchestratorInstanceServiceClient type.
type OrchestratorInstanceServiceClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, resourceName string, parameters delegatednetwork.Orchestrator) (result delegatednetwork.OrchestratorInstanceServiceCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result delegatednetwork.OrchestratorInstanceServiceDeleteFuture, err error)
	GetDetails(ctx context.Context, resourceGroupName string, resourceName string) (result delegatednetwork.Orchestrator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result delegatednetwork.OrchestratorsPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result delegatednetwork.OrchestratorsIterator, err error)
	ListBySubscription(ctx context.Context) (result delegatednetwork.OrchestratorsPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result delegatednetwork.OrchestratorsIterator, err error)
	Patch(ctx context.Context, resourceGroupName string, resourceName string, parameters delegatednetwork.OrchestratorResourceUpdateParameters) (result delegatednetwork.Orchestrator, err error)
}

var _ OrchestratorInstanceServiceClientAPI = (*delegatednetwork.OrchestratorInstanceServiceClient)(nil)

// DelegatedSubnetServiceClientAPI contains the set of methods on the DelegatedSubnetServiceClient type.
type DelegatedSubnetServiceClientAPI interface {
	DeleteDetails(ctx context.Context, resourceGroupName string, resourceName string) (result delegatednetwork.DelegatedSubnetServiceDeleteDetailsFuture, err error)
	GetDetails(ctx context.Context, resourceGroupName string, resourceName string) (result delegatednetwork.DelegatedSubnet, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result delegatednetwork.DelegatedSubnetsPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result delegatednetwork.DelegatedSubnetsIterator, err error)
	ListBySubscription(ctx context.Context) (result delegatednetwork.DelegatedSubnetsPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result delegatednetwork.DelegatedSubnetsIterator, err error)
	PatchDetails(ctx context.Context, resourceGroupName string, resourceName string, parameters delegatednetwork.ResourceUpdateParameters) (result delegatednetwork.DelegatedSubnetServicePatchDetailsFuture, err error)
	PutDetails(ctx context.Context, resourceGroupName string, resourceName string, parameters delegatednetwork.DelegatedSubnet) (result delegatednetwork.DelegatedSubnetServicePutDetailsFuture, err error)
}

var _ DelegatedSubnetServiceClientAPI = (*delegatednetwork.DelegatedSubnetServiceClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result delegatednetwork.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result delegatednetwork.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*delegatednetwork.OperationsClient)(nil)
