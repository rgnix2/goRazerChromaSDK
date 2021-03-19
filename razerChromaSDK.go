package goRazerChromaSDK

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	sdkUrl string = "http://localhost:54235/razer/chromasdk"
)

type version struct {
	Core    string `json:"core"`
	Device  string `json:"device"`
	Version string `json:"version"`
}

type Author struct {
	Name    string `json:"name"`
	Contact string `json:"contact"`
}

type AppInfo struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	Author          `json:"author"`
	DeviceSupported []string `json:"device_supported"`
	Category        string   `json:"category"`
}

type sessionId struct {
	Sessionid int    `json:"sessionid"`
	URI       string `json:"uri"`
}

func GetSession(AppInfo) (sessionId, error) {
	//url := "http://localhost:54235/razer/chromasdk"

	response, err := http.Get(sdkUrl)
	if err != nil {
		log.Fatal(err)

	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	curSession := sessionId{}
	err = json.Unmarshal(responseData, &curSession)
	if err != nil {
		return sessionId{}, err
	}
	fmt.Println(string(responseData))
	return curSession, nil

}

func GetVersion() (version, error) {

	//url := "http://localhost:54235/razer/chromasdk"

	response, err := http.Get(sdkUrl)

	if err != nil {
		log.Fatal(err)

	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	curVersion := version{}
	err = json.Unmarshal(responseData, &curVersion)
	if err != nil {
		return version{}, err
	}
	fmt.Println(string(responseData))
	return curVersion, nil

}
