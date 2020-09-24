package sql

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RecommendedElasticPoolsClient is the the Azure SQL Database management API provides a RESTful set of web services
// that interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve,
// update, and delete databases.
type RecommendedElasticPoolsClient struct {
	BaseClient
}

// NewRecommendedElasticPoolsClient creates an instance of the RecommendedElasticPoolsClient client.
func NewRecommendedElasticPoolsClient(subscriptionID string) RecommendedElasticPoolsClient {
	return NewRecommendedElasticPoolsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecommendedElasticPoolsClientWithBaseURI creates an instance of the RecommendedElasticPoolsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewRecommendedElasticPoolsClientWithBaseURI(baseURI string, subscriptionID string) RecommendedElasticPoolsClient {
	return RecommendedElasticPoolsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a recommended elastic pool.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// recommendedElasticPoolName - the name of the recommended elastic pool to be retrieved.
func (client RecommendedElasticPoolsClient) Get(ctx context.Context, resourceGroupName string, serverName string, recommendedElasticPoolName string) (result RecommendedElasticPool, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecommendedElasticPoolsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, recommendedElasticPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RecommendedElasticPoolsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, recommendedElasticPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recommendedElasticPoolName": autorest.Encode("path", recommendedElasticPoolName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"serverName":                 autorest.Encode("path", serverName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recommendedElasticPools/{recommendedElasticPoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendedElasticPoolsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RecommendedElasticPoolsClient) GetResponder(resp *http.Response) (result RecommendedElasticPool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer returns recommended elastic pools.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
func (client RecommendedElasticPoolsClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result RecommendedElasticPoolListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecommendedElasticPoolsClient.ListByServer")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "ListByServer", resp, "Failure responding to request")
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client RecommendedElasticPoolsClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recommendedElasticPools", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendedElasticPoolsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client RecommendedElasticPoolsClient) ListByServerResponder(resp *http.Response) (result RecommendedElasticPoolListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListMetrics returns recommended elastic pool metrics.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// recommendedElasticPoolName - the name of the recommended elastic pool to be retrieved.
func (client RecommendedElasticPoolsClient) ListMetrics(ctx context.Context, resourceGroupName string, serverName string, recommendedElasticPoolName string) (result RecommendedElasticPoolListMetricsResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecommendedElasticPoolsClient.ListMetrics")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListMetricsPreparer(ctx, resourceGroupName, serverName, recommendedElasticPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "ListMetrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListMetricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "ListMetrics", resp, "Failure sending request")
		return
	}

	result, err = client.ListMetricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.RecommendedElasticPoolsClient", "ListMetrics", resp, "Failure responding to request")
	}

	return
}

// ListMetricsPreparer prepares the ListMetrics request.
func (client RecommendedElasticPoolsClient) ListMetricsPreparer(ctx context.Context, resourceGroupName string, serverName string, recommendedElasticPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recommendedElasticPoolName": autorest.Encode("path", recommendedElasticPoolName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"serverName":                 autorest.Encode("path", serverName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recommendedElasticPools/{recommendedElasticPoolName}/metrics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListMetricsSender sends the ListMetrics request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendedElasticPoolsClient) ListMetricsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListMetricsResponder handles the response to the ListMetrics request. The method always
// closes the http.Response Body.
func (client RecommendedElasticPoolsClient) ListMetricsResponder(resp *http.Response) (result RecommendedElasticPoolListMetricsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}