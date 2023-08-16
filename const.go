package goqontak

const (
	ContentType                   = "Content-Type"
	ApplicationJSONCharsetUtf8    = "application/json; charset=utf-8"
	HeaderKeyAuthorization        = "Authorization"
	PathValidateWhatsAppNumber    = "/api/open/v1/broadcasts/contacts"
	PathSendWhatsAppDirect        = "/api/open/v1/broadcasts/whatsapp/direct"
	PathGenerateAccessToken       = "/oauth/token"
	PathWebhookMessageInteraction = "/api/open/v1/message_interactions"
)
