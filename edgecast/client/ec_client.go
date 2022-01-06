package client

/*
	This file contains the concrete client implementation for the EC SDK
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/eccollections"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecjson"
)

const (
	defaultHeaderAccept      string = "application/json"
	defaultHeaderContentType string = "application/json"
)

// ECClient -
type ECClient struct {
	reqBuilder requestBuilder
	reqSender  requestSender
	config     ClientConfig
}

// NewECClient creates a default instance of ECClient using the provided
// configuration
func NewECClient(config ClientConfig) ECClient {
	return ECClient{
		reqBuilder: newECRequestBuilder(config),
		reqSender:  newECRequestSender(config),
		config:     config,
	}
}

// Do invokes an HTTP request with the given parameters
func (c ECClient) Do(params DoParams) (interface{}, error) {
	req, err := c.reqBuilder.buildRequest(buildRequestParams{
		method:      params.Method,
		path:        params.Path,
		rawBody:     params.Body,
		queryParams: params.QueryParams,
		pathParams:  params.PathParams,
		userAgent:   c.config.UserAgent,
	})

	if err != nil {
		return nil, fmt.Errorf("ECClient.Do: %v", err)
	}

	c.config.Logger.Debug(
		"[REQUEST-URI]:[%s] %s\n",
		req.method,
		req.url.String())
	c.config.Logger.Debug("[REQUEST-BODY]:%v\n", req.rawBody)
	c.config.Logger.Debug("[REQUEST-HEADERS]:%+v\n", req.headers)

	resp, err := c.reqSender.sendRequest(*req)
	if err != nil {
		return nil, fmt.Errorf("ECClient.Do: %v", err)
	}
	bodyAsString, _ := ecjson.ConvertToJSONString(resp.data, true)
	c.config.Logger.Debug("[RESPONSE-BODY]:%s\n", bodyAsString)
	return resp.data, nil
}

// ecRequestBuilder builds requests to be sent to the Edgecast API
type ecRequestBuilder struct {
	baseAPIURL   url.URL
	authProvider *auth.AuthorizationProvider
	userAgent    string
}

// newECRequestBuilder creates a default instance of ecRequestBuilder using the
// provided configuration
func newECRequestBuilder(config ClientConfig) ecRequestBuilder {
	return ecRequestBuilder{
		baseAPIURL:   config.BaseAPIURL,
		authProvider: &config.AuthProvider,
		userAgent:    config.UserAgent,
	}
}

// buildRequest creates a new Request for the Edgecast API with query params,
// adding appropriate headers
func (eb ecRequestBuilder) buildRequest(
	params buildRequestParams,
) (*request, error) {
	relativeURL, err := url.Parse(params.path)
	if err != nil {
		return nil,
			fmt.Errorf("ecRequestBuilder.buildRequest: url.Parse: %v", err)
	}

	req := request{
		method:  params.method,
		url:     eb.baseAPIURL.ResolveReference(relativeURL),
		headers: make(map[string]string),
	}

	req.setPathParams(params.pathParams)
	req.setQueryParams(params.queryParams)

	req.headers["User-Agent"] = params.userAgent
	req.headers["Accept"] = defaultHeaderAccept

	if params.rawBody != nil {
		err := req.setBody(params.rawBody)
		if err != nil {
			return nil, fmt.Errorf(
				"ecRequestBuilder.buildRequest: %v", err)
		}
	}

	if eb.authProvider != nil {
		err := req.setAuthorization(*eb.authProvider)
		if err != nil {
			return nil, fmt.Errorf(
				"ecRequestBuilder.buildRequest: %v", err)
		}
	}

	return &req, nil
}

func (req *request) setPathParams(params map[string]string) {
	// Apply path parameters
	// e.g.
	// path = "/customers/{id}/policies/{id}""
	// params = { id: 1, policy: 99 }
	// result = "customers/1/policies/99"
	for k, v := range params {
		req.url.Path = strings.Replace(
			req.url.Path,
			"{"+k+"}",
			fmt.Sprintf("%v", v),
			-1)
	}
}

func (req *request) setQueryParams(queryParams map[string]string) {
	// Adding Query Params
	query := req.url.Query()
	for k, v := range queryParams {
		query.Add(k, v)
	}
	// Encode the parameters and set the URL
	req.url.RawQuery = query.Encode()
}

func (req *request) setBody(rawBody interface{}) error {
	switch b := rawBody.(type) {
	case string:
		req.rawBody = []byte(b)
		req.headers["Content-Type"] = "text/plain; charset=utf-8"
		req.headers["Accept"] = "application/json, text/html"
	default:
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(rawBody)
		if err != nil {
			return err
		}
		req.rawBody = buf
		req.headers["Accept"] = defaultHeaderAccept
		req.headers["Content-Type"] = defaultHeaderContentType
	}
	return nil
}

func (req *request) setAuthorization(
	auth auth.AuthorizationProvider,
) error {
	authHeader, err := auth.GetAuthorizationHeader()
	if err != nil {
		return fmt.Errorf(
			"request.setAuthorization: failed to get authorization: %v",
			err)
	}
	req.headers["Authorization"] = authHeader
	return nil
}

// ecRequestSender sends requests to the Edgecast API
type ecRequestSender struct {
	clientAdapter clientAdapter
}

// newECRequestSender creates a default instance of ecRequestSender
func newECRequestSender(config ClientConfig) ecRequestSender {
	return ecRequestSender{
		clientAdapter: newRetryableHTTPClientAdapter(config),
	}
}

// literalResponse is used for unmarshaling response data
// that is in an unrecognized format
type literalResponse struct {
	value interface{}
}

// sendRequest sends an HTTP request and returns the Response with any Response
// body data, if applicable
func (es ecRequestSender) sendRequest(req request) (*response, error) {
	httpResp, err := es.clientAdapter.do(req)
	if err != nil {
		return nil, fmt.Errorf("SendRequest:%v", err)
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	bodyAsString := string(body)
	if httpResp.StatusCode >= 400 && httpResp.StatusCode <= 599 {
		if err != nil {
			return nil,
				fmt.Errorf(
					"ecRequestSender.sendRequest: ioutil.ReadAll: %v",
					err)
		}
		return nil,
			fmt.Errorf(
				"ecRequestSender.sendRequest failed: %s", bodyAsString)
	}

	var temp interface{}
	if jsonUnmarshalErr := json.Unmarshal(body, &temp); err != nil {
		return nil,
			fmt.Errorf(
				"ecRequestSender.sendRequest: malformed JSON response: %v",
				jsonUnmarshalErr)
	}

	if eccollections.IsInterfaceArray(temp) {
		if err := json.Unmarshal([]byte(body), &req.parsedResponse); err != nil {
			return nil,
				fmt.Errorf(
					"ecRequestSender.sendRequest: malformed Json Array response:%v",
					err)
		}
	} else {
		if ecjson.IsJSONString(bodyAsString) {
			err = json.Unmarshal([]byte(bodyAsString), &req.parsedResponse)
			if err != nil {
				return nil, fmt.Errorf(
					"ecRequestSender.sendRequest: Decode error: %v",
					err)
			}
		} else {
			// if response is not json string
			switch v := req.parsedResponse.(type) {
			case literalResponse:
				rs, ok := req.parsedResponse.(literalResponse)
				if ok {
					rs.value = bodyAsString
					req.parsedResponse = rs
				}
			case float64:
				fmt.Println("float64:", v)
			default:
				fmt.Println("unknown")
			}
		}
	}
	return &response{
		data:         req.parsedResponse,
		httpResponse: httpResp,
	}, nil
}

// sendRequestWithStringResponse sends an HTTP request and returns the response
// body as a string
func (es ecRequestSender) sendRequestWithStringResponse(
	req request,
) (*string, error) {
	httpResp, err := es.clientAdapter.do(req)
	if err != nil {
		return nil, fmt.Errorf(
			"ecRequestSender.sendRequestWithStringResponse: %v",
			err)
	}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	bodyAsString := string(body)
	if httpResp.StatusCode >= 400 && httpResp.StatusCode <= 599 {
		if err != nil {
			return nil, fmt.Errorf(
				"ecRequestSender.sendRequestWithStringResponse: ioutil.ReadAll: %v",
				err)
		}

		return nil, fmt.Errorf(
			"ecRequestSender.sendRequestWithStringResponse failed: %s",
			bodyAsString)
	}
	return &bodyAsString, nil
}
