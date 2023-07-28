/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SteeringMode struct {
	SteerModeValue SteerModeValue `json:"steerModeValue" yaml:"steerModeValue" bson:"steerModeValue" mapstructure:"SteerModeValue"`
	Active         AccessType     `json:"active,omitempty" yaml:"active" bson:"active" mapstructure:"Active"`
	Standby        *AccessTypeRm  `json:"standby,omitempty" yaml:"standby" bson:"standby" mapstructure:"Standby"`
	Var3gLoad      int32          `json:"3gLoad,omitempty" yaml:"3gLoad" bson:"3gLoad" mapstructure:"Var3gLoad"`
	PrioAcc        AccessType     `json:"prioAcc,omitempty" yaml:"prioAcc" bson:"prioAcc" mapstructure:"PrioAcc"`
}