// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"strings"
)

// SecurityAssessment A security assessment that provides an overall insight into your database security posture.
// The security assessment results are based on the analysis of your database configurations,
// user accounts, and security controls. For more information, see Security Assessment Overview (https://docs.oracle.com/en/cloud/paas/data-safe/udscs/security-assessment-overview.html).
type SecurityAssessment struct {

	// The OCID of the security assessment.
	Id *string `mandatory:"true" json:"id"`

	// The date and time when the security assessment was created. Conforms to the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time when the security assessment was last updated. Conforms to the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the compartment that contains the security assessment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the security assessment.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Array of database target OCIDs.
	TargetIds []string `mandatory:"true" json:"targetIds"`

	// The current state of the security assessment.
	LifecycleState SecurityAssessmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The type of this security assessment. The possible types are:
	// LATEST: The most up-to-date assessment that is running automatically for a target. It is system generated.
	// SAVED: A saved security assessment. LATEST assessments are always saved in order to maintain the history of runs. A SAVED assessment is also generated by a 'refresh' action (triggered by the user).
	// SAVE_SCHEDULE: The schedule for periodic saves of LATEST assessments.
	// COMPARTMENT: An automatically managed assessment type that stores all details of targets in one compartment.
	//  This type keeps an up-to-date assessment of all database risks in one compartment. It is automatically updated when
	//  the latest assessment or refresh action is executed. It is also automatically updated when a target is deleted or move to a different compartment.
	Type SecurityAssessmentTypeEnum `mandatory:"true" json:"type"`

	// List containing maps as values.
	// Example: `{"Operations": [ {"CostCenter": "42"} ] }`
	IgnoredTargets []interface{} `mandatory:"false" json:"ignoredTargets"`

	// List containing maps as values.
	// Example: `{"Operations": [ {"CostCenter": "42"} ] }`
	IgnoredAssessmentIds []interface{} `mandatory:"false" json:"ignoredAssessmentIds"`

	// The version of the target database.
	TargetVersion *string `mandatory:"false" json:"targetVersion"`

	// Indicates whether or not the security assessment is set as a baseline. This is applicable only for saved security assessments.
	IsBaseline *bool `mandatory:"false" json:"isBaseline"`

	// Indicates if the assessment has deviated from the baseline.
	IsDeviatedFromBaseline *bool `mandatory:"false" json:"isDeviatedFromBaseline"`

	// The OCID of the baseline against which the latest security assessment was compared.
	LastComparedBaselineId *string `mandatory:"false" json:"lastComparedBaselineId"`

	// Details about the current state of the security assessment.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID of the security assessment that is responsible for creating this scheduled save assessment.
	ScheduleSecurityAssessmentId *string `mandatory:"false" json:"scheduleSecurityAssessmentId"`

	// Indicates whether the security assessment was created by system or by a user.
	TriggeredBy SecurityAssessmentTriggeredByEnum `mandatory:"false" json:"triggeredBy,omitempty"`

	// The description of the security assessment.
	Description *string `mandatory:"false" json:"description"`

	// Schedule to save the assessment periodically in the specified format:
	// <version-string>;<version-specific-schedule>
	// Allowed version strings - "v1"
	// v1's version specific schedule -<ss> <mm> <hh> <day-of-week> <day-of-month>
	// Each of the above fields potentially introduce constraints. A workrequest is created only
	// when clock time satisfies all the constraints. Constraints introduced:
	// 1. seconds = <ss> (So, the allowed range for <ss> is [0, 59])
	// 2. minutes = <mm> (So, the allowed range for <mm> is [0, 59])
	// 3. hours = <hh> (So, the allowed range for <hh> is [0, 23])
	// <day-of-week> can be either '*' (without quotes or a number between 1(Monday) and 7(Sunday))
	// 4. No constraint introduced when it is '*'. When not, day of week must equal the given value
	// <day-of-month> can be either '*' (without quotes or a number between 1 and 28)
	// 5. No constraint introduced when it is '*'. When not, day of month must equal the given value
	Schedule *string `mandatory:"false" json:"schedule"`

	// The summary of findings for the security assessment
	Link *string `mandatory:"false" json:"link"`

	Statistics *SecurityAssessmentStatistics `mandatory:"false" json:"statistics"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SecurityAssessment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityAssessment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSecurityAssessmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSecurityAssessmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSecurityAssessmentTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetSecurityAssessmentTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSecurityAssessmentTriggeredByEnum(string(m.TriggeredBy)); !ok && m.TriggeredBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TriggeredBy: %s. Supported values are: %s.", m.TriggeredBy, strings.Join(GetSecurityAssessmentTriggeredByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecurityAssessmentTriggeredByEnum Enum with underlying type: string
type SecurityAssessmentTriggeredByEnum string

// Set of constants representing the allowable values for SecurityAssessmentTriggeredByEnum
const (
	SecurityAssessmentTriggeredByUser   SecurityAssessmentTriggeredByEnum = "USER"
	SecurityAssessmentTriggeredBySystem SecurityAssessmentTriggeredByEnum = "SYSTEM"
)

var mappingSecurityAssessmentTriggeredByEnum = map[string]SecurityAssessmentTriggeredByEnum{
	"USER":   SecurityAssessmentTriggeredByUser,
	"SYSTEM": SecurityAssessmentTriggeredBySystem,
}

var mappingSecurityAssessmentTriggeredByEnumLowerCase = map[string]SecurityAssessmentTriggeredByEnum{
	"user":   SecurityAssessmentTriggeredByUser,
	"system": SecurityAssessmentTriggeredBySystem,
}

// GetSecurityAssessmentTriggeredByEnumValues Enumerates the set of values for SecurityAssessmentTriggeredByEnum
func GetSecurityAssessmentTriggeredByEnumValues() []SecurityAssessmentTriggeredByEnum {
	values := make([]SecurityAssessmentTriggeredByEnum, 0)
	for _, v := range mappingSecurityAssessmentTriggeredByEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityAssessmentTriggeredByEnumStringValues Enumerates the set of values in String for SecurityAssessmentTriggeredByEnum
func GetSecurityAssessmentTriggeredByEnumStringValues() []string {
	return []string{
		"USER",
		"SYSTEM",
	}
}

// GetMappingSecurityAssessmentTriggeredByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityAssessmentTriggeredByEnum(val string) (SecurityAssessmentTriggeredByEnum, bool) {
	enum, ok := mappingSecurityAssessmentTriggeredByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecurityAssessmentTypeEnum Enum with underlying type: string
type SecurityAssessmentTypeEnum string

// Set of constants representing the allowable values for SecurityAssessmentTypeEnum
const (
	SecurityAssessmentTypeLatest       SecurityAssessmentTypeEnum = "LATEST"
	SecurityAssessmentTypeSaved        SecurityAssessmentTypeEnum = "SAVED"
	SecurityAssessmentTypeSaveSchedule SecurityAssessmentTypeEnum = "SAVE_SCHEDULE"
	SecurityAssessmentTypeCompartment  SecurityAssessmentTypeEnum = "COMPARTMENT"
)

var mappingSecurityAssessmentTypeEnum = map[string]SecurityAssessmentTypeEnum{
	"LATEST":        SecurityAssessmentTypeLatest,
	"SAVED":         SecurityAssessmentTypeSaved,
	"SAVE_SCHEDULE": SecurityAssessmentTypeSaveSchedule,
	"COMPARTMENT":   SecurityAssessmentTypeCompartment,
}

var mappingSecurityAssessmentTypeEnumLowerCase = map[string]SecurityAssessmentTypeEnum{
	"latest":        SecurityAssessmentTypeLatest,
	"saved":         SecurityAssessmentTypeSaved,
	"save_schedule": SecurityAssessmentTypeSaveSchedule,
	"compartment":   SecurityAssessmentTypeCompartment,
}

// GetSecurityAssessmentTypeEnumValues Enumerates the set of values for SecurityAssessmentTypeEnum
func GetSecurityAssessmentTypeEnumValues() []SecurityAssessmentTypeEnum {
	values := make([]SecurityAssessmentTypeEnum, 0)
	for _, v := range mappingSecurityAssessmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityAssessmentTypeEnumStringValues Enumerates the set of values in String for SecurityAssessmentTypeEnum
func GetSecurityAssessmentTypeEnumStringValues() []string {
	return []string{
		"LATEST",
		"SAVED",
		"SAVE_SCHEDULE",
		"COMPARTMENT",
	}
}

// GetMappingSecurityAssessmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityAssessmentTypeEnum(val string) (SecurityAssessmentTypeEnum, bool) {
	enum, ok := mappingSecurityAssessmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
