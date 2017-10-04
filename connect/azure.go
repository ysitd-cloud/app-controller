package connect

import (
	"os"

	"github.com/Azure/azure-storage-go"
	"github.com/ysitd-cloud/gin-utils/env"
)

var storageClient storage.Client
var storageInitial bool = false

func NewAzureStorageClient() (client storage.Client, err error) {
	if storageInitial {
		client = storageClient
		err = nil
		return
	}

	name := os.Getenv("AZURE_STORAGE_NAME")
	key := os.Getenv("AZURE_STORAGE_KEY")

	client, err = storage.NewBasicClient(name, key)

	storageClient = client
	storageInitial = true

	return
}

var tableServiceClient storage.TableServiceClient
var tableServiceBootstrap = false

func NewAzureTableServiceClient() (client storage.TableServiceClient, err error) {
	if tableServiceBootstrap {
		client = tableServiceClient
		err = nil
		return
	}

	storageClient, err := NewAzureStorageClient()
	client = storageClient.GetTableService()

	tableServiceClient = client
	tableServiceBootstrap = true
	return
}

func NewAzureTable() (storage.Table, error) {
	tableService, err := NewAzureTableServiceClient()
	if err != nil {
		return storage.Table{}, err
	}

	name := env.GetEnvWithDefault("AZURE_STORAGE_TABLE", "app")
	return tableService.GetTableReference(name), nil
}
