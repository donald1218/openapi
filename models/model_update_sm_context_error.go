/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.502 V16.9.0; 5G System; Session Management Services; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.502/
 *
 * API version: 1.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UpdateSmContextError struct {
	JsonData                  *SmContextUpdateError `json:"jsonData,omitempty" yaml:"jsonData" bson:"jsonData" multipart:"contentType:application/json,omitempty"`
	BinaryDataN1SmMessage     []byte                `json:"binaryDataN1SmMessage,omitempty" yaml:"binaryDataN1SmMessage" bson:"binaryDataN1SmMessage" multipart:"contentType:application/vnd.3gpp.5gnas,ref:JsonData.N1SmMsg.ContentId,omitempty"`
	BinaryDataN2SmInformation []byte                `json:"binaryDataN2SmInformation,omitempty" yaml:"binaryDataN2SmInformation" bson:"binaryDataN2SmInformation" multipart:"contentType:application/vnd.3gpp.ngap,ref:JsonData.N2SmInfo.ContentId,omitempty"`
}