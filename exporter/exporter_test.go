package exporter

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"regexp"
	"testing"
)

func TestNewExporter(t *testing.T) {

	Convey("exporter testing", t, func() {

		exporter := NewExporter("flume", "../config.yml", "../metrics.yml")
		fmt.Println(exporter)

		url := "http://localhost:53454/metrics"
		reg := regexp.MustCompile(`//(.*)/metrics`)
		fmt.Println(reg.FindStringSubmatch(url)[1])
	})
}
