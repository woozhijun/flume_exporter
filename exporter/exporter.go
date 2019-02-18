package exporter

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/woozhijun/flume_exporter/config"
	"regexp"
	"strconv"
	"strings"

	"github.com/woozhijun/flume_exporter/collector"
)

type Exporter struct {
	gaugeVecs  map[string]*prometheus.GaugeVec
	configFile string
}

func NewExporter(namespace string, configFile string, metric *config.Metrics) *Exporter {
	gaugeVecs := make(map[string]*prometheus.GaugeVec)
	for k, v := range metric.Metrics {
		for _, m := range v {
			val := fmt.Sprintf("%s_%s", k, m)
			gaugeVecs[val] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      val,
				Help:      val},
				[]string{"host", "type", "name"})
		}
	}
	return &Exporter{
		gaugeVecs:  gaugeVecs,
		configFile: configFile,
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	for _, value := range e.gaugeVecs {
		value.Describe(ch)
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	log.Debugf("%v", e)

	e.collectGaugeVec()
	for _, value := range e.gaugeVecs {
		value.Collect(ch)
	}
}

func (e *Exporter) collectGaugeVec() error {

	f := collector.FlumeMetric{}
	var flumeMetricUrls []string

	conf := config.GetConfig(e.configFile)
	if conf == nil {
		return errors.New("load flume config failed")
	}
	for _, agent := range conf.Agents {
		if agent.Enabled {
			for _, url := range agent.Urls {
				flumeMetricUrls = append(flumeMetricUrls, url)
			}
		}
	}
	// flumeMetricUrls := e.flumeMetricUrls
	log.Debugf("flumeMetricUrls=%v", flumeMetricUrls)

	channel := make(chan collector.FlumeMetric)
	for _, url := range flumeMetricUrls {
		go func(url string) {
			channel <- f.GetMetrics(url)
		}(url)
	}

	// receive from all channels
	for i := 0; i < len(flumeMetricUrls); i++ {
		m := <-channel
		url := flumeMetricUrls[i]
		if m.Metrics == nil {
			log.Warn(">>>.receive metrics channel is nil, url: " + url)
			continue
		}
		reg := regexp.MustCompile(`//(.*)/metrics`)
		host := reg.FindStringSubmatch(url)[1]
		for k, v := range m.Metrics {
			sMetrics := make(map[string]interface{})
			sMetrics = v.(map[string]interface{})
			delete(sMetrics, "Type")

			if strings.HasPrefix(k, "SOURCE.") {
				e.processGaugeVecs(k, host, "SOURCE", sMetrics)
			} else if strings.HasPrefix(k, "CHANNEL.") {
				delete(sMetrics, "Open")
				e.processGaugeVecs(k, host, "CHANNEL", sMetrics)
			} else if strings.HasPrefix(k, "SINK.") {
				e.processGaugeVecs(k, host, "SINK", sMetrics)
			}
		}
	}
	return nil
}

func (e *Exporter) processGaugeVecs(title string, host string, flumeType string, data map[string]interface{}) {

	name := strings.Replace(title, flumeType+".", "", 1)
	for mName, mValue := range data {
		val, err := strconv.ParseFloat(mValue.(string), 64)
		if err != nil {
			log.Errorf("value = %v", val)
			val = 0
		}
		gv := e.gaugeVecs[flumeType+"_"+mName]
		if gv != nil {
			gv.WithLabelValues(host, flumeType, name).Set(val)
		} else {
			fmt.Printf("====> metric: %s, type %s", flumeType, mName)
			fmt.Println()
		}
	}
}
