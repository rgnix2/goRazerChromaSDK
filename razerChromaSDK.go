package goRazerChromaSDK

// NOTES
// use convention commits and something to create change log using them.
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var sessionUrl = ""

const (
	api                string = "https://chromasdk.io:54236/razer/chromasdk"
	apiSession         string = "ads"
	MaxIdleConnections int    = 20
	RequestTimeout     int    = 5
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

type Results struct {
	ID     string `json:"id"`
	Result int    `json:"result"`
}

type StaticColor struct {
	Effect string `json:"effect"`
	Param  struct {
		Color int `json:"color"`
	} `json:"param"`
}

func KeyboardStaticPUT(app AppInfo, color StaticColor) (Results, error) {
	println("pgkURL: ", sessionUrl)
	reqBody, err := json.Marshal(color)
	println("reqbody:", string(reqBody))
	if err != nil {
		panic(err)
	}
	println("\napi: ", sessionUrl+"/keyboard")
	///////////////////////////////////
	// initialize http client
	client := &http.Client{}

	// set the HTTP method, url, and request body
	req, err := http.NewRequest("PUT", sessionUrl+"/keyboard", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		panic(err)
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//////////////
	fmt.Println(resp.StatusCode, resp.Request.URL)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println("stringBody ", string(body))
	curResults := Results{}
	err = json.Unmarshal(body, &curResults)
	if err != nil {
		return Results{}, err
	}

	return curResults, nil
}
func MouseStatic(app AppInfo, color StaticColor) (Results, error) {

	println("pgkURL: ", sessionUrl)
	reqBody, err := json.Marshal(color)
	println("reqbody:", string(reqBody))
	if err != nil {
		panic(err)
	}
	println("\napi: ", sessionUrl+"/mousepad")
	///////////////////////////////////
	// initialize http client
	client := &http.Client{}

	// set the HTTP method, url, and request body
	req, err := http.NewRequest("PUT", sessionUrl+"/mousepad", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		panic(err)
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//////////////
	fmt.Println(resp.StatusCode, resp.Request.URL)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println("stringBody ", string(body))
	curResults := Results{}
	err = json.Unmarshal(body, &curResults)
	if err != nil {
		return Results{}, err
	}

	return curResults, nil
}

// http post for API is not working. Seems put is the only way forward.
func KeyboardStatic(app AppInfo, color StaticColor) (Results, error) {

	println("pgkURL: ", sessionUrl)
	reqBody, err := json.Marshal(color)
	println("reqbody:", string(reqBody))
	if err != nil {
		panic(err)
	}

	println("\napi: ", sessionUrl+"/keyboard")
	resp, err := http.Post(sessionUrl+"/keyboard", "application/json", bytes.NewBuffer(reqBody))
	//fmt.Println("postBody:", resp.Body)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println("stringBody ", string(body))
	curResults := Results{}
	err = json.Unmarshal(body, &curResults)
	if err != nil {
		return Results{}, err
	}

	return curResults, nil
}

func GetSession(app AppInfo) (sessionId, error) {

	reqBody, err := json.Marshal(app)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(api, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	curSession := sessionId{}
	err = json.Unmarshal(body, &curSession)
	if err != nil {
		return sessionId{}, err
	}
	sessionUrl = curSession.URI

	print("\npkgSession ", sessionUrl+"\n")

	time.Sleep(2 * time.Second)
	return curSession, nil

}

func GetVersion() (version, error) {

	response, err := http.Get(api)

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

	return curVersion, nil

}
