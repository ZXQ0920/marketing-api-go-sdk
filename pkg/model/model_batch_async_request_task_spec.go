/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 任务所需条件
type BatchAsyncRequestTaskSpec struct {
	UpdateUnionPositionPackageSpec              *[]UpdateUnionPositionPackageItem              `json:"update_union_position_package_spec,omitempty"`
	UpdateExcludeUnionPositionPackageSpec       *[]UpdateExcludeUnionPositionPackageItem       `json:"update_exclude_union_position_package_spec,omitempty"`
	UpdateTargetingIdSpec                       *[]UpdateTargetingIdItem                       `json:"update_targeting_id_spec,omitempty"`
	UpdateBidStrategySpec                       *[]UpdateBidStrategyItem                       `json:"update_bid_strategy_spec,omitempty"`
	UpdateDeepConversionBehaviorBidSpec         *[]UpdateDeepConversionBehaviorBidItem         `json:"update_deep_conversion_behavior_bid_spec,omitempty"`
	UpdateAdgroupAppAndroidChannelPackageIdSpec *[]UpdateAdgroupAppAndroidChannelPackageIdItem `json:"update_adgroup_app_android_channel_package_id_spec,omitempty"`
	UpdateCampaignSpeedModeSpec                 *[]UpdateCampaignSpeedModeItem                 `json:"update_campaign_speed_mode_spec,omitempty"`
	DeleteCampaignSpec                          *[]DeleteCampaignItem                          `json:"delete_campaign_spec,omitempty"`
	DeleteAdgroupSpec                           *[]DeleteAdgroupItem                           `json:"delete_adgroup_spec,omitempty"`
	DeleteAdSpec                                *[]DeleteAdItem                                `json:"delete_ad_spec,omitempty"`
	UpdateAdgroupDeepConversionWorthRateSpec    *[]UpdateAdgroupDeepConversionWorthRateItem    `json:"update_adgroup_deep_conversion_worth_rate_spec,omitempty"`
	UpdateAdcreativeDeepLinkUrlSpec             *[]UpdateAdcreativeDeepLinkUrlItem             `json:"update_adcreative_deep_link_url_spec,omitempty"`
	TargetingsShareSpec                         *[]TargetingsShareItem                         `json:"targetings_share_spec,omitempty"`
	UpdateCampaignConfiguredStatusSpec          *[]UpdateCampaignConfiguredStatusItem          `json:"update_campaign_configured_status_spec,omitempty"`
	UpdateCampaignDailyBudgetSpec               *[]UpdateCampaignDailyBudgetItem               `json:"update_campaign_daily_budget_spec,omitempty"`
	UpdateAdgroupConfiguredStatusSpec           *[]UpdateAdgroupConfiguredStatusItem           `json:"update_adgroup_configured_status_spec,omitempty"`
	UpdateAdgroupDailyBudgetSpec                *[]UpdateAdgroupDailyBudgetItem                `json:"update_adgroup_daily_budget_spec,omitempty"`
	UpdateAdConfiguredStatusSpec                *[]UpdateAdConfiguredStatusItem                `json:"update_ad_configured_status_spec,omitempty"`
	UpdateAdgroupAutoAcquisitionSpec            *[]UpdateAdgroupAutoAcquisitionItem            `json:"update_adgroup_auto_acquisition_spec,omitempty"`
	UpdateAdcreativeLandingPageSpec             *[]UpdateAdcreativeLandingPageItem             `json:"update_adcreative_landing_page_spec,omitempty"`
}
