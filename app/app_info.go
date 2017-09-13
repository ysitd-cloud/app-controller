package app

type MetaInformation interface {
	GetImage() string
	GetTag() string
}

type appInfo struct {
	image string
	tag   string
}

func (info *appInfo) GetImage() string {
	return info.image
}

func (info *appInfo) GetTag() string {
	return info.tag
}
