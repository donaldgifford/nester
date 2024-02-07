/*
Copyright © 2024 Donald Gifford dgifford06@gmail.com
*/

package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/donaldgifford/nester/src"
	"github.com/spf13/cobra"
)

var printMetrics bool

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs nester",
	Long: `A single run command to pull data from the configured nest and push it to the configured influxDB server.

Typicall this will run in a crontab to pull data every 5 minutes.
'*/5 * * * * /path/to/nester run'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
		// create authconfig
		ctx := context.Background()
		config := src.InitConfig(ctx)
		token := src.Authenticate(config)
		sdm := src.NewSdmService(config, token)
		device := sdm.GetDevice()
		traits, err := sdm.GetDeviceTraits(device)
		if err != nil {
			fmt.Println("Error: ", err)
			log.Fatal(err)
		}
		if printMetrics {
			// fmt.Println(traits.Connectivity)
			fmt.Printf("Connectivity: %v\n", traits.Connectivity.Status)
			fmt.Printf("TimerMode: %v\n", traits.Fan.TimerMode)
			fmt.Printf("AmbientHumidityPercent: %v\n", traits.Humidity.AmbientHumidityPercent)
			fmt.Printf("CustomName: %v\n", traits.Info.CustomName)
			fmt.Printf("TemperatureScale: %v\n", traits.Settings.TemperatureScale)
			fmt.Printf("AmbientTemperatureCelsius: %v\n", traits.Temperature.AmbientTemperatureCelsius)
			fmt.Printf("HeatCelsius: %v\n", traits.ThermostatEco.HeatCelsius)
			fmt.Printf("CoolCelsius: %v\n", traits.ThermostatEco.CoolCelsius)
			fmt.Printf("HeatCelsiusSetPoint: %v\n", traits.ThermostatTemperatureSetpoint.HeatCelsius)
		} else {
			fmt.Println("no traits printed")
		}
		// src.Inf()
		influxClient := src.CreateInfluxDBClient()
		influxClient.WritePoint(traits)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolVarP(&printMetrics, "print", "p", false, "toggle to print trait metrics to stdout")
}
