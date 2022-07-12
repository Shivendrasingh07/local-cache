package models

import "time"

type LocalCacheStruct struct {
	LocalCacheData map[string]interface{}
	ExpiryTime     map[string]interface{}
}

type ExpiryTime struct {
	Key  string
	Time time.Time
}

type KeyValueStruct struct {
	Key         string
	Value       string
	TimeSeconds int
}
