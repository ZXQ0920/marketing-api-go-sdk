/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type AdgroupsGetExample struct {
	TAds            *ads.SDKClient
	AccessToken     string
	AccountId       int64
	AdgroupsGetOpts *api.AdgroupsGetOpts
}

func (e *AdgroupsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.AdgroupsGetOpts = &api.AdgroupsGetOpts{

		Filtering: optional.NewInterface([]model.FilteringStruct{model.FilteringStruct{
			Field:    "promoted_object_type",
			Operator: "EQUALS",
			Values:   []string{"PROMOTED_OBJECT_TYPE_APP_IOS"},
		}}),

		Fields: optional.NewInterface([]string{"adgroup_id", "campaign_id", "adgroup_name"}),
	}
}

func (e *AdgroupsGetExample) RunExample() (model.AdgroupsGetResponseData, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.Adgroups().Get(ctx, e.AccountId, e.AdgroupsGetOpts)
}

func main() {
	e := &AdgroupsGetExample{}
	e.Init()
	response, httpResponse, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Http response:", httpResponse)
}
