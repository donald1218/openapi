/*
 * Nnef_SMContext
 *
 * Nnef SMContext Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.541 V17.3.0; 5G System; Session Management Services for Non-IP Data Delivery (NIDD).
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.541/
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Representation of the SM Context status notification.
type NefSmContextSmContextStatusNotification struct {
	Status SmContextStatus `json:"status" yaml:"status" bson:"status,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	SmContextId         string                   `json:"smContextId" yaml:"smContextId" bson:"smContextId,omitempty"`
	Cause               NefSmContextReleaseCause `json:"cause,omitempty" yaml:"cause" bson:"cause,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus     `json:"smallDataRateStatus,omitempty" yaml:"smallDataRateStatus" bson:"smallDataRateStatus,omitempty"`
	ApnRateStatus       *ApnRateStatus           `json:"apnRateStatus,omitempty" yaml:"apnRateStatus" bson:"apnRateStatus,omitempty"`
}