package goqontak

import (
	"net/http"
	"net/url"
)

// Options ...
type Options struct {
	httpClient *http.Client
	baseURL    *url.URL
}

type Error struct {
	Code     int      `json:"code"`
	Messages []string `json:"messages"`
}

// ResponseError ...
type ResponseError struct {
	Status string `json:"status,omitempty"`
	Error  Error  `json:"error"`
}

type GenerateAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    string `json:"created_at"`
}

type ValidateWhatsAppNumberTokenResponse struct {
	Response
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ValidateWhatsAppNumberTokenParams struct {
	ChannelIntegrationID string   `json:"channel_integration_id"`
	PhoneNumbers         []string `json:"phone_numbers"`
}

type SendWhatsAppDirectParams struct {
	ToNumber             string      `json:"to_number"`
	ToName               string      `json:"to_name"`
	MessageTemplateID    string      `json:"message_template_id"`
	ChannelIntegrationID string      `json:"channel_integration_id"`
	Language             Language    `json:"language"`
	Parameters           interface{} `json:"parameters"`
}

type Language struct {
	Code string `json:"code"`
}

type Parameter struct {
	Body []Body `json:"body"`
}

type ParameterWithHeader struct {
	Body   []Body `json:"body"`
	Header Header `json:"header"`
}

type Header struct {
	Format      string        `json:"format"`
	HeaderParam []HeaderParam `json:"params"`
}

type HeaderParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Body struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	ValueText string `json:"value_text"`
}

type SendWhatsAppDirectResponse struct {
	Response
}

type GenerateAccessTokenParams struct {
	UserName     string `json:"username"`
	Password     string `json:"password"`
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type WebhookMessageInteractionParams struct {
	ReceiveMessageFromAgent    bool   `json:"receive_message_from_agent"`
	ReceiveMessageFromCustomer bool   `json:"receive_message_from_customer"`
	StatusMessage              bool   `json:"status_message"`
	Url                        string `json:"url"`
}

type WebhookMessageInteractionResponse struct {
	Response
}
