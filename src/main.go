// Copyright 2020 Donovan Solms github.com/donovansolms twitter.com/donovansolms
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"github.com/ApexGamesHQ/apex-games-engine/src/core"
	"github.com/ApexGamesHQ/apex-games-engine/src/extract"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// main executes the service
func main() {

	var err error
	// Set up the default logger
	logLevel := logrus.InfoLevel
	logrus.SetOutput(os.Stdout)

	logOutputFormat := logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "Jan 02 15:04:05",
	}
	logrus.SetFormatter(&logOutputFormat)

	// if config.Debug {
	// 	logLevel = logrus.DebugLevel
	// }

	logger := logrus.WithFields(logrus.Fields{
		"service": "engine",
	})
	logger.Info("Starting the Apex Games Engine")

	// Load configuration
	viper.SetConfigName("apex-games") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			logger.Warn("Config file not found, using defaults")
		} else {
			// Config file was found but another error was produced
			logger.Errorf("Config file found but it container errors: %s", err)
		}
	} else {
		logger.Info("Config loaded successfully")
	}

	logLevel, err = logrus.ParseLevel(viper.GetString("LogLevel"))
	if err != nil {
		logger.Warnf("Log level read from config is invalid: %s", err)
	}
	logrus.SetLevel(logLevel)
	logger.Infof("Log level set to %s", logLevel.String())

	// screenCapture := capture.NewScreen()
	// _ = screenCapture
	// img, err := screenCapture.CaptureFullscreen(1)

	// _ = img
	// _ = err
	// fileName := fmt.Sprintf("test.png")
	// file, _ := os.Create(fileName)
	// defer file.Close()
	// png.Encode(file, img)

	pixelExtractor := extract.NewPerPixel()

	engine := core.NewEngine()

	_ = engine
}
