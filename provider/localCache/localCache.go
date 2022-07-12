package localCache

import (
	"context"
	"fmt"
	"local-cache/models"
	"local-cache/provider"
	"time"
)

var LocalCachedData models.LocalCacheStruct

type CacheStruct struct {
	ctx context.Context
}

func NewLocalCacheProvider(confPath string) provider.LocalCacheInterface {

	//err := localcache.Init(confPath)
	//if err != nil {
	//	logrus.Fatalf("NewStorageProvider : %v", err)
	//}
	return &CacheStruct{}
}

func (lc CacheStruct) Set(data models.KeyValueStruct) error {

	tempData := map[string]interface{}{
		data.Key: data.Value,
	}
	timeData := time.Now().Add(time.Duration(data.TimeSeconds))
	tempTimeData := map[string]interface{}{
		data.Key: timeData,
	}
	LocalCachedData.LocalCacheData = tempData
	LocalCachedData.ExpiryTime = tempTimeData

	go func() {
		time.Sleep(time.Second * time.Duration(data.TimeSeconds))
		err := lc.DeleteKeyValue(data.Key)
		if err != nil {
			fmt.Println("unable to delete data")
			return
		}
	}()

	return nil
}

func (lc CacheStruct) Get(key string) interface{} {
	return LocalCachedData.LocalCacheData[key]
}

func (lc CacheStruct) DeleteKeyValue(key string) error {
	LocalCachedData.LocalCacheData[key] = nil
	time.Sleep(time.Second * 1)
	return nil
}

//  func (lc CacheStruct) set(data map[string]interface{}) error {
//	var newKey string
//	newKey = (data["key"]).(string)
//	newValue := data["value"].(string)
//
//	err := localcache.set(newKey, []byte(newValue), time.Second*10)
//	if err != nil {
//		logrus.Fatalf("NewStorageProvider : %v", err)
//		return err
//	}
//	return nil
//  }

//func (lc CacheStruct) set(data map[string]interface{}) error {
//	//var newKey string
//	//newKey = (data["key"]).(string)
//	//newValue := data["value"].(string)
//
//	LocalData = data
//	fmt.Println(LocalData)
//	//lc.data["key"] = newKey
//	//lc.data["value"] = newValue
//
//	return nil
//}
//
//func (lc CacheStruct) Get(key string) string {
//	//	time.Sleep(time.Second * 30)
//	//	return localcache.GetString(key)
//	//var data map[string]interface{}
//	//data =
//
//	return LocalData[key].(string)
//
//}
//
//func (lc CacheStruct) Delete(key string) {
//	//	time.Sleep(time.Second * 30)
//	//localcache.Del(key)
//	LocalData[key] = nil
//}
