package kashangwl

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type ApiProductResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Id                      int64  `json:"id"`                        // 商品编号
		ProductName             string `json:"product_name,omitempty"`    // 商品名称
		Name                    string `json:"name"`                      // 规格名称
		Price                   string `json:"price"`                     // 售价
		ValidPurchasingQuantity string `json:"valid_purchasing_quantity"` // 合法的购买数量
		SuperiorCommissionsRate int    `json:"superior_commissions_rate"` // 上级佣金比例
		Type                    int    `json:"type"`                      // 商品类型（1：充值，2：卡密，3：卡券，4：人工）
		SupplyState             int    `json:"supply_state"`              // 库存状态（1：充足，2：断货）
		StockState              int    `json:"stock_state"`               // 状态（1：上架，2：维护，3：下架）
		BanStartAt              string `json:"ban_start_at,omitempty"`    // 禁售开始时间
		BanEndAt                string `json:"ban_end_at,omitempty"`      // 禁售结束时间
	} `json:"data"`
}

type ApiProductResult struct {
	Result ApiProductResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newApiProductResult(result ApiProductResponse, body []byte, http gorequest.Response) *ApiProductResult {
	return &ApiProductResult{Result: result, Body: body, Http: http}
}

// ApiProduct 获取单个商品信息
// product_id = 商品编号
// http://doc.cqmeihu.cn/sales/product-info.html
func (c *Client) ApiProduct(ctx context.Context, productID int64, notMustParams ...gorequest.Params) (*ApiProductResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("product_id", productID) // 商品编号
	// 请求
	request, err := c.request(ctx, apiUrl+"/api/product", params)
	if err != nil {
		return newApiProductResult(ApiProductResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ApiProductResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiProductResult(response, request.ResponseBody, request), err
}
