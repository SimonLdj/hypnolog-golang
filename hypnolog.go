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

func Create() Message {
	return Message{}
}

func Tag(tagName string) Message {
	return Message{ Tags: []string{tagName}}
}

// todo: consider to remove
func LogStruct(data interface{}) bool {
	return Create().Set(data).Log()
}

// todo: consider to remove
func LogString(str string) bool{
	return Create().Str(str).Log()
}

func Log(message Message) bool {

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

	if resp.StatusCode != 200 {
		fmt.Println("Hypnolog error, unexpected status code from server: ", resp.StatusCode)
		return false

	}
	return true
}

type Message struct {
	Data interface{} `json:"data"`
	Type string      `json:"type"`
	Tags []string    `json:"tags"`
}

func (message Message) Tag(tagName string) Message {
	message.Tags = append(message.Tags, tagName)
	return message
}

func (message Message) Str(str string) Message {
	if message.Data == nil {
		message.Data = str
	} else {
		message.Data = fmt.Sprintln(message.Data) + str
	}
	message.Type = "string"
	return message
}

func (message Message) Set(data interface{}) Message {
	message.Data = data
	message.Type = "json"
	return message
}

func (message Message) Log() bool {
	return Log(message)
}
