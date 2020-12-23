package main

import (
	"github.com/nradchenko/mp707/usb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
)

var (
	Version = "dev"

	listenAddress     = kingpin.Flag("listen-address", "Address to listen on").Short('l').Envar("LISTEN_ADDRESS").Default(":9774").TCP()
	sensorDescription = kingpin.Flag("sensors-description", "Sensor description (e.g. 1800000c911b9228='Room #1')").Short('d').Envar("SENSOR_DESCRIPTION").StringMap()

	verbose = kingpin.Flag("verbose", "Verbose logging").Short('v').Envar("VERBOSE").Bool()
	_       = kingpin.HelpFlag.Short('h')
)

func main() {
	kingpin.Version(Version)
	kingpin.Parse()

	if *verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	if err := usb.InitLib(); err != nil {
		log.Fatal(err)
	}

	defer usb.DesposeLib()

	e := NewExporter(*sensorDescription)
	prometheus.MustRegister(e)
	http.Handle("/metrics", promhttp.Handler())
	log.Info("starting exporter on ", *listenAddress)
	log.Fatal(http.ListenAndServe((*listenAddress).String(), nil))
}
