/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type PlayablePagesGetResponseData struct {
	List     []PlayablePagesGetListStruct `json:"list,omitempty"`
	PageInfo *Conf                        `json:"page_info,omitempty"`
}
