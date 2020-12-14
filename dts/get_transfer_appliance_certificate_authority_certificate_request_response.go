// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dts

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"net/http"
)

// GetTransferApplianceCertificateAuthorityCertificateRequest wrapper for the GetTransferApplianceCertificateAuthorityCertificate operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/GetTransferApplianceCertificateAuthorityCertificate.go.html to see an example of how to use GetTransferApplianceCertificateAuthorityCertificateRequest.
type GetTransferApplianceCertificateAuthorityCertificateRequest struct {

	// ID of the Transfer Job
	Id *string `mandatory:"true" contributesTo:"path" name:"id"`

	// Label of the Transfer Appliance
	TransferApplianceLabel *string `mandatory:"true" contributesTo:"path" name:"transferApplianceLabel"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetTransferApplianceCertificateAuthorityCertificateRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetTransferApplianceCertificateAuthorityCertificateRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetTransferApplianceCertificateAuthorityCertificateRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetTransferApplianceCertificateAuthorityCertificateResponse wrapper for the GetTransferApplianceCertificateAuthorityCertificate operation
type GetTransferApplianceCertificateAuthorityCertificateResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The TransferApplianceCertificate instance
	TransferApplianceCertificate `presentIn:"body"`

	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	Etag *string `presentIn:"header" name:"etag"`
}

func (response GetTransferApplianceCertificateAuthorityCertificateResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetTransferApplianceCertificateAuthorityCertificateResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
