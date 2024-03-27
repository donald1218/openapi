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

import (
	"time"
)

// The redundant transmission experience per Time Slot.
type RedundantTransmissionExpPerTs struct {
	// string with format 'date-time' as defined in OpenAPI.
	TsStart *time.Time `json:"tsStart" yaml:"tsStart" bson:"tsStart,omitempty"`
	// indicating a time in seconds.
	TsDuration      int32                      `json:"tsDuration" yaml:"tsDuration" bson:"tsDuration,omitempty"`
	ObsvRedTransExp *ObservedRedundantTransExp `json:"obsvRedTransExp" yaml:"obsvRedTransExp" bson:"obsvRedTransExp,omitempty"`
	// Redundant Transmission Status. Set to \"true\" if redundant transmission was activated, otherwise set to \"false\". Default value is \"false\" if omitted.
	RedTransStatus bool `json:"redTransStatus,omitempty" yaml:"redTransStatus" bson:"redTransStatus,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	UeRatio int32 `json:"ueRatio,omitempty" yaml:"ueRatio" bson:"ueRatio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty" yaml:"confidence" bson:"confidence,omitempty"`
}