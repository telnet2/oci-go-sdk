// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// AnnouncementUserStatusDetails An announcement's status regarding whether it has been acknowledged by a user.
type AnnouncementUserStatusDetails struct {

	// The OCID of the announcement that this status is associated with.
	UserStatusAnnouncementId *string `mandatory:"true" json:"userStatusAnnouncementId"`

	// The OCID of the user that this status is associated with.
	UserId *string `mandatory:"true" json:"userId"`

	// The date and time the announcement was acknowledged, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-01-01T17:43:01.389+0000`
	TimeAcknowledged *common.SDKTime `mandatory:"false" json:"timeAcknowledged"`
}

func (m AnnouncementUserStatusDetails) String() string {
	return common.PointerString(m)
}
