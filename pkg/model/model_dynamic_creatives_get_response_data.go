/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type DynamicCreativesGetResponseData struct {
	List     []DynamicCreativesGetListStruct `json:"list,omitempty"`
	PageInfo *ConfPageSize500                `json:"page_info,omitempty"`
}
