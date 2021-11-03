package main

import (
	"fmt"
	"github.com/nradchenko/mp707/usb"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	namespace = "mp707"
)

var (
	sensorTemperatureDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "sensor", "temp_celsius"),
		"Sensor temperature",
		[]string{"address", "description"}, nil)
)

type Exporter struct {
	descriptions map[string]string
}

func NewExporter(descriptions map[string]string) *Exporter {
	return &Exporter{descriptions: descriptions}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- sensorTemperatureDesc
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	log.Debug("processing collect request")
	defer log.Debug("done processing collect request")

	devices, err := usb.Lookup()
	if err != nil {
		log.WithError(err).Error("error enumerating devices")
		return
	}

	for _, device := range devices {
		defer device.Close()

		log.WithField("id", device.GetId()).Debug("processing device")

		if roms, err := device.GetSensors(); err != nil {
			log.WithError(err).Error("cannot get sensors list")
		} else {
			for _, rom := range roms {
				address := fmt.Sprintf("%s", rom)
				log.WithField("address", address).Debug("processing sensor")

				if temperature, err := device.GetTemperature(rom); err != nil {
					log.WithField("address", address).Error("cannot get sensor reading")
				} else {
					description, ok := e.descriptions[address]
					if !ok {
						log.WithField("address", address).Debugf("no description for sensor")
					}

					ch <- prometheus.MustNewConstMetric(sensorTemperatureDesc, prometheus.GaugeValue, float64(temperature), address, description)
				}
			}
		}
	}
}
