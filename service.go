package goqontak

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Service ...
type Service struct {
	options Options
}

func (s *Service) setOptions(options Options) HTTPServiceClient {
	s.options = options
	return s
}

func (s *Service) GenerateAccessToken(ctx context.Context, request interface{}) (returnData GenerateAccessTokenResponse, errorResponse ResponseError, resp *http.Response, err error) {
	values, ok := request.(GenerateAccessTokenParams)
	if !ok {
		err = ErrStatusBadRequest
		return
	}

	body, _ := json.Marshal(values)
	req, _ := s.newRequest(ctx, http.MethodPost, PathGenerateAccessToken, string(body))
	resp, err = s.do(ctx, req, &returnData, &errorResponse)

	return
}

func (s *Service) ValidateWhatsAppNumber(ctx context.Context, accessToken string, request interface{}) (returnData ValidateWhatsAppNumberTokenResponse, errorResponse ResponseError, resp *http.Response, err error) {
	values, ok := request.(ValidateWhatsAppNumberTokenParams)
	if !ok {
		err = ErrStatusBadRequest
		return
	}

	body, _ := json.Marshal(values)
	req, _ := s.newRequest(ctx, http.MethodPost, PathValidateWhatsAppNumber, string(body))
	req.Header.Add(HeaderKeyAuthorization, fmt.Sprintf("Bearer %s", accessToken))
	resp, err = s.do(ctx, req, &returnData, &errorResponse)

	return
}

func (s *Service) SendWhatsAppDirect(ctx context.Context, accessToken string, request interface{}) (returnData SendWhatsAppDirectResponse, errorResponse ResponseError, resp *http.Response, err error) {
	values, ok := request.(SendWhatsAppDirectParams)
	if !ok {
		err = ErrStatusBadRequest
		return
	}

	body, _ := json.Marshal(values)
	req, _ := s.newRequest(ctx, http.MethodPost, PathSendWhatsAppDirect, string(body))
	req.Header.Add(HeaderKeyAuthorization, fmt.Sprintf("Bearer %s", accessToken))
	resp, err = s.do(ctx, req, &returnData, &errorResponse)

	return
}

func (s *Service) newRequest(ctx context.Context, method, path string, body interface{}) (req *http.Request, err error) {
	rel := &url.URL{Path: path}
	u := s.options.baseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = bytes.NewBufferString(body.(string))
	}

	req, err = http.NewRequestWithContext(ctx, method, u.String(), buf)

	if err != nil {
		return
	}

	req.Header.Set(ContentType, ApplicationJSONCharsetUtf8)
	return
}

func (s *Service) do(ctx context.Context, req *http.Request, normalResponse interface{}, errorResponse interface{}) (resp *http.Response, err error) {
	resp, err = s.options.httpClient.Do(req)

	if err != nil {
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	// reset the response body to the original unread state
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByte))

	_ = json.Unmarshal(bodyByte, &normalResponse)
	_ = json.Unmarshal(bodyByte, &errorResponse)

	err = s.lookUpErrStatusCode(resp.StatusCode)

	return
}

// lookUpErrStatusCode ...
func (s *Service) lookUpErrStatusCode(statusCode int) (err error) {
	// 400	You have sent bad request data.
	// 401	Access denied due to unauthorized transaction. Please check AccessToken.
	// 404	The requested resource is not found.
	// 405	HTTP method is not allowed.
	// 500	Happens within the web server attempting to show you a web page. It's usually a server-side problem out of your control.
	// 502	The server, while acting as a gateway or proxy, received an invalid response from the upstream server
	draftErrors := map[int]error{
		http.StatusBadRequest:          ErrStatusBadRequest,
		http.StatusUnauthorized:        ErrStatusUnauthorized,
		http.StatusNotFound:            ErrStatusNotFound,
		http.StatusMethodNotAllowed:    ErrStatusMethodNotAllowed,
		http.StatusInternalServerError: ErrStatusInternalServerError,
		http.StatusBadGateway:          ErrStatusBadGateway,
		http.StatusUnprocessableEntity: ErrStatusUnprocessableEntity,
	}

	if msg, ok := draftErrors[statusCode]; ok {
		err = msg
	}

	return
}

func (s *Service) WebhookMessageInteraction(ctx context.Context, accessToken string, request interface{}) (returnData WebhookMessageInteractionResponse, errorResponse ResponseError, resp *http.Response, err error) {
	values, ok := request.(WebhookMessageInteractionParams)
	if !ok {
		err = ErrStatusBadRequest
		return
	}

	body, _ := json.Marshal(values)
	req, _ := s.newRequest(ctx, http.MethodPut, PathWebhookMessageInteraction, string(body))
	req.Header.Add(HeaderKeyAuthorization, fmt.Sprintf("Bearer %s", accessToken))
	resp, err = s.do(ctx, req, &returnData, &errorResponse)

	return
}
