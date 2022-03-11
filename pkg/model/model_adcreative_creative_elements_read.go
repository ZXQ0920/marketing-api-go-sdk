/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意元素
type AdcreativeCreativeElementsRead struct {
	Image                      *string                            `json:"image,omitempty"`
	Image2                     *string                            `json:"image2,omitempty"`
	Image3                     *string                            `json:"image3,omitempty"`
	Title                      *string                            `json:"title,omitempty"`
	Description                *string                            `json:"description,omitempty"`
	Corporate                  *AdcreativeCorporate               `json:"corporate,omitempty"`
	Video                      *string                            `json:"video,omitempty"`
	DeepLinkType               *string                            `json:"deep_link_type,omitempty"`
	LinkNameType               LinkNameType                       `json:"link_name_type,omitempty"`
	ImageList                  *[]string                          `json:"image_list,omitempty"`
	ElementStory               *[]AdcreativeElementStoryArrayItem `json:"element_story,omitempty"`
	Url                        *string                            `json:"url,omitempty"`
	ButtonText                 *string                            `json:"button_text,omitempty"`
	BottomText                 *string                            `json:"bottom_text,omitempty"`
	ExcitationText             *string                            `json:"excitation_text,omitempty"`
	CountdownBegin             *int64                             `json:"countdown_begin,omitempty"`
	CountdownPrice             *string                            `json:"countdown_price,omitempty"`
	CountdownTimeType          AdCreativeCountdownTimeType        `json:"countdown_time_type,omitempty"`
	Label                      *[]CreativeLabel                   `json:"label,omitempty"`
	ProductTags                *[]string                          `json:"product_tags,omitempty"`
	LogoDescription            *string                            `json:"logo_description,omitempty"`
	Logo                       *string                            `json:"logo,omitempty"`
	LeftButton                 *string                            `json:"left_button,omitempty"`
	RightButton                *string                            `json:"right_button,omitempty"`
	LeftBottomTxt              *string                            `json:"left_bottom_txt,omitempty"`
	AnimationEffect            *string                            `json:"animation_effect,omitempty"`
	Phone                      *string                            `json:"phone,omitempty"`
	Caption                    *string                            `json:"caption,omitempty"`
	LogoPage                   *AdcreativeLogoPage                `json:"logo_page,omitempty"`
	VideoPopupUrl              *string                            `json:"video_popup_url,omitempty"`
	VideoPopupButton           *AdcreativeVideoPopupButton        `json:"video_popup_button,omitempty"`
	VideoPopupButtonText       *string                            `json:"video_popup_button_text,omitempty"`
	VideoPopupButtonUrl        *string                            `json:"video_popup_button_url,omitempty"`
	ButtonUrl                  *string                            `json:"button_url,omitempty"`
	LongVideo1                 *string                            `json:"long_video1,omitempty"`
	LongVideo2                 *string                            `json:"long_video2,omitempty"`
	ShortVideoStruct           *ShortVideoStruct                  `json:"short_video_struct,omitempty"`
	QzoneVideoPageId           *string                            `json:"qzone_video_page_id,omitempty"`
	Qq                         *string                            `json:"qq,omitempty"`
	LeftCanvas                 *string                            `json:"left_canvas,omitempty"`
	RightCanvas                *string                            `json:"right_canvas,omitempty"`
	SunText                    *string                            `json:"sun_text,omitempty"`
	CloudText                  *string                            `json:"cloud_text,omitempty"`
	OvercastText               *string                            `json:"overcast_text,omitempty"`
	RainText                   *string                            `json:"rain_text,omitempty"`
	SnowText                   *string                            `json:"snow_text,omitempty"`
	FogText                    *string                            `json:"fog_text,omitempty"`
	SandText                   *string                            `json:"sand_text,omitempty"`
	HazeText                   *string                            `json:"haze_text,omitempty"`
	LabelledImg                *AdcreativeLabelledImg             `json:"labelled_img,omitempty"`
	ShareImg                   *string                            `json:"share_img,omitempty"`
	LongVideoStruct            *LongVideoStruct                   `json:"long_video_struct,omitempty"`
	BannerContent              *AdcreativeBannerContent           `json:"banner_content,omitempty"`
	CardContent                *AdcreativeCardContent             `json:"card_content,omitempty"`
	Brand                      *AdCreativeBrand                   `json:"brand,omitempty"`
	FullScreenImage            *string                            `json:"full_screen_image,omitempty"`
	ZipUrl                     *string                            `json:"zip_url,omitempty"`
	EndPage                    *AdCreativeEndPage                 `json:"end_page,omitempty"`
	ShopImage                  *string                            `json:"shop_image,omitempty"`
	HeadLine                   *string                            `json:"head_line,omitempty"`
	ShopImageStruct            *AdCreativeShopImageStruct         `json:"shop_image_struct,omitempty"`
	ChosenButton               *ChosenButton                      `json:"chosen_button,omitempty"`
	LivingDescStruct           *AdCreativeLivingDescStruct        `json:"living_desc_struct,omitempty"`
	FloatingZoneStruct         *FloatingZone                      `json:"floating_zone_struct,omitempty"`
	CanvasShareImage           *string                            `json:"canvas_share_image,omitempty"`
	CountdownExpiringTimestamp *int64                             `json:"countdown_expiring_timestamp,omitempty"`
}
