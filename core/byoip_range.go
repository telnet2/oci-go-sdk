// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// ByoipRange A ByoipRange, is an IP address prefix that the user owns and wishes to import into OCI.
type ByoipRange struct {

	// The address range the user is on-boarding.
	CidrBlock *string `mandatory:"true" json:"cidrBlock"`

	// The OCID of the compartment containing the Byoip Range.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Oracle ID (OCID) of the Byoip Range.
	Id *string `mandatory:"true" json:"id"`

	// The Byoip Range's current state.
	LifecycleState ByoipRangeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the Byoip Range was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// This is an internally generated ASCII string that the user will then use as part of the validation process. Specifically, they will need to add the token string generated by the service to their Internet Registry record.
	ValidationToken *string `mandatory:"true" json:"validationToken"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The Byoip Range's current substate.
	LifecycleDetails ByoipRangeLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// The date and time the Byoip Range was validated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeValidated *common.SDKTime `mandatory:"false" json:"timeValidated"`

	// The date and time the Byoip Range was advertised, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeAdvertised *common.SDKTime `mandatory:"false" json:"timeAdvertised"`

	// The date and time the Byoip Range was withdrawn, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeWithdrawn *common.SDKTime `mandatory:"false" json:"timeWithdrawn"`
}

func (m ByoipRange) String() string {
	return common.PointerString(m)
}

// ByoipRangeLifecycleDetailsEnum Enum with underlying type: string
type ByoipRangeLifecycleDetailsEnum string

// Set of constants representing the allowable values for ByoipRangeLifecycleDetailsEnum
const (
	ByoipRangeLifecycleDetailsCreating    ByoipRangeLifecycleDetailsEnum = "CREATING"
	ByoipRangeLifecycleDetailsValidating  ByoipRangeLifecycleDetailsEnum = "VALIDATING"
	ByoipRangeLifecycleDetailsProvisioned ByoipRangeLifecycleDetailsEnum = "PROVISIONED"
	ByoipRangeLifecycleDetailsActive      ByoipRangeLifecycleDetailsEnum = "ACTIVE"
	ByoipRangeLifecycleDetailsFailed      ByoipRangeLifecycleDetailsEnum = "FAILED"
	ByoipRangeLifecycleDetailsDeleting    ByoipRangeLifecycleDetailsEnum = "DELETING"
	ByoipRangeLifecycleDetailsDeleted     ByoipRangeLifecycleDetailsEnum = "DELETED"
)

var mappingByoipRangeLifecycleDetails = map[string]ByoipRangeLifecycleDetailsEnum{
	"CREATING":    ByoipRangeLifecycleDetailsCreating,
	"VALIDATING":  ByoipRangeLifecycleDetailsValidating,
	"PROVISIONED": ByoipRangeLifecycleDetailsProvisioned,
	"ACTIVE":      ByoipRangeLifecycleDetailsActive,
	"FAILED":      ByoipRangeLifecycleDetailsFailed,
	"DELETING":    ByoipRangeLifecycleDetailsDeleting,
	"DELETED":     ByoipRangeLifecycleDetailsDeleted,
}

// GetByoipRangeLifecycleDetailsEnumValues Enumerates the set of values for ByoipRangeLifecycleDetailsEnum
func GetByoipRangeLifecycleDetailsEnumValues() []ByoipRangeLifecycleDetailsEnum {
	values := make([]ByoipRangeLifecycleDetailsEnum, 0)
	for _, v := range mappingByoipRangeLifecycleDetails {
		values = append(values, v)
	}
	return values
}

// ByoipRangeLifecycleStateEnum Enum with underlying type: string
type ByoipRangeLifecycleStateEnum string

// Set of constants representing the allowable values for ByoipRangeLifecycleStateEnum
const (
	ByoipRangeLifecycleStateInactive ByoipRangeLifecycleStateEnum = "INACTIVE"
	ByoipRangeLifecycleStateUpdating ByoipRangeLifecycleStateEnum = "UPDATING"
	ByoipRangeLifecycleStateActive   ByoipRangeLifecycleStateEnum = "ACTIVE"
	ByoipRangeLifecycleStateDeleting ByoipRangeLifecycleStateEnum = "DELETING"
	ByoipRangeLifecycleStateDeleted  ByoipRangeLifecycleStateEnum = "DELETED"
)

var mappingByoipRangeLifecycleState = map[string]ByoipRangeLifecycleStateEnum{
	"INACTIVE": ByoipRangeLifecycleStateInactive,
	"UPDATING": ByoipRangeLifecycleStateUpdating,
	"ACTIVE":   ByoipRangeLifecycleStateActive,
	"DELETING": ByoipRangeLifecycleStateDeleting,
	"DELETED":  ByoipRangeLifecycleStateDeleted,
}

// GetByoipRangeLifecycleStateEnumValues Enumerates the set of values for ByoipRangeLifecycleStateEnum
func GetByoipRangeLifecycleStateEnumValues() []ByoipRangeLifecycleStateEnum {
	values := make([]ByoipRangeLifecycleStateEnum, 0)
	for _, v := range mappingByoipRangeLifecycleState {
		values = append(values, v)
	}
	return values
}
