// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"errors"
)

// DescribeNetworkPackagesRequest is the request struct for DescribeNetworkPackages
type DescribeNetworkPackagesRequest struct {
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
}

// Validate validates the DescribeNetworkPackagesRequest
func (s *DescribeNetworkPackagesRequest) Validate() error {
	if s.BizRegionId == nil || *s.BizRegionId == "" {
		return errors.New("BizRegionId is required")
	}
	return nil
}
