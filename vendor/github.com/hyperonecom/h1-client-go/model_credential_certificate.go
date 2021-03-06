/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CredentialCertificate struct for CredentialCertificate
type CredentialCertificate struct {
	CreatedBy   string `json:"createdBy,omitempty"`
	CreatedOn   string `json:"createdOn,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Id          string `json:"id,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	Value       string `json:"value,omitempty"`
}
