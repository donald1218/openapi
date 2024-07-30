/*
 * Nbsf_Management
 *
 * Binding Support Management Service API. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nbsf_Management

// APIClient manages communication with the Nbsf_Management API v1.1.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	IndividualPCFBindingDocumentApi *IndividualPCFBindingDocumentApiService
	PCFBindingsCollectionApi        *PCFBindingsCollectionApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.IndividualPCFBindingDocumentApi = (*IndividualPCFBindingDocumentApiService)(&c.common)
	c.PCFBindingsCollectionApi = (*PCFBindingsCollectionApiService)(&c.common)

	return c
}