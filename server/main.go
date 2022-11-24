package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	URL      string = "https://api.assemblyai.com/v2/upload"
	key      string = ""
	filePath string
)

func main() {
	getStarted()
}

func getStarted() {
	var file, fileErr = ioutil.ReadFile(filePath)
	if fileErr != nil {
		log.Println("read File issue")
		os.Exit(2)
	}

	var getUrl = getUrl(file)
	var id = getTranscript(getUrl)
	getText(id["id"])
}

func getUrl(file []byte) map[string]string {
	var req, reqErr = http.NewRequest("POST", URL, bytes.NewBuffer(file))
	req.Header.Set("authorization", key)
	if reqErr != nil {
		log.Println("req Error")
		os.Exit(2)
	}

	defer req.Body.Close()
	var client *http.Client = &http.Client{}
	var result, resultErr = client.Do(req)
	if resultErr != nil {
		log.Println("transcript Url Error")
		os.Exit(2)
	}
	var data = map[string]string{}

	json.NewDecoder(result.Body).Decode(&data)
	data["audio_url"] = data["upload_url"]
	delete(data, "upload_url")
	return data
}

func getTranscript(url map[string]string) map[string]string {
	var jsml, _ = json.Marshal(url)
	var req, reqErr = http.NewRequest("POST", "https://api.assemblyai.com/v2/transcript", bytes.NewBuffer(jsml))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", key)
	if reqErr != nil {
		log.Println("Error During transcript req")
		os.Exit(2)
	}
	var client *http.Client = &http.Client{}
	var res, resErr = client.Do(req)
	if resErr != nil {
		log.Println("Error During transcript res")
		os.Exit(2)
	}
	var id = map[string]string{}
	json.NewDecoder(res.Body).Decode(&id)

	return id
}
func getText(id string) {
	time.Sleep(time.Second * 10)
	var req, reqErr = http.NewRequest("GET", "https://api.assemblyai.com/v2/transcript/"+id, nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", key)

	if reqErr != nil {
		log.Println("error During getting result Text")
		os.Exit(2)
	}

	var client *http.Client = &http.Client{}
	var res, resErr = client.Do(req)
	if resErr != nil {
		log.Println("Error During result text")
		os.Exit(2)
	}
	var data = map[string]string{}
	json.NewDecoder(res.Body).Decode(&data)
	if data["text"] != "" {
		log.Println(data["text"])
		os.Exit(1)
	}
	time.Sleep(time.Second * 10)
	getText(id)
}
