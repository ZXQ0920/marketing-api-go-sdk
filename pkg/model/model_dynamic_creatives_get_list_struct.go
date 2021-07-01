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
type DynamicCreativesGetListStruct struct {
	DynamicCreativeId         *int64                       `json:"dynamic_creative_id,omitempty"`
	OuterAdcreativeId         *int64                       `json:"outer_adcreative_id,omitempty"`
	DynamicCreativeName       *string                      `json:"dynamic_creative_name,omitempty"`
	DynamicCreativeTemplateId *int64                       `json:"dynamic_creative_template_id,omitempty"`
	DynamicCreativeElements   *DynamicCreativeElementsRead `json:"dynamic_creative_elements,omitempty"`
	PageType                  PageTypeRead                 `json:"page_type,omitempty"`
	PageSpec                  *DynamicCreativePageSpec     `json:"page_spec,omitempty"`
	DeepLinkUrl               *string                      `json:"deep_link_url,omitempty"`
	AutomaticSiteEnabled      *bool                        `json:"automatic_site_enabled,omitempty"`
	SiteSet                   *[]string                    `json:"site_set,omitempty"`
	PromotedObjectType        PromotedObjectType           `json:"promoted_object_type,omitempty"`
	PromotedObjectId          *string                      `json:"promoted_object_id,omitempty"`
	ProfileId                 *int64                       `json:"profile_id,omitempty"`
	CreatedTime               *int64                       `json:"created_time,omitempty"`
	LastModifiedTime          *int64                       `json:"last_modified_time,omitempty"`
	IsDeleted                 *bool                        `json:"is_deleted,omitempty"`
	CampaignType              CampaignType                 `json:"campaign_type,omitempty"`
	ImpressionTrackingUrl     *string                      `json:"impression_tracking_url,omitempty"`
	ClickTrackingUrl          *string                      `json:"click_tracking_url,omitempty"`
	FeedsVideoCommentSwitch   *bool                        `json:"feeds_video_comment_switch,omitempty"`
	UnionMarketSwitch         *bool                        `json:"union_market_switch,omitempty"`
	VideoEndPage              *VideoEndPageSpec            `json:"video_end_page,omitempty"`
	BarrageList               *[]BarrageListReadStruct     `json:"barrage_list,omitempty"`
	DynamicCreativeGroupUsed  DynamicCreativeGroupUsed     `json:"dynamic_creative_group_used,omitempty"`
}
