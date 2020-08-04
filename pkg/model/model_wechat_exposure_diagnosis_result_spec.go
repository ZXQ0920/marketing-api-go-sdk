/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 微信广告的曝光诊断结论
type WechatExposureDiagnosisResultSpec struct {
	DiagnosisConclusionSpec            *DiagnosisConclusionSpec            `json:"diagnosis_conclusion_spec,omitempty"`
	TargetingDiagnosisConclusionSpec   *TargetingDiagnosisConclusionSpec   `json:"targeting_diagnosis_conclusion_spec,omitempty"`
	CostDiagnosisConclusionSpec        *CostDiagnosisConclusionSpec        `json:"cost_diagnosis_conclusion_spec,omitempty"`
	CompititionDiagnosisConclusionSpec *CompititionDiagnosisConclusionSpec `json:"compitition_diagnosis_conclusion_spec,omitempty"`
}
