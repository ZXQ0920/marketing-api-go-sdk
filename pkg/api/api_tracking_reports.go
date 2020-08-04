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

type TrackingReportsApiService service

/*
TrackingReportsApiService 点击追踪报表
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId
 * @param dateRange
 * @param optional nil or *TrackingReportsGetOpts - Optional Parameters:
     * @param "TimeGranularity" (optional.String) -
     * @param "PromotedObjectType" (optional.String) -
     * @param "PromotedObjectId" (optional.String) -
     * @param "FeedbackUrl" (optional.String) -
     * @param "Page" (optional.Int64) -
     * @param "PageSize" (optional.Int64) -
     * @param "Fields" (optional.Interface of []string) -  返回参数的字段列表

@return TrackingReportsGetResponse
*/

type TrackingReportsGetOpts struct {
	TimeGranularity    optional.String
	PromotedObjectType optional.String
	PromotedObjectId   optional.String
	FeedbackUrl        optional.String
	Page               optional.Int64
	PageSize           optional.Int64
	Fields             optional.Interface
}

func (a *TrackingReportsApiService) Get(ctx context.Context, accountId int64, dateRange DateRange, localVarOptionals *TrackingReportsGetOpts) (TrackingReportsGetResponseData, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue TrackingReportsGetResponseData
		localVarResponse    TrackingReportsGetResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/tracking_reports/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("account_id", parameterToString(accountId, ""))
	localVarQueryParams.Add("date_range", parameterToString(dateRange, ""))
	if localVarOptionals != nil && localVarOptionals.TimeGranularity.IsSet() {
		localVarQueryParams.Add("time_granularity", parameterToString(localVarOptionals.TimeGranularity.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PromotedObjectType.IsSet() {
		localVarQueryParams.Add("promoted_object_type", parameterToString(localVarOptionals.PromotedObjectType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PromotedObjectId.IsSet() {
		localVarQueryParams.Add("promoted_object_id", parameterToString(localVarOptionals.PromotedObjectId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FeedbackUrl.IsSet() {
		localVarQueryParams.Add("feedback_url", parameterToString(localVarOptionals.FeedbackUrl.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
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
			var v TrackingReportsGetResponse
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
