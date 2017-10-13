package core

import "github.com/ysitd-cloud/app-controller/app"

func NewAutoScaleManager() AutoScaleManager {
	return &noAutoScale{}
}

func (*noAutoScale) GetEntry(id string) app.AutoScale {
	return app.NewAutoScale(int32(1))
}

func (*noAutoScale) Close() {}
