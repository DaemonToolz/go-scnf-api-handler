package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// curl -X GET "network-and-schedule/typical-flight-schedule/iata-seasons" -H "Api-Key:****"
const (
	NetworkSchedule = "network-and-schedule"
)

func sendTrainLinesRequest() TrainData {
	data := TrainData{}
	json.Unmarshal(sendRequest("GET", "https://api.sncf.com/v1/coverage/sncf/stop_areas/stop_area:SNCF:87271494/lines/", nil, nil), &data)
	return data
}

func sendRequest(method string, uri string, inBody interface{}, headers map[string]string) []byte {
	client := &http.Client{}

	var req *http.Request
	if inBody != nil {
		bodyByte, _ := json.Marshal(inBody)
		req, _ = http.NewRequest(method, uri, bytes.NewBuffer(bodyByte))
	} else {
		req, _ = http.NewRequest(method, uri, nil)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", appConfig.Key)

	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}
