/*
Copyright © 2024 Donald Gifford dgifford06@gmail.com
*/
package src

type DeviceTraits struct {
	traits Traits
	error  error
}

type Traits struct {
	Connectivity struct {
		Status string `json:"status"`
	} `json:"sdm.devices.traits.Connectivity"`
	Fan struct {
		TimerMode string `json:"timerMode"`
	} `json:"sdm.devices.traits.Fan"`
	Humidity struct {
		AmbientHumidityPercent int `json:"ambientHumidityPercent"`
	} `json:"sdm.devices.traits.Humidity"`
	Info struct {
		CustomName string `json:"customName"`
	} `json:"sdm.devices.traits.Info"`
	Settings struct {
		TemperatureScale string `json:"temperatureScale"`
	} `json:"sdm.devices.traits.Settings"`
	Temperature struct {
		AmbientTemperatureCelsius float64 `json:"ambientTemperatureCelsius"`
	} `json:"sdm.devices.traits.Temperature"`
	ThermostatEco struct {
		AvailableModes []string `json:"availableModes"`
		Mode           string   `json:"mode"`
		CoolCelsius    float64  `json:"coolCelsius"`
		HeatCelsius    float64  `json:"heatCelsius"`
	} `json:"sdm.devices.traits.ThermostatEco"`
	ThermostatHvac struct {
		Status string `json:"status"`
	} `json:"sdm.devices.traits.ThermostatHvac"`
	ThermostatMode struct {
		AvailableModes []string `json:"availableModes"`
		Mode           string   `json:"mode"`
	} `json:"sdm.devices.traits.ThermostatMode"`
	ThermostatTemperatureSetpoint struct {
		HeatCelsius float64 `json:"heatCelsius"`
		CoolCelsius float64 `json:"coolCelsius"`
	} `json:"sdm.devices.traits.ThermostatTemperatureSetpoint"`
}
