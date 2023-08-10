package goqontak

import (
	"net/http"
	"net/url"
)

// Builder ...
type Builder struct {
	option  Options
	builder HTTPServiceClient
}

// SetBuilder ...
func (c *Builder) SetBuilder(b HTTPServiceClient) {
	c.builder = b
}

// SetHTTPClient ...
func (c *Builder) SetHTTPClient(httpClient *http.Client) {
	c.option.httpClient = httpClient
}

// SetBaseURL ...
func (c *Builder) SetBaseURL(baseURL *url.URL) {
	c.option.baseURL = baseURL
}

// Build ...
func (c *Builder) Build() HTTPServiceClient {
	c.builder.setOptions(c.option)
	return c.builder
}
