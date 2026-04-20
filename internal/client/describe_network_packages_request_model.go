// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"errors"
)

// DescribeNetworkPackagesRequest is the request struct for DescribeNetworkPackages
type DescribeNetworkPackagesRequest struct {
	UserAliUid  *string `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
}

// Validate validates the DescribeNetworkPackagesRequest
func (s *DescribeNetworkPackagesRequest) Validate() error {
	if s.UserAliUid == nil || *s.UserAliUid == "" {
		return errors.New("UserAliUid is required")
	}
	if s.BizRegionId == nil || *s.BizRegionId == "" {
		return errors.New("BizRegionId is required")
	}
	return nil
}
