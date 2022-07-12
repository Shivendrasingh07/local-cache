package main

import (
	"local-cache/server"
)

var LocalCacheData map[string]interface{}

func main() {
	srv := server.InitializeServer()

	srv.LocalCacheMain()

	//srv := server.InitializeServer()
	//fmt.Println("test InitializeServer")
	//srv.Start()
	//
	////	srv.LocalCache()
	//err := database.Connect()
	////	fmt.Println("test database")
	//if err != nil {
	//	panic(err)
	//}
	//	localCache()
}

//func (a *Server) start() {
//	a.initializeRoutes()
//}

//func localCache() {
//	var opp models.Operations
//
//	for opp != models.ExitOperations {
//		time.Sleep(time.Second * 1)
//		fmt.Println("select operations")
//		_, err := fmt.Scan(&opp)
//		if err != nil {
//			panic(err)
//			return
//		}
//
//		switch opp {
//		case models.SetOperations:
//			set()
//		case models.GetOperations:
//			get()
//		case models.DeleteOperations:
//			deleteKeyValue()
//		case models.ExitOperations:
//			return
//		default:
//			return
//		}
//
//	}
//}
//
//func set() {
//	var key string
//	var value string
//	time.Sleep(time.Second * 2)
//	fmt.Println("enter [key]")
//	_, err := fmt.Scan(&key)
//
//	if err != nil {
//		panic(err)
//		return
//	}
//	fmt.Println("enter [value]")
//	_, err = fmt.Scan(&value)
//	data := map[string]interface{}{
//		key: value,
//	}
//	LocalCacheData = data
//	time.Sleep(time.Second * 1)
//	fmt.Println("value added at key successfully")
//}
//
//func get() {
//	var key string
//	fmt.Println("enter [key]")
//	_, err := fmt.Scan(&key)
//	if err != nil {
//		panic(err)
//		return
//	}
//	fmt.Println("[value]", LocalCacheData[key])
//
//}
//
//func deleteKeyValue() {
//	var key string
//	fmt.Println("enter [key]")
//	_, err := fmt.Scan(&key)
//	if err != nil {
//		panic(err)
//		return
//	}
//	LocalCacheData[key] = nil
//	time.Sleep(time.Second * 1)
//	fmt.Println("deleted")
//}
