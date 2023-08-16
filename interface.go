package goqontak

import (
	"context"
	"net/http"
)

// HTTPServiceClient ...
type HTTPServiceClient interface {
	// setOptions ...
	setOptions(options Options) HTTPServiceClient

	// GenerateAccessToken ...
	GenerateAccessToken(ctx context.Context, request interface{}) (returnData GenerateAccessTokenResponse, errorResponse ResponseError, resp *http.Response, err error)

	// ValidateWhatsAppNumber ...
	ValidateWhatsAppNumber(ctx context.Context, accessToken string, request interface{}) (returnData ValidateWhatsAppNumberTokenResponse, errorResponse ResponseError, resp *http.Response, err error)

	// SendWhatsAppDirect ...
	SendWhatsAppDirect(ctx context.Context, accessToken string, request interface{}) (returnData SendWhatsAppDirectResponse, errorResponse ResponseError, resp *http.Response, err error)

	WebhookMessageInteraction(ctx context.Context, accessToken string, request interface{}) (returnData WebhookMessageInteractionResponse, errorResponse ResponseError, resp *http.Response, err error)
}
