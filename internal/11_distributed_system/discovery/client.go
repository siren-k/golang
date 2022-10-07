package discovery

import "github.com/hashicorp/consul/api"

// Client는 우리가 관심이 있는 api 메소드를 노출시킨다.
type Client interface {
	Register(tags []string) error
	Service(service, tag string) ([]*api.ServiceEntry, *api.QueryMeta, error)
}

type client struct {
	client  *api.Client
	address string
	name    string
	port    int
}

// NewClient 함수는 consul 클라이언트를 초기화한다.
func NewClient(config *api.Config, address, name string, port int) (Client, error) {
	c, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	cli := &client{
		client:  c,
		name:    name,
		address: address,
		port:    port,
	}
	return cli, nil
}
