/*
 * 3gpp-as-session-with-qos
 *
 * API for setting us an AS session with required QoS.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.122 V17.9.0 T8 reference point for Northbound APIs
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.122/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents QoS requirements for time sensitive communication.
type TscQosRequirement struct {
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	ReqGbrDl string `json:"reqGbrDl,omitempty" yaml:"reqGbrDl" bson:"reqGbrDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	ReqGbrUl string `json:"reqGbrUl,omitempty" yaml:"reqGbrUl" bson:"reqGbrUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	ReqMbrDl string `json:"reqMbrDl,omitempty" yaml:"reqMbrDl" bson:"reqMbrDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	ReqMbrUl string `json:"reqMbrUl,omitempty" yaml:"reqMbrUl" bson:"reqMbrUl,omitempty"`
	// Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.
	MaxTscBurstSize int32 `json:"maxTscBurstSize,omitempty" yaml:"maxTscBurstSize" bson:"maxTscBurstSize,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	Req5Gsdelay int32 `json:"req5Gsdelay,omitempty" yaml:"req5Gsdelay" bson:"req5Gsdelay,omitempty"`
	// Represents the priority level of TSC Flows.
	Priority int32 `json:"priority,omitempty" yaml:"priority" bson:"priority,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	TscaiTimeDom int32                `json:"tscaiTimeDom,omitempty" yaml:"tscaiTimeDom" bson:"tscaiTimeDom,omitempty"`
	TscaiInputDl *TscaiInputContainer `json:"tscaiInputDl,omitempty" yaml:"tscaiInputDl" bson:"tscaiInputDl,omitempty"`
	TscaiInputUl *TscaiInputContainer `json:"tscaiInputUl,omitempty" yaml:"tscaiInputUl" bson:"tscaiInputUl,omitempty"`
}