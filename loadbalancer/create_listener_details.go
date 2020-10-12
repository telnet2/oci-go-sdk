// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/v27/common"
)

// CreateListenerDetails The configuration details for adding a listener to a backend set.
// For more information on listener configuration, see
// Managing Load Balancer Listeners (https://docs.cloud.oracle.com/Content/Balance/Tasks/managinglisteners.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateListenerDetails struct {

	// The name of the associated backend set.
	// Example: `example_backend_set`
	DefaultBackendSetName *string `mandatory:"true" json:"defaultBackendSetName"`

	// The communication port for the listener.
	// Example: `80`
	Port *int `mandatory:"true" json:"port"`

	// The protocol on which the listener accepts connection requests.
	// To get a list of valid protocols, use the ListProtocols
	// operation.
	// Example: `HTTP`
	Protocol *string `mandatory:"true" json:"protocol"`

	// A friendly name for the listener. It must be unique and it cannot be changed.
	// Avoid entering confidential information.
	// Example: `example_listener`
	Name *string `mandatory:"true" json:"name"`

	// An array of hostname resource names.
	HostnameNames []string `mandatory:"false" json:"hostnameNames"`

	// The name of the set of path-based routing rules, PathRouteSet,
	// applied to this listener's traffic.
	// Example: `example_path_route_set`
	PathRouteSetName *string `mandatory:"false" json:"pathRouteSetName"`

	SslConfiguration *SslConfigurationDetails `mandatory:"false" json:"sslConfiguration"`

	ConnectionConfiguration *ConnectionConfiguration `mandatory:"false" json:"connectionConfiguration"`

	// The names of the RuleSet to apply to the listener.
	// Example: ["example_rule_set"]
	RuleSetNames []string `mandatory:"false" json:"ruleSetNames"`
}

func (m CreateListenerDetails) String() string {
	return common.PointerString(m)
}
