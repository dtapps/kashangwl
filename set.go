package kashangwl

type SetConfigConfig struct {
	CustomerId  int    // 商家编号
	CustomerKey string // 商家密钥
}

// SetConfig 配置
func (c *Client) SetConfig(config *SetConfigConfig) *Client {
	c.config.CustomerId = config.CustomerId
	c.config.CustomerKey = config.CustomerKey
	return c
}
