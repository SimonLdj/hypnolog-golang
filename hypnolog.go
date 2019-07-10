package hypnolog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	defaltDomain = "localhost"
	defaultPort = "7000"
)

var (
	host  = defaltDomain
)

func SetHost(newHost string) {
	host = newHost
}

func LogStruct(data interface{}) bool {
	return Log(HypnologMessage{Data: data, Type: "json"})
}

func LogString(str string) bool{
	return Log(HypnologMessage{Data: str, Type: "string"})
}

func Log(message HypnologMessage) bool {

	//fmt.Println("URL:>", postUrl)
	postUrl := "http://" + host + ":" + defaultPort + "/logger/in"

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(message)
	if err != nil {
		fmt.Println("Hypnolog error, while serializing json request: ", err)
		return false
	}
	resp , err := http.Post(postUrl, "application/json; charset=utf-8", b)
	if err != nil {
		fmt.Println("Hypnolog error, while sending http request: ", err)
		return false
	}

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	if resp.StatusCode != 200 {
		fmt.Println("Hypnolog error, unexpected status code from server: ", resp.StatusCode)
		return false

	}
	return true
}

type HypnologMessage struct {
	Data interface{} `json:"data"`
	Type string      `json:"type"`
	Tags []string    `json:"tags"`
}
