/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"context"
	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	. "github.com/tencentad/marketing-api-go-sdk/pkg/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type CustomAudienceInsightsApiService service

/*
CustomAudienceInsightsApiService 人群洞察分析
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId
 * @param audienceId
 * @param dimensionType
 * @param optional nil or *CustomAudienceInsightsGetOpts - Optional Parameters:
     * @param "Fields" (optional.Interface of []string) -  返回参数的字段列表

@return CustomAudienceInsightsGetResponse
*/

type CustomAudienceInsightsGetOpts struct {
	Fields optional.Interface
}

func (a *CustomAudienceInsightsApiService) Get(ctx context.Context, accountId int64, audienceId int64, dimensionType []string, localVarOptionals *CustomAudienceInsightsGetOpts) (CustomAudienceInsightsGetResponseData, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue CustomAudienceInsightsGetResponseData
		localVarResponse    CustomAudienceInsightsGetResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/custom_audience_insights/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("account_id", parameterToString(accountId, ""))
	localVarQueryParams.Add("audience_id", parameterToString(audienceId, ""))
	localVarQueryParams.Add("dimension_type", parameterToString(dimensionType, "multi"))
	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, localVarFileKey)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarResponse, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			if localVarResponse.Code > 0 {
				err = errors.NewError(localVarResponse.Code, localVarResponse.Message, localVarResponse.MessageCn, localVarResponse.Errors)
				return localVarReturnValue, localVarHttpResponse, err
			}
			return *localVarResponse.Data, localVarHttpResponse, err
		} else {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v CustomAudienceInsightsGetResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
