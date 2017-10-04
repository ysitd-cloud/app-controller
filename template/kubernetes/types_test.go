package kubernetes

import (
	"github.com/ysitd-cloud/app-controller/app"
)

func createTestApplication() app.Application {
	env := make(app.Environment)
	env["usage"] = "test"
	meta := app.NewMetaInformation("golang", "1.9-alpine")
	autoScale := app.NewAutoScale(32)
	network := app.NewNetwork("test.app.ysitd.cloud", 12345)

	return app.NewApplication("testing", env, meta, autoScale, network)
}
