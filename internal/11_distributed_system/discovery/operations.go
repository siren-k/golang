package discovery

import "github.com/hashicorp/consul/api"

// Register 함수는 consul에 서비스를 추가한다.
func (c *client) Register(tags []string) error {
	reg := &api.AgentServiceRegistration{
		ID:      c.name,
		Name:    c.name,
		Port:    c.port,
		Address: c.address,
		Tags:    tags,
	}
	return c.client.Agent().ServiceRegister(reg)
}

// Service 함수는 서비스를 반환한다.
func (c *client) Service(service, tag string) ([]*api.ServiceEntry, *api.QueryMeta, error) {
	return c.client.Health().Service(service, tag, false, nil)
}
