package kashangwl

func (c *Client) GetCustomerId() int {
	return c.config.customerId
}

func (c *Client) GetCustomerKey() string {
	return c.config.customerKey
}
