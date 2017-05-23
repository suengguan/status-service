package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	_ "app-service/status-service/routers"
	"model"
)

const (
	base_url = "http://localhost:8080/v1/status"
)

func Test_JobStatus_GetAll(t *testing.T) {
	// create user and resource
	daoApi.UserDaoApi.Init("http://user-dao-service:8080")
	daoApi.ResourceDaoApi.Init("http://resource-dao-service:8080")
	var user model.User
	var resource model.Resource
	resource.Id = 0
	user.Id = 0
	user.Name = "user"
	user.EncryptedPassword = "user"
	user.Resource = &resource
	newUser, err := daoApi.UserDaoApi.Create(&user)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("user:", *newUser)
	resource.User = newUser
	newResource, err := daoApi.ResourceDaoApi.Create(&resource)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("resource:", *newResource)

	// get all job status
	res, err := http.Get(base_url + "/job/1")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}
