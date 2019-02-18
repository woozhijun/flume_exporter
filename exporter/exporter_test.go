package exporter

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/woozhijun/flume_exporter/config"
	"regexp"
	"testing"
)

func TestNewExporter(t *testing.T) {

	Convey("exporter testing", t, func() {

		metrics := config.GetCollectMetrics("../metrics.yml")

		exporter := NewExporter("flume", "config.yml", metrics)
		fmt.Println(exporter)

		url := "http://localhost:53454/metrics"
		reg := regexp.MustCompile(`//(.*)/metrics`)
		fmt.Println(reg.FindStringSubmatch(url)[1])
	})
}
