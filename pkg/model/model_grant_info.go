/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 权限信息
type GrantInfo struct {
	AudienceId int64             `json:"audience_id,omitempty"`
	GrantType  AudienceGrantType `json:"grant_type,omitempty"`
	GrantSpec  *GrantSpec        `json:"grant_spec,omitempty"`
}
