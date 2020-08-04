/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type TargetingsGetListStruct struct {
	TargetingId                   int64                 `json:"targeting_id,omitempty"`
	TargetingName                 string                `json:"targeting_name,omitempty"`
	IsIncludeUnsupportedTargeting bool                  `json:"is_include_unsupported_targeting,omitempty"`
	Description                   string                `json:"description,omitempty"`
	IsDeleted                     bool                  `json:"is_deleted,omitempty"`
	CreatedTime                   int64                 `json:"created_time,omitempty"`
	LastModifiedTime              int64                 `json:"last_modified_time,omitempty"`
	AdLockStatus                  AdLockStatus          `json:"ad_lock_status,omitempty"`
	TargetingTranslation          string                `json:"targeting_translation,omitempty"`
	Targeting                     *ReadTargetingSetting `json:"targeting,omitempty"`
}
