package providers

import (
	"os"

	"github.com/Azure/azure-storage-go"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/gin-utils/env"
)

type azureServiceProvider struct {
	*container.AbstractServiceProvider
}

func NewAzureServiceProvider(kernel container.Container) container.ServiceProvider {
	sp := azureServiceProvider{
		AbstractServiceProvider: container.NewAbstractServiceProvider(true),
	}

	sp.SetContainer(kernel)

	return &sp
}

func (sp *azureServiceProvider) Provides() []string {
	return []string{
		"azure.storage.account.name",
		"azure.storage.account.key",
		"azure.storage.client",
		"azure.storage.client.table",
		"azure.storage.table.app.name",
		"azure.storage.table.app",
	}
}

func (sp *azureServiceProvider) Register(kernel container.Container) {
	sp.registerClient(kernel)
	sp.registerTableClient(kernel)
	sp.registerTable(kernel)
}

func (sp *azureServiceProvider) registerClient(kernel container.Container) {
	kernel.Instance("azure.storage.account.name", os.Getenv("AZURE_STORAGE_NAME"))
	kernel.Instance("azure.storage.account.key", os.Getenv("AZURE_STORAGE_KEY"))

	kernel.Singleton("azure.storage.client", func(kernel container.Container) interface{} {
		name := kernel.Make("azure.storage.account.name").(string)
		key := kernel.Make("azure.storage.account.key").(string)
		client, err := storage.NewBasicClient(name, key)
		if err != nil {
			panic(err)
		}
		return client
	})
}

func (sp *azureServiceProvider) registerTableClient(kernel container.Container) {
	kernel.Singleton("azure.storage.client.table", func(kernel container.Container) interface{} {
		client := kernel.Make("azure.storage.client").(storage.Client)
		return client.GetTableService()
	})
}

func (sp *azureServiceProvider) registerTable(kernel container.Container) {
	kernel.Instance("azure.storage.table.app.name", env.GetEnvWithDefault("AZURE_STORAGE_TABLE", "app"))

	kernel.Bind("azure.storage.table.app", func(kernel container.Container) interface{} {
		client := kernel.Make("azure.storage.client.table").(storage.TableServiceClient)
		return client.GetTableReference(kernel.Make("azure.storage.table.app.name").(string))
	})
}
