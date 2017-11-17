// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"net/http"
)

// ListShapesRequest wrapper for the ListShapes operation
type ListShapesRequest struct {

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancer shapes to list.
	CompartmentID *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int64 `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`
}

func (request ListShapesRequest) String() string {
	return common.PointerString(request)
}

// ListShapesResponse wrapper for the ListShapes operation
type ListShapesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []LoadBalancerShape instance
	Items []LoadBalancerShape `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListShapesResponse) String() string {
	return common.PointerString(response)
}