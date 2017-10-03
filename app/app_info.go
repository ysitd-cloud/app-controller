package app

type MetaInformation interface {
	GetImage() string
	GetTag() string
}

func NewMetaInformation(image, tag string) MetaInformation {
	return &appInfo{
		image: image,
		tag:   tag,
	}
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
