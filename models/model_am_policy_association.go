/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.507 V16.9.0; 5G System; Access and Mobility Policy Control Service.
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.507/
 *
 * API version: 1.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AmPolicyAssociation struct {
	Request *AmPolicyAssociationRequest `json:"request,omitempty" yaml:"request" bson:"request,omitempty"`
	// Request Triggers that the PCF subscribes.
	Triggers      []AmPolicyRequestTrigger        `json:"triggers,omitempty" yaml:"triggers" bson:"triggers,omitempty"`
	ServAreaRes   *ServiceAreaRestriction         `json:"servAreaRes,omitempty" yaml:"servAreaRes" bson:"servAreaRes,omitempty"`
	WlServAreaRes *WirelineServiceAreaRestriction `json:"wlServAreaRes,omitempty" yaml:"wlServAreaRes" bson:"wlServAreaRes,omitempty"`
	Rfsp          int32                           `json:"rfsp,omitempty" yaml:"rfsp" bson:"rfsp,omitempty"`
	SmfSelInfo    *SmfSelectionData               `json:"smfSelInfo,omitempty" yaml:"smfSelInfo" bson:"smfSelInfo,omitempty"`
	UeAmbr        *Ambr                           `json:"ueAmbr,omitempty" yaml:"ueAmbr" bson:"ueAmbr,omitempty"`
	Pras          map[string]PresenceInfo         `json:"pras,omitempty" yaml:"pras" bson:"pras,omitempty"`
	SuppFeat      string                          `json:"suppFeat" yaml:"suppFeat" bson:"suppFeat,omitempty"`
}