package kashangwl

import (
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	CustomerId  int    // 商家编号
	CustomerKey string // 商家密钥
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		customerId  int    // 商家编号
		customerKey string // 商家密钥
	}
	log struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.customerId = config.CustomerId
	c.config.customerKey = config.CustomerKey

	c.requestClient = gorequest.NewHttp()

	return c, nil
}
