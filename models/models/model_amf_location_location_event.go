/*
 * Namf_Location
 *
 * AMF Location Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AmfLocationLocationEvent string

// List of AmfLocationLocationEvent
const (
	AmfLocationLocationEvent_EMERGENCY_CALL_ORIGINATION        AmfLocationLocationEvent = "EMERGENCY_CALL_ORIGINATION"
	AmfLocationLocationEvent_EMERGENCY_CALL_RELEASE            AmfLocationLocationEvent = "EMERGENCY_CALL_RELEASE"
	AmfLocationLocationEvent_EMERGENCY_CALL_HANDOVER           AmfLocationLocationEvent = "EMERGENCY_CALL_HANDOVER"
	AmfLocationLocationEvent_ACTIVATION_OF_DEFERRED_LOCATION   AmfLocationLocationEvent = "ACTIVATION_OF_DEFERRED_LOCATION"
	AmfLocationLocationEvent_UE_MOBILITY_FOR_DEFERRED_LOCATION AmfLocationLocationEvent = "UE_MOBILITY_FOR_DEFERRED_LOCATION"
	AmfLocationLocationEvent_CANCELLATION_OF_DEFERRED_LOCATION AmfLocationLocationEvent = "CANCELLATION_OF_DEFERRED_LOCATION"
)