/*
Copyright © 2024 Donald Gifford dgifford06@gmail.com
*/
package src

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/smartdevicemanagement/v1"
)

type Sdm struct {
	service smartdevicemanagement.Service
	error   error
}

type NestDevice struct {
	device smartdevicemanagement.GoogleHomeEnterpriseSdmV1Device
	error  error
}

func NewSdmService(config AuthConfig, t *oauth2.Token) Sdm {
	httpClient := config.conf.Client(config.c, t)
	sdmService, err := smartdevicemanagement.NewService(config.c, option.WithHTTPClient(httpClient))
	if err != nil {
		fmt.Printf("Error creating sdmService: %v", err)
		return Sdm{service: smartdevicemanagement.Service{}, error: err}
	}
	return Sdm{service: *sdmService, error: nil}
}

func (s *Sdm) GetDevice() NestDevice {
	projectID := viper.GetString("nest.project_id")
	deviceID := viper.GetString("nest.device_id")
	device, err := s.service.Enterprises.Devices.Get(fmt.Sprintf("enterprises/%s/devices/%s", projectID, deviceID)).Do()
	if err != nil {
		fmt.Printf("Error getting device: %v", err)
		return NestDevice{device: smartdevicemanagement.GoogleHomeEnterpriseSdmV1Device{}, error: err}
	}

	return NestDevice{device: *device, error: nil}
}

func (s *Sdm) GetDeviceTraits(nd NestDevice) (Traits, error) {
	var nestTraits Traits
	err := json.Unmarshal(nd.device.Traits, &nestTraits)
	if err != nil {
		fmt.Printf("Error unmarshaling traits: %v", err)
		return Traits{}, err
	}
	return nestTraits, nil
}
