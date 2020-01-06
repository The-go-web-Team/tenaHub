package service

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/TenaHub/client/entity"
	"fmt"
	"net/url"
	"errors"
)

func FetchHealthCenters() ([]clientEntity.HealthCenter, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s/healthcenters", baseURL)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var healthcenters []clientEntity.HealthCenter
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &healthcenters)
	if err != nil {
		return nil, err
	}
	return healthcenters, nil
}

func FetchHealthCenter(id uint) (*clientEntity.HealthCenter, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("%s/healthcenter/%d", baseURL, id)
	req, _ := http.NewRequest("GET", URL, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var healthcenters *clientEntity.HealthCenter
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &healthcenters)
	if err != nil {
		return nil, err
	}
	return healthcenters, nil
}



// Authenticate authenticates user
func HealthCenterAuthenticate(healthcenter *clientEntity.HealthCenter) (*clientEntity.HealthCenter, error) {
	URL := fmt.Sprintf("%s/%s", baseURL, "healthcenter")

	formval := url.Values{}
	formval.Add("email", healthcenter.Email)
	formval.Add("password", healthcenter.Password)

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
		Content clientEntity.HealthCenter
	}{}

	err = json.Unmarshal(body, &respjson)

	fmt.Println(respjson)

	if respjson.Status == "error" {
		return nil, errors.New("error")
	}
	return &respjson.Content, nil
}





