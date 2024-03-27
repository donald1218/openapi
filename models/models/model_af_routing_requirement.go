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

// Describes the event information delivered in the subscription.
type AfRoutingRequirement struct {
	AppReloc      bool               `json:"appReloc,omitempty" yaml:"appReloc" bson:"appReloc,omitempty"`
	RouteToLocs   []*RouteToLocation `json:"routeToLocs,omitempty" yaml:"routeToLocs" bson:"routeToLocs,omitempty"`
	SpVal         *SpatialValidity   `json:"spVal,omitempty" yaml:"spVal" bson:"spVal,omitempty"`
	TempVals      []TemporalValidity `json:"tempVals,omitempty" yaml:"tempVals" bson:"tempVals,omitempty"`
	UpPathChgSub  *UpPathChgEvent    `json:"upPathChgSub,omitempty" yaml:"upPathChgSub" bson:"upPathChgSub,omitempty"`
	AddrPreserInd bool               `json:"addrPreserInd,omitempty" yaml:"addrPreserInd" bson:"addrPreserInd,omitempty"`
	// Indicates whether simultaneous connectivity should be temporarily maintained for the source and target PSA.
	SimConnInd bool `json:"simConnInd,omitempty" yaml:"simConnInd" bson:"simConnInd,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	SimConnTerm int32 `json:"simConnTerm,omitempty" yaml:"simConnTerm" bson:"simConnTerm,omitempty"`
	// Contains EAS IP replacement information.
	EasIpReplaceInfos []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty" yaml:"easIpReplaceInfos" bson:"easIpReplaceInfos,omitempty"`
	// Indicates the EAS rediscovery is required.
	EasRedisInd bool `json:"easRedisInd,omitempty" yaml:"easRedisInd" bson:"easRedisInd,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxAllowedUpLat int32 `json:"maxAllowedUpLat,omitempty" yaml:"maxAllowedUpLat" bson:"maxAllowedUpLat,omitempty"`
}