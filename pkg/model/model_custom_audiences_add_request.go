/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type CustomAudiencesAddRequest struct {
	AccountId       int64         `json:"account_id,omitempty"`
	Name            string        `json:"name,omitempty"`
	Type_           AudienceType  `json:"type,omitempty"`
	OuterAudienceId string        `json:"outer_audience_id,omitempty"`
	Description     string        `json:"description,omitempty"`
	AudienceSpec    *AudienceSpec `json:"audience_spec,omitempty"`
	Platform        DataPlatform  `json:"platform,omitempty"`
}
