package aadapi

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
	"github.com/Azure/azure-sdk-for-go/services/domainservices/mgmt/2020-01-01/aad"
)

// DomainServiceOperationsClientAPI contains the set of methods on the DomainServiceOperationsClient type.
type DomainServiceOperationsClientAPI interface {
	List(ctx context.Context) (result aad.OperationEntityListResultPage, err error)
	ListComplete(ctx context.Context) (result aad.OperationEntityListResultIterator, err error)
}

var _ DomainServiceOperationsClientAPI = (*aad.DomainServiceOperationsClient)(nil)

// DomainServicesClientAPI contains the set of methods on the DomainServicesClient type.
type DomainServicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, domainServiceName string, domainService aad.DomainService) (result aad.DomainServicesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, domainServiceName string) (result aad.DomainServicesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, domainServiceName string) (result aad.DomainService, err error)
	List(ctx context.Context) (result aad.DomainServiceListResultPage, err error)
	ListComplete(ctx context.Context) (result aad.DomainServiceListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result aad.DomainServiceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result aad.DomainServiceListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, domainServiceName string, domainService aad.DomainService) (result aad.DomainServicesUpdateFuture, err error)
}

var _ DomainServicesClientAPI = (*aad.DomainServicesClient)(nil)

// OuContainerOperationsClientAPI contains the set of methods on the OuContainerOperationsClient type.
type OuContainerOperationsClientAPI interface {
	List(ctx context.Context) (result aad.OperationEntityListResultPage, err error)
	ListComplete(ctx context.Context) (result aad.OperationEntityListResultIterator, err error)
}

var _ OuContainerOperationsClientAPI = (*aad.OuContainerOperationsClient)(nil)

// OuContainerClientAPI contains the set of methods on the OuContainerClient type.
type OuContainerClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount aad.ContainerAccount) (result aad.OuContainerCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string) (result aad.OuContainerDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string) (result aad.OuContainer, err error)
	List(ctx context.Context, resourceGroupName string, domainServiceName string) (result aad.OuContainerListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, domainServiceName string) (result aad.OuContainerListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount aad.ContainerAccount) (result aad.OuContainerUpdateFuture, err error)
}

var _ OuContainerClientAPI = (*aad.OuContainerClient)(nil)
