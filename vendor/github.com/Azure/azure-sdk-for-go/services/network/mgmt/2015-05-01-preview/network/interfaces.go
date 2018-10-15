package network

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// InterfacesClient is the network Client
type InterfacesClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// NewInterfacesClient creates an instance of the InterfacesClient client.
func NewInterfacesClient(subscriptionID string) InterfacesClient {
	return NewInterfacesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// NewInterfacesClientWithBaseURI creates an instance of the InterfacesClient client.
func NewInterfacesClientWithBaseURI(baseURI string, subscriptionID string) InterfacesClient {
	return InterfacesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// CreateOrUpdate the Put NetworkInterface operation creates/updates a networkInterface
//
// resourceGroupName is the name of the resource group. networkInterfaceName is the name of the network interface.
// parameters is parameters supplied to the create/update NetworkInterface operation
func (client InterfacesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string, parameters Interface) (result InterfacesCreateOrUpdateFuture, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, networkInterfaceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client InterfacesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, networkInterfaceName string, parameters Interface) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) CreateOrUpdateSender(req *http.Request) (future InterfacesCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client InterfacesClient) CreateOrUpdateResponder(resp *http.Response) (result Interface, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// Delete the delete netwokInterface operation deletes the specified netwokInterface.
//
// resourceGroupName is the name of the resource group. networkInterfaceName is the name of the network interface.
func (client InterfacesClient) Delete(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result InterfacesDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, networkInterfaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// DeletePreparer prepares the Delete request.
func (client InterfacesClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkInterfaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) DeleteSender(req *http.Request) (future InterfacesDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client InterfacesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// Get the Get ntework interface operation retreives information about the specified network interface.
//
// resourceGroupName is the name of the resource group. networkInterfaceName is the name of the network interface.
func (client InterfacesClient) Get(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result Interface, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, networkInterfaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// GetPreparer prepares the Get request.
func (client InterfacesClient) GetPreparer(ctx context.Context, resourceGroupName string, networkInterfaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InterfacesClient) GetResponder(resp *http.Response) (result Interface, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// GetVirtualMachineScaleSetNetworkInterface the Get ntework interface operation retreives information about the
// specified network interface in a virtual machine scale set.
//
// resourceGroupName is the name of the resource group. virtualMachineScaleSetName is the name of the virtual
// machine scale set. virtualmachineIndex is the virtual machine index. networkInterfaceName is the name of the
// network interface.
func (client InterfacesClient) GetVirtualMachineScaleSetNetworkInterface(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string) (result Interface, err error) {
	req, err := client.GetVirtualMachineScaleSetNetworkInterfacePreparer(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "GetVirtualMachineScaleSetNetworkInterface", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetVirtualMachineScaleSetNetworkInterfaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "GetVirtualMachineScaleSetNetworkInterface", resp, "Failure sending request")
		return
	}

	result, err = client.GetVirtualMachineScaleSetNetworkInterfaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "GetVirtualMachineScaleSetNetworkInterface", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// GetVirtualMachineScaleSetNetworkInterfacePreparer prepares the GetVirtualMachineScaleSetNetworkInterface request.
func (client InterfacesClient) GetVirtualMachineScaleSetNetworkInterfacePreparer(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName":       autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"virtualmachineIndex":        autorest.Encode("path", virtualmachineIndex),
		"virtualMachineScaleSetName": autorest.Encode("path", virtualMachineScaleSetName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// GetVirtualMachineScaleSetNetworkInterfaceSender sends the GetVirtualMachineScaleSetNetworkInterface request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) GetVirtualMachineScaleSetNetworkInterfaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// GetVirtualMachineScaleSetNetworkInterfaceResponder handles the response to the GetVirtualMachineScaleSetNetworkInterface request. The method always
// closes the http.Response Body.
func (client InterfacesClient) GetVirtualMachineScaleSetNetworkInterfaceResponder(resp *http.Response) (result Interface, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// List the List networkInterfaces opertion retrieves all the networkInterfaces in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client InterfacesClient) List(ctx context.Context, resourceGroupName string) (result InterfaceListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "List", resp, "Failure sending request")
		return
	}

	result.ilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "List", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListPreparer prepares the List request.
func (client InterfacesClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client InterfacesClient) ListResponder(resp *http.Response) (result InterfaceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client InterfacesClient) listNextResults(lastResults InterfaceListResult) (result InterfaceListResult, err error) {
	req, err := lastResults.interfaceListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client InterfacesClient) ListComplete(ctx context.Context, resourceGroupName string) (result InterfaceListResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName)
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListAll the List networkInterfaces opertion retrieves all the networkInterfaces in a subscription.
func (client InterfacesClient) ListAll(ctx context.Context) (result InterfaceListResultPage, err error) {
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.ilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.ilr, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListAllPreparer prepares the ListAll request.
func (client InterfacesClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkInterfaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client InterfacesClient) ListAllResponder(resp *http.Response) (result InterfaceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client InterfacesClient) listAllNextResults(lastResults InterfaceListResult) (result InterfaceListResult, err error) {
	req, err := lastResults.interfaceListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client InterfacesClient) ListAllComplete(ctx context.Context) (result InterfaceListResultIterator, err error) {
	result.page, err = client.ListAll(ctx)
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListVirtualMachineScaleSetNetworkInterfaces the list network interface operation retrieves information about all
// network interfaces in a virtual machine scale set.
//
// resourceGroupName is the name of the resource group. virtualMachineScaleSetName is the name of the virtual
// machine scale set.
func (client InterfacesClient) ListVirtualMachineScaleSetNetworkInterfaces(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (result InterfaceListResultPage, err error) {
	result.fn = client.listVirtualMachineScaleSetNetworkInterfacesNextResults
	req, err := client.ListVirtualMachineScaleSetNetworkInterfacesPreparer(ctx, resourceGroupName, virtualMachineScaleSetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListVirtualMachineScaleSetNetworkInterfaces", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListVirtualMachineScaleSetNetworkInterfacesSender(req)
	if err != nil {
		result.ilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListVirtualMachineScaleSetNetworkInterfaces", resp, "Failure sending request")
		return
	}

	result.ilr, err = client.ListVirtualMachineScaleSetNetworkInterfacesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListVirtualMachineScaleSetNetworkInterfaces", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListVirtualMachineScaleSetNetworkInterfacesPreparer prepares the ListVirtualMachineScaleSetNetworkInterfaces request.
func (client InterfacesClient) ListVirtualMachineScaleSetNetworkInterfacesPreparer(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"virtualMachineScaleSetName": autorest.Encode("path", virtualMachineScaleSetName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/networkInterfaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListVirtualMachineScaleSetNetworkInterfacesSender sends the ListVirtualMachineScaleSetNetworkInterfaces request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) ListVirtualMachineScaleSetNetworkInterfacesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListVirtualMachineScaleSetNetworkInterfacesResponder handles the response to the ListVirtualMachineScaleSetNetworkInterfaces request. The method always
// closes the http.Response Body.
func (client InterfacesClient) ListVirtualMachineScaleSetNetworkInterfacesResponder(resp *http.Response) (result InterfaceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listVirtualMachineScaleSetNetworkInterfacesNextResults retrieves the next set of results, if any.
func (client InterfacesClient) listVirtualMachineScaleSetNetworkInterfacesNextResults(lastResults InterfaceListResult) (result InterfaceListResult, err error) {
	req, err := lastResults.interfaceListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listVirtualMachineScaleSetNetworkInterfacesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListVirtualMachineScaleSetNetworkInterfacesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listVirtualMachineScaleSetNetworkInterfacesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListVirtualMachineScaleSetNetworkInterfacesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "listVirtualMachineScaleSetNetworkInterfacesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListVirtualMachineScaleSetNetworkInterfacesComplete enumerates all values, automatically crossing page boundaries as required.
func (client InterfacesClient) ListVirtualMachineScaleSetNetworkInterfacesComplete(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (result InterfaceListResultIterator, err error) {
	result.page, err = client.ListVirtualMachineScaleSetNetworkInterfaces(ctx, resourceGroupName, virtualMachineScaleSetName)
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListVirtualMachineScaleSetVMNetworkInterfaces the list network interface operation retrieves information about all
// network interfaces in a virtual machine from a virtual machine scale set.
//
// resourceGroupName is the name of the resource group. virtualMachineScaleSetName is the name of the virtual
// machine scale set. virtualmachineIndex is the virtual machine index.
func (client InterfacesClient) ListVirtualMachineScaleSetVMNetworkInterfaces(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string) (result InterfaceListResultPage, err error) {
	result.fn = client.listVirtualMachineScaleSetVMNetworkInterfacesNextResults
	req, err := client.ListVirtualMachineScaleSetVMNetworkInterfacesPreparer(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListVirtualMachineScaleSetVMNetworkInterfaces", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListVirtualMachineScaleSetVMNetworkInterfacesSender(req)
	if err != nil {
		result.ilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListVirtualMachineScaleSetVMNetworkInterfaces", resp, "Failure sending request")
		return
	}

	result.ilr, err = client.ListVirtualMachineScaleSetVMNetworkInterfacesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListVirtualMachineScaleSetVMNetworkInterfaces", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListVirtualMachineScaleSetVMNetworkInterfacesPreparer prepares the ListVirtualMachineScaleSetVMNetworkInterfaces request.
func (client InterfacesClient) ListVirtualMachineScaleSetVMNetworkInterfacesPreparer(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"virtualmachineIndex":        autorest.Encode("path", virtualmachineIndex),
		"virtualMachineScaleSetName": autorest.Encode("path", virtualMachineScaleSetName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListVirtualMachineScaleSetVMNetworkInterfacesSender sends the ListVirtualMachineScaleSetVMNetworkInterfaces request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) ListVirtualMachineScaleSetVMNetworkInterfacesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListVirtualMachineScaleSetVMNetworkInterfacesResponder handles the response to the ListVirtualMachineScaleSetVMNetworkInterfaces request. The method always
// closes the http.Response Body.
func (client InterfacesClient) ListVirtualMachineScaleSetVMNetworkInterfacesResponder(resp *http.Response) (result InterfaceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listVirtualMachineScaleSetVMNetworkInterfacesNextResults retrieves the next set of results, if any.
func (client InterfacesClient) listVirtualMachineScaleSetVMNetworkInterfacesNextResults(lastResults InterfaceListResult) (result InterfaceListResult, err error) {
	req, err := lastResults.interfaceListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listVirtualMachineScaleSetVMNetworkInterfacesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListVirtualMachineScaleSetVMNetworkInterfacesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listVirtualMachineScaleSetVMNetworkInterfacesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListVirtualMachineScaleSetVMNetworkInterfacesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "listVirtualMachineScaleSetVMNetworkInterfacesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/network/mgmt/2015-05-01-preview/network instead.
// ListVirtualMachineScaleSetVMNetworkInterfacesComplete enumerates all values, automatically crossing page boundaries as required.
func (client InterfacesClient) ListVirtualMachineScaleSetVMNetworkInterfacesComplete(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string) (result InterfaceListResultIterator, err error) {
	result.page, err = client.ListVirtualMachineScaleSetVMNetworkInterfaces(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex)
	return
}