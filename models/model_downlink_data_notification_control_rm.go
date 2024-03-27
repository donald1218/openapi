/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.512 V17.11.0; 5G System; Session Management Policy Control Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// This data type is defined in the same way as the DownlinkDataNotificationControl data type, but with the nullable:true property.
type DownlinkDataNotificationControlRm struct {
	NotifCtrlInds []NotificationControlIndication `json:"notifCtrlInds,omitempty" yaml:"notifCtrlInds" bson:"notifCtrlInds,omitempty"`
	TypesOfNotif  []DlDataDeliveryStatus          `json:"typesOfNotif,omitempty" yaml:"typesOfNotif" bson:"typesOfNotif,omitempty"`
}