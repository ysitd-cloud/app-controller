package connect

import (
	"os"

	"github.com/Azure/azure-storage-go"
	"github.com/ysitd-cloud/gin-utils/env"
)

func NewAzureStorageClient() (storage.Client, error) {
	name := os.Getenv("AZURE_STORAGE_NAME")
	key := os.Getenv("AZURE_STORAGE_KEY")

	return storage.NewBasicClient(name, key)
}

func NewAzureTable() (storage.Table, error) {
	client, err := NewAzureStorageClient()
	if err != nil {
		return storage.Table{}, err
	}

	tableService := client.GetTableService()
	name := env.GetEnvWithDefault("AZURE_STORAGE_TABLE", "app")
	return tableService.GetTableReference(name), nil
}
