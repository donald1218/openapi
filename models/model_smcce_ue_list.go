/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.10.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents the List of UEs classified based on experience level of Session Management  congestion control.
type SmcceUeList struct {
	HighLevel   []string `json:"highLevel,omitempty" yaml:"highLevel" bson:"highLevel,omitempty"`
	MediumLevel []string `json:"mediumLevel,omitempty" yaml:"mediumLevel" bson:"mediumLevel,omitempty"`
	LowLevel    []string `json:"lowLevel,omitempty" yaml:"lowLevel" bson:"lowLevel,omitempty"`
}