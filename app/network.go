package app

func NewNetwork(domain string, port int32) Network {
	return &network{
		domain: domain,
		port:   port,
	}
}

func (d *network) GetDomain() string {
	return d.domain
}

func (d *network) GetPort() int32 {
	return d.port
}
