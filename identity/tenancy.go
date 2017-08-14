// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

//Tenancy The root compartment that contains all of your organization's compartments and other\nOracle Bare Metal Cloud Services cloud resources. When you sign up for Oracle Bare Metal Cloud Services,\nOracle creates a tenancy for your company, which is a secure and isolated partition\nwhere you can create, organize, and administer your cloud resources.\n\nTo use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,\ntalk to an administrator. If you're an administrator who needs to write policies to give users access,\nsee [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).\n

type Tenancy struct {

    // The OCID of the tenancy.
    Id string `json:"id,omitempty"`

    // The name of the tenancy.
    Name string `json:"name,omitempty"`

    // The description of the tenancy.
    Description string `json:"description,omitempty"`

    // The region key for the tenancy's home region.\n\nAllowed values are:\n- `IAD`\n- `PHX`\n
    HomeRegionKey string `json:"homeRegionKey,omitempty"`
}