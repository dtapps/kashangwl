package kashangwl

import (
	"go.dtapp.net/golog"
)

// ClientConfig 实例配置
type ClientConfig struct {
	CustomerId  int64  // 商家编号
	CustomerKey string // 商家密钥
}

// Client 实例
type Client struct {
	config struct {
		customerId  int64  // 商家编号
		customerKey string // 商家密钥
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.customerId = config.CustomerId
	c.config.customerKey = config.CustomerKey

	return c, nil
}
