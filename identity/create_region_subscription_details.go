// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

type CreateRegionSubscriptionDetails struct {

	// The regions's key.
	// Allowed values are:
	// - `PHX`
	// - `IAD`
	// - `FRA`
	// Example: `PHX`
	RegionKey *string `mandatory:"true" json:"regionKey,omitempty"`
}

func (model CreateRegionSubscriptionDetails) String() string {
	return common.PointerString(model)
}
