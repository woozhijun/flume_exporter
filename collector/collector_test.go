package collector

import (
	log "github.com/Sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCollector(t *testing.T) {

	url := "http://localhost:8080/metrics"
	Convey("metrics collector tests", t, func() {

		flumeMetric, err := GetMetrics(url)

		log.Info(flumeMetric)
		So(err, ShouldBeNil)
		So(flumeMetric, ShouldNotBeNil)
	})
}