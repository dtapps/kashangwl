package kashangwl

import (
	"context"
	"encoding/json"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
)

type ApiProductRechargeParamsCacheResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		PurchaseTips string `json:"purchase_tips"`
	} `json:"data"`
	Time    int    `json:"time"`
	TraceId string `json:"trace_id"`
}

type ApiProductRechargeParamsCacheResult struct {
	Result ApiProductRechargeParamsCacheResponse // 结果
	Body   []byte                                // 内容
	Http   gorequest.Response                    // 请求
	Err    error                                 // 错误
}

func newApiProductRechargeParamsCacheResult(result ApiProductRechargeParamsCacheResponse, body []byte, http gorequest.Response, err error) *ApiProductRechargeParamsCacheResult {
	return &ApiProductRechargeParamsCacheResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiProductRechargeParamsCache 接口说明
// 获取商品的充值参数（仅支持充值类商品）
// http://doc.cqmeihu.cn/sales/ProductParams.html
func (c *Client) ApiProductRechargeParamsCache(ctx context.Context, productId int64) *ApiProductRechargeParamsCacheResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.requestCache(ctx, fmt.Sprintf("%s/%d/goods_info_params/%d", apiUrlCache, c.GetCustomerId(), productId), params, http.MethodGet)
	// 定义
	var response ApiProductRechargeParamsCacheResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newApiProductRechargeParamsCacheResult(response, request.ResponseBody, request, err)
}
