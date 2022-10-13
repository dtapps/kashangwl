package kashangwl

import "go.dtapp.net/golog"

func (c *Client) GetCustomerId() int64 {
	return c.config.customerId
}

func (c *Client) GetCustomerKey() string {
	return c.config.customerKey
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
