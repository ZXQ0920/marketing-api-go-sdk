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
type AdvertiserGetListStruct struct {
	DailyBudget                      int64                        `json:"daily_budget,omitempty"`
	SystemStatus                     CustomerSystemStatus         `json:"system_status,omitempty"`
	CorporationName                  string                       `json:"corporation_name,omitempty"`
	CertificationImageId             string                       `json:"certification_image_id,omitempty"`
	IndividualQualification          *IndividualQualificationRead `json:"individual_qualification,omitempty"`
	IntroductionUrl                  string                       `json:"introduction_url,omitempty"`
	IndustryQualificationImageIdList []string                     `json:"industry_qualification_image_id_list,omitempty"`
	AdQualificationImageIdList       []string                     `json:"ad_qualification_image_id_list,omitempty"`
	ContactPerson                    string                       `json:"contact_person,omitempty"`
	ContactPersonEmail               string                       `json:"contact_person_email,omitempty"`
	ContactPersonTelephone           string                       `json:"contact_person_telephone,omitempty"`
	ContactPersonMobile              string                       `json:"contact_person_mobile,omitempty"`
	WechatSpec                       *MpInfoRead                  `json:"wechat_spec,omitempty"`
	Websites                         []WebsiteReadStruct          `json:"websites,omitempty"`
	AccountId                        int64                        `json:"account_id,omitempty"`
	AdQualificationImage             []string                     `json:"ad_qualification_image,omitempty"`
	CertificationImage               string                       `json:"certification_image,omitempty"`
	IndustryQualificationImage       []string                     `json:"industry_qualification_image,omitempty"`
	CorporateImageName               string                       `json:"corporate_image_name,omitempty"`
	CorporateImageLogo               string                       `json:"corporate_image_logo,omitempty"`
	CorporationLicence               string                       `json:"corporation_licence,omitempty"`
	CustomizedIndustry               string                       `json:"customized_industry,omitempty"`
	IdentityNumber                   string                       `json:"identity_number,omitempty"`
	SystemIndustryId                 int64                        `json:"system_industry_id,omitempty"`
	RejectMessage                    string                       `json:"reject_message,omitempty"`
}
