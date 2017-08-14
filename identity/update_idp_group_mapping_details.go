// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity


type UpdateIdpGroupMappingDetails struct {

    // The idp group name.
    IdpGroupName string `json:"idpGroupName,omitempty"`

    // The OCID of the group.
    GroupId string `json:"groupId,omitempty"`
}