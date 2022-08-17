package kashangwl

import (
	"context"
	"encoding/json"
	"fmt"
	"go.dtapp.net/gorequest"
	"net/http"
)

type ApiProductCacheResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		GoodsId            string  `json:"goods_id"`
		ApiGoodsId         int64   `json:"api_goods_id"`
		GoodsName          string  `json:"goods_name"`
		ClassificationName string  `json:"classification_name"`
		GoodsPrice         float64 `json:"goods_price"`
		PurchaseTips       string  `json:"purchase_tips"`
	} `json:"data"`
	Time    int    `json:"time"`
	TraceId string `json:"trace_id"`
}

type ApiProductCacheResult struct {
	Result ApiProductCacheResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
	Err    error                   // 错误
}

func newApiProductCacheResult(result ApiProductCacheResponse, body []byte, http gorequest.Response, err error) *ApiProductCacheResult {
	return &ApiProductCacheResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiProductCache [缓存，需托管授权]获取单个商品信息
func (c *Client) ApiProductCache(ctx context.Context, productId int64) *ApiProductCacheResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.requestCache(ctx, fmt.Sprintf("%s/%d/goods_info/%d", apiUrlCache, c.GetCustomerId(), productId), params, http.MethodGet)
	// 定义
	var response ApiProductCacheResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newApiProductCacheResult(response, request.ResponseBody, request, err)
}
