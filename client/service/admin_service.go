package service

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/TenaHub/client/entity"
	"fmt"
	"net/url"
	"errors"
	"time"
)
var baseURL = "http://localhost:8181/v1"

func FetchAdmin(id int) (*clientEntity.Admin, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s/admin/%d", baseURL, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	adminData := clientEntity.Admin{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &adminData)
	fmt.Println("error is ",err)
	if err != nil {
		return nil, err
	}
	return &adminData, nil
}




type cookie struct {
	Key        string
	Expiration time.Time
}

type response struct {
	Status string
	Content interface{}
}

var loggedIn = make([]cookie, 10)

func getResponse(request *http.Request) []byte {
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	if 200 != resp.StatusCode {
		panic(errors.New("status not correct"))
	}
	return body
}


// Authenticate authenticates user
func AdminAuthenticate(admin *clientEntity.Admin) (*clientEntity.Admin, error) {
	URL := fmt.Sprintf("%s/%s", baseURL, "admin")

	formval := url.Values{}
	formval.Add("email", admin.Email)
	formval.Add("password", admin.Password)

	resp, err := http.PostForm(URL, formval)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	respjson := struct {
		Status string
		Content clientEntity.Admin
	}{}

	err = json.Unmarshal(body, &respjson)

	fmt.Println(respjson)

	if respjson.Status == "error" {
		return nil, errors.New("error")
	}
	return &respjson.Content, nil
}


//
//func FetchHealthCenter(id int) (*clientEntity.HealthCenter, error) {
//	client := &http.Client{}
//	URL := fmt.Sprintf("%s/healthcenter/%d", baseURL, id)
//	req, _ := http.NewRequest("GET", URL, nil)
//	res, err := client.Do(req)
//	if err != nil {
//		return nil, err
//	}
//	healthcenter := &clientEntity.HealthCenter{}
//	body, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		return nil, err
//	}
//	err = json.Unmarshal(body, healthcenter)
//	if err != nil {
//		return nil, err
//	}
//	return healthcenter, nil
//}
//
//func FetchUser(id int) (*clientEntity.User, error) {
//	client := &http.Client{}
//	URL := fmt.Sprintf("%s/user/%d", baseURL, id)
//	req, _ := http.NewRequest("GET", URL, nil)
//	res, err := client.Do(req)
//	if err != nil {
//		return nil, err
//	}
//	userdata := &clientEntity.User{}
//	body, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		return nil, err
//	}
//	err = json.Unmarshal(body, userdata)
//	if err != nil {
//		return nil, err
//	}
//	return userdata, nil
//}
//
