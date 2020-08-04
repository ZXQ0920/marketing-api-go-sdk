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
type WechatPagesGetListStruct struct {
	PageId               int64                `json:"page_id,omitempty"`
	PageName             string               `json:"page_name,omitempty"`
	CreatedTime          int64                `json:"created_time,omitempty"`
	LastModifiedTime     int64                `json:"last_modified_time,omitempty"`
	PageTemplateId       int64                `json:"page_template_id,omitempty"`
	ShareContentSpec     *ShareContentSpec    `json:"share_content_spec,omitempty"`
	PreviewUrl           string               `json:"preview_url,omitempty"`
	PageType             PageTypeRead         `json:"page_type,omitempty"`
	SourceType           WechatPageSourceType `json:"source_type,omitempty"`
	PageElementsSpecList []PageElementsStruct `json:"page_elements_spec_list,omitempty"`
}
