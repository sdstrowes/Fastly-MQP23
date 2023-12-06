package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	madParamDefault   = 10
	maxBucketDefault  = 20
	windowSizeDefault = 360
)

type SubscriptionMsg struct {
	Host   string `json:"host,omitempty"` //aka collector
	Peer   string `json:"peer,omitempty"`
	Asn    int    `json:"asn,omitempty"`
	Prefix string `json:"prefix,omitempty"`
}

type Configuration struct {
	FileInputOption string            `json:"dataOption"`      //required
	StaticFile      string            `json:"staticFilePath"`  //required for static
	Subscriptions   []SubscriptionMsg `json:"subscriptions"`   //required for live
	MadParameters   int               `json:"madParameters"`   //optional - set to default if omitted
	MaxBuckets      int               `json:"maxBuckets"`      //optional - set to default if omitted
	WindowSize      int               `json:"windowSize"`      //optional - set to default if omitted
	URLStaticData   string            `json:"staticFilesLink"` //only required for use in get_static_data.go
}

// load config json file
func LoadConfig(filename string) (*Configuration, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Configuration

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// checks if each value in the config file is valid input
func (c *Configuration) ValidateConfiguration() error {

	//convert all strings to lower case to ignore any capitalizations
	fileInputL := strings.ToLower(c.FileInputOption)

	//checks for valid input corresponding to choice of live vs static data analysis
	if fileInputL == "live" {
		//require at least 1 subscription
		if len(c.Subscriptions) == 0 {
			return errors.New("choosing live data input stream requires to input at least one subscription")
		}
	} else if fileInputL == "static" {
		//require valid file path
		_, err := os.Stat(c.StaticFile)
		if err != nil {
			if os.IsNotExist(err) {
				return errors.New("Invalid pathway to static file: " + err.Error())
			} else {
				return errors.New("Error validating static file path: " + err.Error())
			}
		}
	} else {
		//return error if input is neither live or static
		return errors.New("DataOption in config.json must be either 'live' or 'static'")
	}

	//require mad parameter
	//set mad param to default value if no val set in config
	if c.MadParameters <= 0 || c.MadParameters > 1000 {
		c.MadParameters = madParamDefault
		fmt.Println("No valid mad parameter given. The default mad parameter was set to 10")
	}

	//require maxBuckets
	//set maxBuckets to default value if not set in config
	if c.MaxBuckets <= 0 {
		c.MaxBuckets = maxBucketDefault
		fmt.Println("No valid maxBuckets value given. The default maxBuckets was set to 20")
	}

	//added to check if no window size was put in
	if c.WindowSize <= 0 {
		c.WindowSize = windowSizeDefault
		fmt.Println("No valid window size given. The default window size was set to 360")
	}

	//check if maxBuckets * windowSize is not greater than 1k - could lead to messy graph
	if (c.MaxBuckets * c.WindowSize) >= 10000 {
		//return error
		return errors.New("WindowSize and/or maxBuckets too large - plot may crash")
	}

	//no errors found - valid config file
	fmt.Println("Configuration successful")
	return nil
}
