// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import "net/http"

// GetGroupRequest wrapper for the GetGroup operation
type GetGroupRequest struct {

	// The OCID of the group.
	GroupID string `mandatory:"true" contributesTo:"path" name:"groupId"`
}

// GetGroupResponse wrapper for the GetGroup operation
type GetGroupResponse struct {

	// The underlying http response
	RawResponse http.Response

	// The Group instance
	Group `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag string `presentIn:"header" name:"etag"`
}
