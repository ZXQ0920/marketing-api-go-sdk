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
type WechatAdvertiserDetailGetListStruct struct {
	AccountId         int64                       `json:"account_id,omitempty"`
	WechatAccountName string                      `json:"wechat_account_name,omitempty"`
	SystemIndustryId  int64                       `json:"system_industry_id,omitempty"`
	WechatAccountId   string                      `json:"wechat_account_id,omitempty"`
	AccountType       WechatAdvertiserAccountType `json:"account_type,omitempty"`
	CorporationName   string                      `json:"corporation_name,omitempty"`
	AuthStatus        WechatAuthStatus            `json:"auth_status,omitempty"`
	AuthTime          int64                       `json:"auth_time,omitempty"`
	AgencyIdList      []int64                     `json:"agency_id_list,omitempty"`
	StaffWechatIdList []string                    `json:"staff_wechat_id_list,omitempty"`
	DailyBudget       int64                       `json:"daily_budget,omitempty"`
}
