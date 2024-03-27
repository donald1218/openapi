/*
 * Nudm_NIDDAU
 *
 * Nudm NIDD Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.8.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents a NIDD authorization update notification.
type NiddAuthUpdateNotification struct {
	NiddAuthUpdateInfoList []NiddAuthUpdateInfo `json:"niddAuthUpdateInfoList" yaml:"niddAuthUpdateInfoList" bson:"niddAuthUpdateInfoList,omitempty"`
}