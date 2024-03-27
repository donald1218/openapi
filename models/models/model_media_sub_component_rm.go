/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.514 V17.9.0; 5G System; Policy Authorization Service; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.514/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// This data type is defined in the same way as the MediaSubComponent data type, but with the OpenAPI nullable property set to true. Removable attributes marBwDl and marBwUl are defined with the corresponding removable data type.
type MediaSubComponentRm struct {
	AfSigProtocol AfSigProtocol        `json:"afSigProtocol,omitempty" yaml:"afSigProtocol" bson:"afSigProtocol,omitempty"`
	EthfDescs     []EthFlowDescription `json:"ethfDescs,omitempty" yaml:"ethfDescs" bson:"ethfDescs,omitempty"`
	FNum          int32                `json:"fNum" yaml:"fNum" bson:"fNum,omitempty"`
	FDescs        []string             `json:"fDescs,omitempty" yaml:"fDescs" bson:"fDescs,omitempty"`
	FStatus       FlowStatus           `json:"fStatus,omitempty" yaml:"fStatus" bson:"fStatus,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	MarBwDl string `json:"marBwDl,omitempty" yaml:"marBwDl" bson:"marBwDl,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property.
	MarBwUl string `json:"marBwUl,omitempty" yaml:"marBwUl" bson:"marBwUl,omitempty"`
	// This data type is defined in the same way as the TosTrafficClass data type, but with the OpenAPI nullable property set to true.
	TosTrCl   string    `json:"tosTrCl,omitempty" yaml:"tosTrCl" bson:"tosTrCl,omitempty"`
	FlowUsage FlowUsage `json:"flowUsage,omitempty" yaml:"flowUsage" bson:"flowUsage,omitempty"`
}