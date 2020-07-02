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

type ProductCatalogsGetExample struct {
	TAds                   *ads.SDKClient
	AccessToken            string
	AccountId              int64
	ProductCatalogsGetOpts *api.ProductCatalogsGetOpts
}

func (e *ProductCatalogsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.ProductCatalogsGetOpts = &api.ProductCatalogsGetOpts{

		Fields: optional.NewInterface([]string{"product_catalog_id", "product_catalog_name", "product_catalog_type", "product_catalog_vertical", "product_catalog_status", "product_recommend_methods", "deep_link_enabled"}),
	}
}

func (e *ProductCatalogsGetExample) RunExample() (model.ProductCatalogsGetResponseData, *http.Response, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.ProductCatalogs().Get(ctx, e.AccountId, e.ProductCatalogsGetOpts)
}

func main() {
	e := &ProductCatalogsGetExample{}
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
