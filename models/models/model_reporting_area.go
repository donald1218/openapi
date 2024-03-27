/*
 * LMF Location
 *
 * LMF Location Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.572 V17.9.0; 5G System; Location Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.572/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Indicates an area for event reporting.
type ReportingArea struct {
	AreaType ReportingAreaType `json:"areaType" yaml:"areaType" bson:"areaType,omitempty"`
	Tai      *Tai              `json:"tai,omitempty" yaml:"tai" bson:"tai,omitempty"`
	Ecgi     *Ecgi             `json:"ecgi,omitempty" yaml:"ecgi" bson:"ecgi,omitempty"`
	Ncgi     *Ncgi             `json:"ncgi,omitempty" yaml:"ncgi" bson:"ncgi,omitempty"`
}