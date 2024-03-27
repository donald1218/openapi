/*
 * Namf_MBSBroadcast
 *
 * AMF MBSBroadcast Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.518 V17.12.0; 5G System; Access and Mobility Management Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/
 *
 * API version: 1.0.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Data within ContextUpdate Request
type AmfMbsBroadcastContextUpdateReqData struct {
	MbsServiceArea         *MbsServiceArea             `json:"mbsServiceArea,omitempty" yaml:"mbsServiceArea" bson:"mbsServiceArea,omitempty"`
	MbsServiceAreaInfoList []MbsServiceAreaInfo        `json:"mbsServiceAreaInfoList,omitempty" yaml:"mbsServiceAreaInfoList" bson:"mbsServiceAreaInfoList,omitempty"`
	N2MbsSmInfo            *AmfMbsBroadcastN2MbsSmInfo `json:"n2MbsSmInfo,omitempty" yaml:"n2MbsSmInfo" bson:"n2MbsSmInfo,omitempty"`
	RanIdList              []GlobalRanNodeId           `json:"ranIdList,omitempty" yaml:"ranIdList" bson:"ranIdList,omitempty"`
	NoNgapSignallingInd    bool                        `json:"noNgapSignallingInd,omitempty" yaml:"noNgapSignallingInd" bson:"noNgapSignallingInd,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NotifyUri string `json:"notifyUri,omitempty" yaml:"notifyUri" bson:"notifyUri,omitempty"`
	// indicating a time in seconds.
	MaxResponseTime    int32 `json:"maxResponseTime,omitempty" yaml:"maxResponseTime" bson:"maxResponseTime,omitempty"`
	N2MbsInfoChangeInd bool  `json:"n2MbsInfoChangeInd,omitempty" yaml:"n2MbsInfoChangeInd" bson:"n2MbsInfoChangeInd,omitempty"`
}