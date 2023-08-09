package sdk

type Client struct {
	client *sdk.Request
}

func (c *Client) CreateUser(req *CreateUserRequest) (*CreateUserResponse, error) {
	//normal code
	resp := &CreateUserResponse{}
	err := c.client.Send(req, resp)
	return resp, err
}
