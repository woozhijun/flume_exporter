package collector

import (
	log "github.com/Sirupsen/logrus"
	simpleJson "github.com/bitly/go-simplejson"

	"io/ioutil"
	"net/http"
)

type Job struct{}

type FlumeMetric struct {
	Metrics map[string]interface{}
}

func (f *FlumeMetric) GetMetrics(flumeMetricUrl string) FlumeMetric {
	httpClient := HttpClient{}
	json, err := httpClient.Get(flumeMetricUrl)
	if err != nil {
		log.Errorf("HttpClient.Get = %v", err)
		return FlumeMetric{nil}
	}

	js, err := simpleJson.NewJson([]byte(json))
	if err != nil {
		log.Errorf("simpleJson.NewJson = %v", err)
		return FlumeMetric{nil}
	}

	flumeMetricMap := make(map[string]interface{})
	flumeMetricMap, _ = js.Map()
	return FlumeMetric{flumeMetricMap}
}

type HttpClient struct{}

func (httpClient *HttpClient) Get(url string) (string, error) {
	log.Debug(url)

	response, err := http.Get(url)
	if err != nil {
		log.Errorf("http.Get = %v", err)
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf("ioutil.ReadAll = %v", err)
		return "", err
	}

	if response.StatusCode != 200 {
		log.Errorf("response.StatusCode = %v", response.StatusCode)
		return "", err
	}

	json := string(body)
	log.Debug(json)
	return json, nil
}
