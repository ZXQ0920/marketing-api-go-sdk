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
type DynamicCreativeElements struct {
	Image                      string                            `json:"image,omitempty"`
	Image2                     string                            `json:"image2,omitempty"`
	Image3                     string                            `json:"image3,omitempty"`
	Title                      string                            `json:"title,omitempty"`
	Description                string                            `json:"description,omitempty"`
	Corporate                  *AdcreativeCorporate              `json:"corporate,omitempty"`
	Video                      string                            `json:"video,omitempty"`
	ImageOptions               []string                          `json:"image_options,omitempty"`
	TitleOptions               []string                          `json:"title_options,omitempty"`
	DescriptionOptions         []string                          `json:"description_options,omitempty"`
	VideoOptions               []string                          `json:"video_options,omitempty"`
	DeepLinkType               string                            `json:"deep_link_type,omitempty"`
	LinkNameType               LinkNameType                      `json:"link_name_type,omitempty"`
	ImageList                  []string                          `json:"image_list,omitempty"`
	ElementStory               []AdcreativeElementStoryArrayItem `json:"element_story,omitempty"`
	Url                        string                            `json:"url,omitempty"`
	ButtonText                 string                            `json:"button_text,omitempty"`
	BottomText                 string                            `json:"bottom_text,omitempty"`
	MiniProgramId              string                            `json:"mini_program_id,omitempty"`
	MiniProgramPath            string                            `json:"mini_program_path,omitempty"`
	Label                      []AdcreativeLabel                 `json:"label,omitempty"`
	ProductTags                []string                          `json:"product_tags,omitempty"`
	LogoDescription            string                            `json:"logo_description,omitempty"`
	Logo                       string                            `json:"logo,omitempty"`
	LeftBottomTxt              string                            `json:"left_bottom_txt,omitempty"`
	AnimationEffect            string                            `json:"animation_effect,omitempty"`
	Phone                      string                            `json:"phone,omitempty"`
	ShortVideoStruct           *ShortVideoStruct                 `json:"short_video_struct,omitempty"`
	LongVideoStruct            *LongVideoStruct                  `json:"long_video_struct,omitempty"`
	BannerContent              *AdcreativeBannerContent          `json:"banner_content,omitempty"`
	CardContent                *AdcreativeCardContent            `json:"card_content,omitempty"`
	VideoPopupButton           *AdcreativeVideoPopupButton       `json:"video_popup_button,omitempty"`
	ButtonUrl                  string                            `json:"button_url,omitempty"`
	Brand                      *BrandStruct                      `json:"brand,omitempty"`
	Caption                    string                            `json:"caption,omitempty"`
	CountdownExpiringTimestamp int64                             `json:"countdown_expiring_timestamp,omitempty"`
}
