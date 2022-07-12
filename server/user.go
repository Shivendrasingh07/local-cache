package server

import (
	"encoding/json"
	"fmt"
	"local-cache/database/helper"
	"local-cache/models"
	"time"

	"net/http"
)

// Login function takes email and password, returns userID
func (srv *Server) Login(resp http.ResponseWriter, req *http.Request) {
	var Cred models.LoginUser
	err := json.NewDecoder(req.Body).Decode(&Cred)
	if err != nil {
		fmt.Println("err")
	}
	loginUser, loginErr := helper.Login(Cred.Email, Cred.Password)
	if loginErr != nil {
		fmt.Println("err")
	}

	err2 := json.NewEncoder(resp).Encode(loginUser)
	if err2 != nil {
		fmt.Println("err")
	}
}

func (srv *Server) Signup(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("err")
	}

	userid, newerr := helper.NewUser(user.Name, user.Email, user.Password)

	if newerr != nil {
		fmt.Println("err")
	}
	err2 := json.NewEncoder(w).Encode(userid)
	if err2 != nil {
		fmt.Println("err")
	}
}

func (srv *Server) LocalCacheMain() {
	var opp models.Operations

	for opp != models.ExitOperations {
		time.Sleep(time.Second * 1)
		fmt.Println("select operations")
		_, err := fmt.Scan(&opp)
		if err != nil {
			panic(err)
			return
		}

		switch opp {
		case models.SetOperations:
			srv.set()
		case models.GetOperations:
			srv.get()
		case models.DeleteOperations:
			srv.deleteKeyValue()
		case models.ExitOperations:
			return
		default:
			return
		}
	}
}

func (srv *Server) set() {
	var key string
	var value string
	var expTime int
	time.Sleep(time.Second * 1)
	fmt.Println("enter [key]")
	_, err := fmt.Scan(&key)
	if err != nil || key == "" {
		fmt.Println("invalid key ")
		return
	}
	fmt.Println("enter [value]")
	_, err = fmt.Scan(&value)
	if err != nil || value == "" {
		fmt.Println("invalid  value")
		return
	}
	fmt.Println("enter [EXP time in seconds]")
	_, err = fmt.Scan(&expTime)
	if err != nil || expTime == 0 {
		fmt.Println("invalid time")
		return
	}

	data := models.KeyValueStruct{Key: key, Value: value, TimeSeconds: expTime}

	err = srv.LocalCache.Set(data)
	if err != nil {
		fmt.Println("unable to set data")
		return
	}

	time.Sleep(time.Second * 1)
	fmt.Println("value and Expiry time added at key successfully")
}

func (srv *Server) get() {
	var key string
	fmt.Println("enter [key]")
	_, err := fmt.Scan(&key)
	if err != nil || key == "" {
		fmt.Println("invalid key ")
		return
	}
	value := srv.LocalCache.Get(key)
	if value == nil {
		fmt.Println("unable to get data")
		return
	}

	fmt.Println("[key]", key)
	fmt.Println("[value]", value)
}

func (srv *Server) deleteKeyValue() {
	var key string
	fmt.Println("enter [key]")
	_, err := fmt.Scan(&key)
	if err != nil && key == "" {
		fmt.Println("invalid key ")
		return
	}

	err = srv.LocalCache.DeleteKeyValue(key)
	if err != nil {
		fmt.Println("unable to delete data")
		return
	}
	fmt.Println("deleted")
}

//func (srv *Server) set(resp http.ResponseWriter, req *http.Request) {
//	//uc := srv.getUserContext(req)
//
//	//var body struct {
//	//	Key   string      `json:"key"`
//	//	Value interface{} `json:"value"`
//	//}
//
//	body := make(map[string]interface{})
//
//	//err := json.NewDecoder(req.Body).Decode(&body)
//	//if err != nil {
//	//	utils.RespondClientErr(resp, req, err, http.StatusBadRequest, "Unable to do this operation right now", "error parsing request")
//	//	return
//	//}
//	//fmt.Println("test InitializeServer")
//
//	var key string
//	var value string
//	fmt.Println("enter [key]")
//	_, err := fmt.Scan(&key)
//	err = srv.LocalCache.set(body)
//	if err != nil {
//		utils.RespondGenericServerErr(resp, req, err, "unable to get candidate's pre fill data for a job")
//		return
//	}
//	fmt.Println("enter [value]")
//	_, err = fmt.Scan(&value)
//	data := map[string]interface{}{
//		key: value,
//	}
//
//	err = srv.LocalCache.set(data)
//	if err != nil {
//		utils.RespondGenericServerErr(resp, req, err, "unable to get candidate's pre fill data for a job")
//		return
//	}
//	//if err != nil {
//	//	utils.RespondGenericServerErr(resp, req, err, "unable to get candidate's pre fill data for a job")
//	//	return
//	//}
//
//	//cache := Newcache()
//
//	utils.EncodeJSONBody(resp, http.StatusCreated, map[string]interface{}{
//		"message": "success",
//	})
//}
//
//func (srv *Server) Get(resp http.ResponseWriter, req *http.Request) {
//	key := chi.URLParam(req, "key")
//	value := srv.LocalCache.Get(key)
//
//	if value == "" {
//		utils.RespondClientErr(resp, req, nil, http.StatusBadRequest, "value at this key is expired", "expired")
//		return
//	}
//
//	utils.EncodeJSONBody(resp, http.StatusOK, map[string]interface{}{
//		"key":   key,
//		"value": value,
//	})
//}
//
//func (srv *Server) Delete(resp http.ResponseWriter, req *http.Request) {
//	key := chi.URLParam(req, "key")
//	srv.LocalCache.Delete(key)
//
//	utils.EncodeJSONBody(resp, http.StatusCreated, map[string]interface{}{
//		"message": "success",
//	})
//}
