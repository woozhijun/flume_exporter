package exporter

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"github.com/woozhijun/flume_exporter/collector"
	"github.com/woozhijun/flume_exporter/config"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type Exporter struct {
	gaugeVecs       map[string]*prometheus.GaugeVec
	flumeMetricUrls []string
}

func NewExporter(namespace string, configFile string, metricFile string) *Exporter {

	metrics := config.GetCollectMetrics(metricFile)
	if metrics == nil {
		log.Fatal("load metrics.yml failed.")
		log.Exit(2)
	}
	gaugeVecs := make(map[string]*prometheus.GaugeVec)
	for k, v := range metrics.Metrics {
		for _, m := range v {
			val := fmt.Sprintf("%s_%s", k, m)
			gaugeVecs[val] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      val,
				Help:      val},
				[]string{"host", "type", "name"})
		}
	}

	var flumeUrls []string
	conf := config.GetConfig(configFile)
	if conf == nil {
		log.Fatal("load flume config.yml failed.")
		log.Exit(2)
	}
	for _, agent := range conf.Agents {
		if agent.Enabled {
			for _, url := range agent.Urls {
				flumeUrls = append(flumeUrls, url)
			}
		}
	}
	log.Debugf("flumeUrls=%v", flumeUrls)

	return &Exporter{
		gaugeVecs:       gaugeVecs,
		flumeMetricUrls: flumeUrls,
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

func (e *Exporter) collectGaugeVec() bool {

	var wg = sync.WaitGroup{}
	f := collector.FlumeMetric{}
	channel := make(chan collector.FlumeMetric)
	wg.Add(2)
	go func(metricUrls []string) {

		defer wg.Done()
		for _, url := range metricUrls {
			channel <- f.GetMetrics(url)
		}
	}(e.flumeMetricUrls)

	go func() {
		defer wg.Done()
		for _, url := range e.flumeMetricUrls {

			m := <-channel
			if m.Metrics[url] == nil {
				log.Warn(">>>.receive metrics channel is nil, url: " + url)
				continue
			}
			reg := regexp.MustCompile(`//(.*)/metrics`)
			host := reg.FindStringSubmatch(url)[1]
			for k, v := range m.Metrics[url] {
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
	}()
	return true
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
		if gv == nil {
			// filter metrics: StartTime StopTime
			continue
		} else {
			gv.WithLabelValues(host, flumeType, name).Set(val)
		}
	}
}
