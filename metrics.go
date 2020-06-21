package proxydns

import (
	"sync"

	"github.com/coredns/coredns/plugin"

	"github.com/prometheus/client_golang/prometheus"
)

// requestCount exports a prometheus metric that is incremented every time a query is seen by the example plugin.
var requestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace: plugin.Namespace,
	Subsystem: "proxydns",
	Name:      "request_count_total",
	Help:      "Counter of requests made.",
}, []string{"server"})

var once sync.Once
