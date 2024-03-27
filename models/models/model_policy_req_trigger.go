/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PolicyReqTrigger string

// List of PolicyReqTrigger
const (
	PolicyReqTrigger_LOCATION_CHANGE      PolicyReqTrigger = "LOCATION_CHANGE"
	PolicyReqTrigger_PRA_CHANGE           PolicyReqTrigger = "PRA_CHANGE"
	PolicyReqTrigger_ALLOWED_NSSAI_CHANGE PolicyReqTrigger = "ALLOWED_NSSAI_CHANGE"
	PolicyReqTrigger_NWDAF_DATA_CHANGE    PolicyReqTrigger = "NWDAF_DATA_CHANGE"
	PolicyReqTrigger_PLMN_CHANGE          PolicyReqTrigger = "PLMN_CHANGE"
	PolicyReqTrigger_CON_STATE_CHANGE     PolicyReqTrigger = "CON_STATE_CHANGE"
	PolicyReqTrigger_SMF_SELECT_CHANGE    PolicyReqTrigger = "SMF_SELECT_CHANGE"
	PolicyReqTrigger_ACCESS_TYPE_CHANGE   PolicyReqTrigger = "ACCESS_TYPE_CHANGE"
)