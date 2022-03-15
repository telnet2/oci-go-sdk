// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"strings"
)

// CreateBootVolumeBackupDetails The representation of CreateBootVolumeBackupDetails
type CreateBootVolumeBackupDetails struct {

	// The OCID of the boot volume that needs to be backed up.
	BootVolumeId *string `mandatory:"true" json:"bootVolumeId"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The type of backup to create. If omitted, defaults to incremental.
	Type CreateBootVolumeBackupDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m CreateBootVolumeBackupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBootVolumeBackupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateBootVolumeBackupDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCreateBootVolumeBackupDetailsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateBootVolumeBackupDetailsTypeEnum Enum with underlying type: string
type CreateBootVolumeBackupDetailsTypeEnum string

// Set of constants representing the allowable values for CreateBootVolumeBackupDetailsTypeEnum
const (
	CreateBootVolumeBackupDetailsTypeFull        CreateBootVolumeBackupDetailsTypeEnum = "FULL"
	CreateBootVolumeBackupDetailsTypeIncremental CreateBootVolumeBackupDetailsTypeEnum = "INCREMENTAL"
)

var mappingCreateBootVolumeBackupDetailsTypeEnum = map[string]CreateBootVolumeBackupDetailsTypeEnum{
	"FULL":        CreateBootVolumeBackupDetailsTypeFull,
	"INCREMENTAL": CreateBootVolumeBackupDetailsTypeIncremental,
}

var mappingCreateBootVolumeBackupDetailsTypeEnumLowerCase = map[string]CreateBootVolumeBackupDetailsTypeEnum{
	"full":        CreateBootVolumeBackupDetailsTypeFull,
	"incremental": CreateBootVolumeBackupDetailsTypeIncremental,
}

// GetCreateBootVolumeBackupDetailsTypeEnumValues Enumerates the set of values for CreateBootVolumeBackupDetailsTypeEnum
func GetCreateBootVolumeBackupDetailsTypeEnumValues() []CreateBootVolumeBackupDetailsTypeEnum {
	values := make([]CreateBootVolumeBackupDetailsTypeEnum, 0)
	for _, v := range mappingCreateBootVolumeBackupDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateBootVolumeBackupDetailsTypeEnumStringValues Enumerates the set of values in String for CreateBootVolumeBackupDetailsTypeEnum
func GetCreateBootVolumeBackupDetailsTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"INCREMENTAL",
	}
}

// GetMappingCreateBootVolumeBackupDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateBootVolumeBackupDetailsTypeEnum(val string) (CreateBootVolumeBackupDetailsTypeEnum, bool) {
	enum, ok := mappingCreateBootVolumeBackupDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
