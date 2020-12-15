// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// CreateApplianceExportJobDetails The representation of CreateApplianceExportJobDetails
type CreateApplianceExportJobDetails struct {
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	BucketName *string `mandatory:"true" json:"bucketName"`

	DisplayName *string `mandatory:"true" json:"displayName"`

	CustomerShippingAddress *ShippingAddress `mandatory:"true" json:"customerShippingAddress"`

	// List of objects with names matching this prefix would be part of this export job.
	Prefix *string `mandatory:"false" json:"prefix"`

	// Object names returned by a list query must be greater or equal to this parameter.
	RangeStart *string `mandatory:"false" json:"rangeStart"`

	// Object names returned by a list query must be strictly less than this parameter.
	RangeEnd *string `mandatory:"false" json:"rangeEnd"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateApplianceExportJobDetails) String() string {
	return common.PointerString(m)
}
