package collector

import (
	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCollector(t *testing.T) {

	url := "http://localhost:8080/metrics"
	Convey("metrics collector tests", t, func() {

		flumeMetric := FlumeMetric{}
		flumeMetric = flumeMetric.GetMetrics(url)

		log.Info(flumeMetric)
		So(flumeMetric, ShouldNotBeNil)
	})
}
