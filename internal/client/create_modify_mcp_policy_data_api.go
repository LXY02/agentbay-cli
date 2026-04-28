// Copyright 2025 AgentBay CLI Contributors
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// buildCreateModifyPolicyBody serializes the shared request into a formData body map.
func buildCreateModifyPolicyBody(request *CreateModifyMcpPolicyDataRequest) map[string]interface{} {
	marshalNested := func(v interface{}) string {
		b, _ := json.Marshal(v)
		return string(b)
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}
	if request.SandboxLifeCycle != nil {
		body["SandboxLifeCycle"] = marshalNested(map[string]interface{}{
			"Mode":              request.SandboxLifeCycle.Mode,
			"IdleTimeoutSwitch": request.SandboxLifeCycle.IdleTimeoutSwitch,
			"HibernateTimeout":  request.SandboxLifeCycle.HibernateTimeout,
			"DesktopMaxRuntime": request.SandboxLifeCycle.DesktopMaxRuntime,
			"UserIdleTimeout":   request.SandboxLifeCycle.UserIdleTimeout,
		})
	}
	if request.NetworkConfig != nil {
		body["NetworkConfig"] = marshalNested(map[string]interface{}{
			"Enabled": request.NetworkConfig.Enabled,
		})
	}
	if request.DisplayConfig != nil {
		body["DisplayConfig"] = marshalNested(map[string]interface{}{
			"DisplayMode": request.DisplayConfig.DisplayMode,
		})
	}
	if !dara.IsNil(request.Taskbar) {
		body["Taskbar"] = request.Taskbar
	}
	if !dara.IsNil(request.ScreenDisplayMode) {
		body["ScreenDisplayMode"] = request.ScreenDisplayMode
	}
	if !dara.IsNil(request.ClientControlMenu) {
		body["ClientControlMenu"] = request.ClientControlMenu
	}
	if !dara.IsNil(request.BusinessType) {
		body["BusinessType"] = request.BusinessType
	}
	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}
	if !dara.IsNil(request.DisconnectKeepSession) {
		body["DisconnectKeepSession"] = request.DisconnectKeepSession
	}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}
	if !dara.IsNil(request.InternetCommunicationProtocol) {
		body["InternetCommunicationProtocol"] = request.InternetCommunicationProtocol
	}
	if !dara.IsNil(request.ResolutionWidth) {
		body["ResolutionWidth"] = request.ResolutionWidth
	}
	if !dara.IsNil(request.ResolutionHeight) {
		body["ResolutionHeight"] = request.ResolutionHeight
	}
	if !dara.IsNil(request.RegionName) {
		body["RegionName"] = request.RegionName
	}
	return body
}

// =============== CreateMcpPolicyData ===============

func (client *Client) CreateMcpPolicyDataWithOptions(request *CreateModifyMcpPolicyDataRequest, runtime *dara.RuntimeOptions) (_result *CreateMcpPolicyDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := buildCreateModifyPolicyBody(request)

	req := &openapiutil.OpenApiRequest{
		Body:    openapiutil.ParseToMap(body),
		Headers: map[string]*string{"Accept": dara.String("application/json")},
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcpPolicyData"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &CreateMcpPolicyDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		reqID := ""
		if _body != nil {
			reqID = extractRequestIDFromResponse(_body)
		}
		return _result, &ErrWithRequestID{Err: _err, RequestID: reqID}
	}
	_result, _err = parseCreateMcpPolicyDataResponse(_body)
	return _result, _err
}

func (client *Client) CreateMcpPolicyData(request *CreateModifyMcpPolicyDataRequest) (_result *CreateMcpPolicyDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMcpPolicyDataResponse{}
	_body, _err := client.CreateMcpPolicyDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMcpPolicyDataWithContext(ctx context.Context, request *CreateModifyMcpPolicyDataRequest, runtime *dara.RuntimeOptions) (_result *CreateMcpPolicyDataResponse, _err error) {
	return client.CreateMcpPolicyDataWithOptions(request, runtime)
}

func parseCreateMcpPolicyDataResponse(res map[string]interface{}) (*CreateMcpPolicyDataResponse, error) {
	out := &CreateMcpPolicyDataResponse{}
	bodyStr := ""
	switch v := res["body"].(type) {
	case string:
		bodyStr = v
	case []byte:
		bodyStr = string(v)
	default:
		return nil, &ErrWithRequestID{Err: errors.New("missing or invalid body in response"), RequestID: extractRequestIDFromResponse(res)}
	}
	out.RawBody = bodyStr
	parsed := &CreateMcpPolicyDataResponseBody{}
	if bodyStr != "" {
		trimmed := strings.TrimSpace(bodyStr)
		if len(trimmed) > 0 && trimmed[0] == '<' {
			return nil, &ErrWithRequestID{Err: errors.New("unexpected XML response for CreateMcpPolicyData"), RequestID: extractRequestIDFromResponse(res)}
		}
		if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
			return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
		}
	}
	out.Body = parsed
	if h, ok := res["headers"].(map[string]*string); ok {
		out.Headers = h
	} else if h, ok := res["headers"].(map[string]interface{}); ok {
		out.Headers = make(map[string]*string)
		for k, v := range h {
			if s, ok := v.(string); ok {
				out.Headers[k] = dara.String(s)
			} else if p, ok := v.(*string); ok && p != nil {
				out.Headers[k] = p
			}
		}
	}
	if sc, ok := res["statusCode"].(int); ok {
		out.StatusCode = dara.Int32(int32(sc))
	}
	if sc, ok := res["statusCode"].(int32); ok {
		out.StatusCode = &sc
	}
	if out.StatusCode == nil && res["statusCode"] != nil {
		if n, err := strconv.Atoi(dara.ToString(res["statusCode"])); err == nil {
			out.StatusCode = dara.Int32(int32(n))
		}
	}
	return out, nil
}

// =============== ModifyMcpPolicyData ===============

func (client *Client) ModifyMcpPolicyDataWithOptions(request *CreateModifyMcpPolicyDataRequest, runtime *dara.RuntimeOptions) (_result *ModifyMcpPolicyDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := buildCreateModifyPolicyBody(request)

	req := &openapiutil.OpenApiRequest{
		Body:    openapiutil.ParseToMap(body),
		Headers: map[string]*string{"Accept": dara.String("application/json")},
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMcpPolicyData"),
		Version:     dara.String("2025-05-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &ModifyMcpPolicyDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		reqID := ""
		if _body != nil {
			reqID = extractRequestIDFromResponse(_body)
		}
		return _result, &ErrWithRequestID{Err: _err, RequestID: reqID}
	}
	_result, _err = parseModifyMcpPolicyDataResponse(_body)
	return _result, _err
}

func (client *Client) ModifyMcpPolicyDataSDK(request *CreateModifyMcpPolicyDataRequest) (_result *ModifyMcpPolicyDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMcpPolicyDataResponse{}
	_body, _err := client.ModifyMcpPolicyDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMcpPolicyDataWithContext(ctx context.Context, request *CreateModifyMcpPolicyDataRequest, runtime *dara.RuntimeOptions) (_result *ModifyMcpPolicyDataResponse, _err error) {
	return client.ModifyMcpPolicyDataWithOptions(request, runtime)
}

func parseModifyMcpPolicyDataResponse(res map[string]interface{}) (*ModifyMcpPolicyDataResponse, error) {
	out := &ModifyMcpPolicyDataResponse{}
	bodyStr := ""
	switch v := res["body"].(type) {
	case string:
		bodyStr = v
	case []byte:
		bodyStr = string(v)
	default:
		return nil, &ErrWithRequestID{Err: errors.New("missing or invalid body in response"), RequestID: extractRequestIDFromResponse(res)}
	}
	out.RawBody = bodyStr
	parsed := &ModifyMcpPolicyDataResponseBody{}
	if bodyStr != "" {
		trimmed := strings.TrimSpace(bodyStr)
		if len(trimmed) > 0 && trimmed[0] == '<' {
			return nil, &ErrWithRequestID{Err: errors.New("unexpected XML response for ModifyMcpPolicyData"), RequestID: extractRequestIDFromResponse(res)}
		}
		if err := json.Unmarshal([]byte(bodyStr), parsed); err != nil {
			return nil, &ErrWithRequestID{Err: err, RequestID: extractRequestIDFromResponse(res)}
		}
	}
	out.Body = parsed
	if h, ok := res["headers"].(map[string]*string); ok {
		out.Headers = h
	} else if h, ok := res["headers"].(map[string]interface{}); ok {
		out.Headers = make(map[string]*string)
		for k, v := range h {
			if s, ok := v.(string); ok {
				out.Headers[k] = dara.String(s)
			} else if p, ok := v.(*string); ok && p != nil {
				out.Headers[k] = p
			}
		}
	}
	if sc, ok := res["statusCode"].(int); ok {
		out.StatusCode = dara.Int32(int32(sc))
	}
	if sc, ok := res["statusCode"].(int32); ok {
		out.StatusCode = &sc
	}
	if out.StatusCode == nil && res["statusCode"] != nil {
		if n, err := strconv.Atoi(dara.ToString(res["statusCode"])); err == nil {
			out.StatusCode = dara.Int32(int32(n))
		}
	}
	return out, nil
}
