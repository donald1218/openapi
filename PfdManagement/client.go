/*
 * 3gpp-pfd-management
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package PfdManagement

import (
	"github.com/free5gc/openapi"
)

// APIClient manages communication with the 3gpp-pfd-management API v1.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	DefaultApi *DefaultApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.httpClient == nil {
		cfg.httpClient = openapi.GetDefaultHttpClient()
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.DefaultApi = (*DefaultApiService)(&c.common)

	return c
}
