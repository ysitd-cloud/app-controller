package app

type Application interface {
	GetID() string
	GetEnvironment() Environment
	GetMeta() MetaInformation
	GetAutoScale() AutoScale
	GetNetwork() Network
}

type app struct {
	id        string
	info      MetaInformation
	env       Environment
	autoScale AutoScale
	network   Network
}

type MetaInformation interface {
	GetImage() string
	GetTag() string
}

type metaInformation struct {
	image string
	tag   string
}

type AutoScale interface {
	GetReplicas() int32
}

type autoScale struct {
	replicas int32
}

type Environment map[string]string

type Network interface {
	GetDomain() string
	GetPort() int32
}

type network struct {
	domain string
	port   int32
}
