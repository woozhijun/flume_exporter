package config

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	Agents []Agent
}

type Agent struct {
	Name string
	Enabled bool
	Urls []string
}

func GetConfig(configFile string) *Conf  {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("could not read config.yml file; err: <%s>", err)
		return nil
	}
	conf := Conf{}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil
	}
	return &conf
}

type Metrics struct {
	Metrics		map[string][]string
}

type CollectMetrics struct {
	Sources  	[]string        `sources`
	Channels    []string        `channels`
	Sinks		[]string		`sinks`
}

func GetCollectMetrics(configFile string) *Metrics  {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("could not read metrics.yml file; err: <%s>", err)
		return nil
	}
	conf := CollectMetrics{}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil
	}
	result := make(map[string][]string)
	result["SOURCE"] = conf.Sources
	result["CHANNEL"] = conf.Channels
	result["SINK"] = conf.Sinks
	return &Metrics{result}
}
