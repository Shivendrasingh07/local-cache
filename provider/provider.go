package provider

import "local-cache/models"

type LocalCacheInterface interface {

	//  Get(key string) string
	//  Set(data map[string]interface{}) error
	//  Delete(key string)

	Get(key string) interface{}
	Set(data models.KeyValueStruct) error
	DeleteKeyValue(key string) error
}

type ConfigProvider interface {
	GetString(key string) string
	GetServerPort() string
}
