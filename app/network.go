package app

type Network interface {
	GetDomain() string
	GetPort() int32
}

func NewNetwork(domain string, port int32) Network {
	return &discoverNetwork{
		domain: domain,
		port:   port,
	}
}

type discoverNetwork struct {
	domain string
	port   int32
}

func (d *discoverNetwork) GetDomain() string {
	return d.domain
}

func (d *discoverNetwork) GetPort() int32 {
	return d.port
}
