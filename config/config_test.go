package config

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetCollectMetrics(t *testing.T) {
	Convey("init metrics test", t, func() {

		conf := GetCollectMetrics("../metrics.yml")
		fmt.Println(conf)
		So(conf, ShouldNotBeNil)
	})
}
