/*
 * Npcf_UEPolicyControl
 *
 * UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.525 V17.9.0; 5G System; UE Policy Control Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.525/
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type N1N2MessageTransferCause string

// List of N1N2MessageTransferCause
const (
	N1N2MessageTransferCause_ATTEMPTING_TO_REACH_UE                N1N2MessageTransferCause = "ATTEMPTING_TO_REACH_UE"
	N1N2MessageTransferCause_N1_N2_TRANSFER_INITIATED              N1N2MessageTransferCause = "N1_N2_TRANSFER_INITIATED"
	N1N2MessageTransferCause_WAITING_FOR_ASYNCHRONOUS_TRANSFER     N1N2MessageTransferCause = "WAITING_FOR_ASYNCHRONOUS_TRANSFER"
	N1N2MessageTransferCause_UE_NOT_RESPONDING                     N1N2MessageTransferCause = "UE_NOT_RESPONDING"
	N1N2MessageTransferCause_N1_MSG_NOT_TRANSFERRED                N1N2MessageTransferCause = "N1_MSG_NOT_TRANSFERRED"
	N1N2MessageTransferCause_N2_MSG_NOT_TRANSFERRED                N1N2MessageTransferCause = "N2_MSG_NOT_TRANSFERRED"
	N1N2MessageTransferCause_UE_NOT_REACHABLE_FOR_SESSION          N1N2MessageTransferCause = "UE_NOT_REACHABLE_FOR_SESSION"
	N1N2MessageTransferCause_TEMPORARY_REJECT_REGISTRATION_ONGOING N1N2MessageTransferCause = "TEMPORARY_REJECT_REGISTRATION_ONGOING"
	N1N2MessageTransferCause_TEMPORARY_REJECT_HANDOVER_ONGOING     N1N2MessageTransferCause = "TEMPORARY_REJECT_HANDOVER_ONGOING"
	N1N2MessageTransferCause_REJECTION_DUE_TO_PAGING_RESTRICTION   N1N2MessageTransferCause = "REJECTION_DUE_TO_PAGING_RESTRICTION"
	N1N2MessageTransferCause_AN_NOT_RESPONDING                     N1N2MessageTransferCause = "AN_NOT_RESPONDING"
	N1N2MessageTransferCause_FAILURE_CAUSE_UNSPECIFIED             N1N2MessageTransferCause = "FAILURE_CAUSE_UNSPECIFIED"
)