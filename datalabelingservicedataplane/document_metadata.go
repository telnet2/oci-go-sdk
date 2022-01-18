// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// DocumentMetadata Collection of metadata related to document record.
type DocumentMetadata struct {
}

func (m DocumentMetadata) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DocumentMetadata) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDocumentMetadata DocumentMetadata
	s := struct {
		DiscriminatorParam string `json:"recordType"`
		MarshalTypeDocumentMetadata
	}{
		"DOCUMENT_METADATA",
		(MarshalTypeDocumentMetadata)(m),
	}

	return json.Marshal(&s)
}
