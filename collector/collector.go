package collector

import (
	simpleJson "github.com/bitly/go-simplejson"
	log "github.com/sirupsen/logrus"

	"io/ioutil"
	"net/http"
)

type Job struct{}

type FlumeMetric struct {
	Metrics map[string]map[string]interface{}
}

func (f *FlumeMetric) GetMetrics(flumeMetricUrl string) FlumeMetric {
	httpClient := HttpClient{}
	json, err := httpClient.Get(flumeMetricUrl)
	result := make(map[string]map[string]interface{})
	if err != nil {
		result[flumeMetricUrl] = nil
		return FlumeMetric{result}
	}

	js, err := simpleJson.NewJson([]byte(json))
	if err != nil {
		log.Errorf("simpleJson.NewJson = %v", err)
		result[flumeMetricUrl] = nil
		return FlumeMetric{result}
	}

	flumeMetricMap := make(map[string]interface{})
	flumeMetricMap, _ = js.Map()
	result[flumeMetricUrl] = flumeMetricMap
	return FlumeMetric{result}
}

type HttpClient struct{}

func (httpClient *HttpClient) Get(url string) (string, error) {
	log.Debug(url)

	response, err := http.Get(url)
	if err != nil {
		log.Errorf("httpClient.Get = %v", err)
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
