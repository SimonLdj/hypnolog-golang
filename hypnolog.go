package hypnolog

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	defaltDomain = "localhost"
	//defaltDomain = "host.docker.internal"
	defaultPort = "7000"
)

func HypnoLogStruct(data interface{}) {
	jsonStr, _ := json.Marshal(data)
	HypnoLog(string(jsonStr), "json")
}

func HypnoLogString(str string) {
	HypnoLog(str, "string")
}


func HypnoLog(data string, typeName string) {
	url := "http://"+defaltDomain+":"+defaultPort+"/logger/in"
	//fmt.Println("URL:>", url)

	msg := HypnologMessage{
		Data: data,
		Type: typeName,
	}

	jsonStr, _ := json.Marshal(msg)

	//var jsonStr = []byte(`{"type":"json", "data":"Buy cheese and bread for breakfast."}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	//if err != nil {
	//panic(err)
	//}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
}

type HypnologMessage struct {
	Data string   `json:"data"`
	Type string   `json:"type"`
	Tags []string `json:"tags"`
}
