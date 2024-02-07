/*
Copyright © 2024 Donald Gifford dgifford06@gmail.com
*/
package src

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type InfluxDBClientConfig struct {
	token  string
	url    string
	org    string
	bucket string
}

type InfluxDBClient struct {
	client influxdb2.Client
	config InfluxDBClientConfig
}

func CreateInfluxDBClient() InfluxDBClient {
	config := InfluxDBClientConfig{
		token:  viper.GetString("influxdb.token"),
		url:    viper.GetString("influxdb.url"),
		org:    viper.GetString("influxdb.org"),
		bucket: viper.GetString("influxdb.bucket"),
	}

	client := influxdb2.NewClient(config.url, config.token)
	return InfluxDBClient{client: client, config: config}
}

type NestMetric struct {
	measurement string
	tags        map[string]string
	fields      map[string]interface{}
	timestamp   time.Time
}

func convertCelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func convertStatusToBool(s string) bool {
	if s == "ONLINE" {
		return true
	} else {
		return false
	}
}

func (i *InfluxDBClient) WritePoint(t Traits) {
	writeAPI := i.client.WriteAPIBlocking(i.config.org, i.config.bucket)
	p := converTraitToPoint(t)
	if err := writeAPI.WritePoint(context.Background(), p); err != nil {
		log.Fatal(err)
	}
}

func converTraitToPoint(t Traits) *write.Point {
	nm := NestMetric{
		measurement: "ThermostatMetrics",
		tags: map[string]string{
			"device": "thermostat",
			"type":   "nest",
		},
		fields: map[string]interface{}{
			"AmbientTempetureCelcius":       convertCelsiusToFahrenheit(t.Temperature.AmbientTemperatureCelsius),
			"ThermostatTemperatureSetpoint": convertCelsiusToFahrenheit(t.ThermostatTemperatureSetpoint.HeatCelsius),
			"AmbientHumidityPercent":        t.Humidity.AmbientHumidityPercent,
			"Status":                        convertStatusToBool(t.Connectivity.Status),
		},
		timestamp: time.Now(),
	}

	return influxdb2.NewPoint(nm.measurement, nm.tags, nm.fields, nm.timestamp)
}
