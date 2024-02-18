package src

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

func Daemon(ctx context.Context) error {
	defaultTickConfig := viper.GetInt64("daemon.default_tick")
	defaultTick := time.Duration(defaultTickConfig) * time.Minute
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.Tick(defaultTick):
			ctx := context.Background()
			config := InitConfig(ctx)
			token := Authenticate(config)
			sdm := NewSdmService(config, token)
			device := sdm.GetDevice()
			traits, err := sdm.GetDeviceTraits(device)
			if err != nil {
				fmt.Println("Error: ", err)
				log.Fatal(err)
			}
			influxClient := CreateInfluxDBClient()
			influxClient.WritePoint(traits)

			fmt.Println("woo")
		}
	}
}
