module github.com/woozhijun/flume_exporter

go 1.13

require (
	github.com/alecthomas/units v0.0.0-20201120081800-1786d5ef83d4 // indirect
	github.com/bitly/go-simplejson v0.5.0
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/prometheus/client_golang v1.9.0
	github.com/prometheus/common v0.15.0
	github.com/prometheus/procfs v0.3.0 // indirect
	github.com/prometheus/promu v0.7.0 // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/smartystreets/goconvey v1.6.4
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b // indirect
	golang.org/x/oauth2 v0.0.0-20210113205817-d3ed898aa8a3 // indirect
	golang.org/x/sys v0.0.0-20210113181707-4bcb84eeeb78 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad => github.com/golang/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7 => github.com/golang/net v0.0.0-20190813141303-74dc4d7220e7
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae => github.com/golang/sys v0.0.0-20200625212154-ddb9806d33ae
)
