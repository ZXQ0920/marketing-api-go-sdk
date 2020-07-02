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

type EcommerceOrderGetExample struct {
	TAds                  *ads.SDKClient
	AccessToken           string
	AccountId             int64
	DateRange             model.DateRange
	EcommerceOrderGetOpts *api.EcommerceOrderGetOpts
}

func (e *EcommerceOrderGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.DateRange = model.DateRange{
		StartDate: "REPORT START DATE",
		EndDate:   "REPORT END DATE",
	}
	e.EcommerceOrderGetOpts = &api.EcommerceOrderGetOpts{

		Fields: optional.NewInterface([]string{"account_id", "ecommerce_order_id", "customized_page_name", "commodity_package_detail", "quantity", "price", "total_price", "ecommerce_order_time", "ecommerce_order_status", "user_name", "user_phone", "user_province", "user_city", "user_area", "user_address", "user_ip", "user_message", "destination_url", "adgroup_id", "adgroup_name", "from_account_id", "delivery_spec"}),
	}
}

func (e *EcommerceOrderGetExample) RunExample() (model.EcommerceOrderGetResponseData, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.EcommerceOrder().Get(ctx, e.AccountId, e.DateRange, e.EcommerceOrderGetOpts)
}

func main() {
	e := &EcommerceOrderGetExample{}
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
