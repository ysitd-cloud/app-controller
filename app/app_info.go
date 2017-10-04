package app

func NewMetaInformation(image, tag string) MetaInformation {
	return &metaInformation{
		image: image,
		tag:   tag,
	}
}

func (info *metaInformation) GetImage() string {
	return info.image
}

func (info *metaInformation) GetTag() string {
	return info.tag
}
