package goRazerChromaSDK

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	SKD_Url string = "http://localhost:54235/razer/chromasdk"
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

func GetSession(app AppInfo) (sessionId, error) {

	//app := AppInfo{}
	//log.Println("appAuthorName: ", app.Author.Name)

	reqBody, err := json.Marshal(app)
	//println("reqbody:", string(reqBody), err)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(SKD_Url, "application/json", bytes.NewBuffer(reqBody))
	//fmt.Println("postBody:", resp.Body)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	//fmt.Println("stringBody ", string(body))
	curSession := sessionId{}
	err = json.Unmarshal(body, &curSession)
	if err != nil {
		return sessionId{}, err
	}
	//fmt.Println("stringBody2 ", string(body))
	return curSession, nil

}

func GetVersion() (version, error) {

	response, err := http.Get(SKD_Url)

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
	//fmt.Println(string(responseData))
	return curVersion, nil

}
