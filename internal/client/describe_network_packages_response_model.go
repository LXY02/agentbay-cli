// This file is auto-generated, don't edit it. Thanks.
package client

// DescribeNetworkPackagesResponseBodyDataItem represents a single network package item
type DescribeNetworkPackagesResponseBodyDataItem struct {
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	EipAddresses     *string `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty"`
}

// GetNetworkPackageId returns the NetworkPackageId value or empty string if nil
func (s *DescribeNetworkPackagesResponseBodyDataItem) GetNetworkPackageId() string {
	if s == nil || s.NetworkPackageId == nil {
		return ""
	}
	return *s.NetworkPackageId
}

// GetEipAddresses returns the EipAddresses value or empty string if nil
func (s *DescribeNetworkPackagesResponseBodyDataItem) GetEipAddresses() string {
	if s == nil || s.EipAddresses == nil {
		return ""
	}
	return *s.EipAddresses
}

// DescribeNetworkPackagesResponseBodyData represents the Data field in the response
type DescribeNetworkPackagesResponseBodyData struct {
	Items []*DescribeNetworkPackagesResponseBodyDataItem `json:"Items,omitempty" xml:"Items,omitempty"`
}

// DescribeNetworkPackagesResponseBody is the response body struct for DescribeNetworkPackages
type DescribeNetworkPackagesResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeNetworkPackagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

// DescribeNetworkPackagesResponse is the response struct for DescribeNetworkPackages
type DescribeNetworkPackagesResponse struct {
	Headers    map[string]*string
	StatusCode *int32
	Body       *DescribeNetworkPackagesResponseBody
}

// GetCode returns the Code value or empty string if nil
func (s *DescribeNetworkPackagesResponseBody) GetCode() string {
	if s == nil || s.Code == nil {
		return ""
	}
	return *s.Code
}

// GetRequestId returns the RequestId pointer
func (s *DescribeNetworkPackagesResponseBody) GetRequestId() *string {
	if s == nil {
		return nil
	}
	return s.RequestId
}

// GetSuccess returns the Success pointer
func (s *DescribeNetworkPackagesResponseBody) GetSuccess() *bool {
	if s == nil {
		return nil
	}
	return s.Success
}

// GetMessage returns the Message pointer
func (s *DescribeNetworkPackagesResponseBody) GetMessage() *string {
	if s == nil {
		return nil
	}
	return s.Message
}

// GetData returns the Data field
func (s *DescribeNetworkPackagesResponseBody) GetData() *DescribeNetworkPackagesResponseBodyData {
	if s == nil {
		return nil
	}
	return s.Data
}
