package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/woozhijun/flume_exporter/exporter"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
)

const (
	namespace = "FLUME"
	clientID  = "flume_exporter"
	versionID = "1.0.0"
)

func main() {

	var (
		configFile       = kingpin.Flag("config-file", "Set config file").Default("config.yml").String()
		listeningAddress = kingpin.Flag("listen.address", "The app listen address.").Default(":9360").String()
		metricEndpoint   = kingpin.Flag("metric.endpiont", "The app listen endpiont.").Default("/metrics").String()
		logLevel         = kingpin.Flag("log-level", "Set Logging level").Default("info").String()
		metricFile       = kingpin.Flag("metric-file", "Set metrics file").Default("metrics.yml").String()
	)
	//plog.AddFlags(kingpin.CommandLine)
	kingpin.Version(version.Print(clientID + " " + versionID))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	setupLogging(*logLevel)
	exporter := exporter.NewExporter(namespace, *configFile, *metricFile)
	prometheus.MustRegister(exporter)

	http.Handle(*metricEndpoint, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`<html>
	        <head><title>flume Exporter</title></head>
	        <body>
	        <h1>flume Exporter</h1>
	        <p><a href='` + *metricEndpoint + `'>Metrics</a></p>
	        </body>
	        </html>`))
	})

	if err := http.ListenAndServe(*listeningAddress, nil); err != nil {
		log.Fatal(err)
	}
}

func setupLogging(logLevel string) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("could not set log level to '%s';err:<%s>", logLevel, err)
	}
	log.SetLevel(level)
}
