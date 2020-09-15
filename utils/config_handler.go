/*
Package utils for go_cloud_auth
Copyright © 2020 Hans Mayer <hans.mayer83@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

import (
	"errors"
	"fmt"
	"go_cloud_auth/models"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

var home, _ = homedir.Dir()

//TODO: add custom configfile as func param
var configfile = home + "/.go_cloud_auth.yml"

// SaveConfig stores a yaml Configuration
func SaveConfig(conf *models.Configs) error {
	existingConf, _ := LoadConfig()

	c, err := createConfig(conf, existingConf)

	if err != nil {
		fmt.Println("Profile already exists!")
		return err
	}

	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	fmt.Println("New config saved!")
	return ioutil.WriteFile(configfile, bytes, 0644)
}

// ChangeConfig updates an existing yaml Configuration
func ChangeConfig() {
	//TODO: add change config
}

// DeleteConfig deletes a config from yaml Configuration
func DeleteConfig() {
	//TODO: add delete config
}

// LoadConfig yaml json Configuration
func LoadConfig() (*models.Configuration, error) {
	fileExist, _ := fileExists(configfile)
	if !fileExist {
		return &models.Configuration{}, errors.New("File: configfile could not be read")
	}

	bytes, err := ioutil.ReadFile(configfile)
	if err != nil {
		return &models.Configuration{}, err
	}

	c := models.Configuration{}
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return &models.Configuration{}, err
	}

	return &c, nil
}

// LoadConfigByName configuration by given name
func LoadConfigByName(name string) (*models.AwsUserConfig, error) {
	fileExist, _ := fileExists(configfile)
	if !fileExist {
		return &models.AwsUserConfig{}, errors.New("File: configfile could not be read")
	}

	bytes, err := ioutil.ReadFile(configfile)
	if err != nil {
		return &models.AwsUserConfig{}, err
	}

	c := models.Configuration{}
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return &models.AwsUserConfig{}, err
	}

	var p *models.AwsUserConfig
	for i := 0; i < len(c.Configs); i++ {
		if c.Configs[i].ConfigName == name {
			p = &c.Configs[i].AwsConfig
		}
	}

	return p, nil
}

// GetAllProfileNames gets all profiles from config file
func GetAllProfileNames() ([]string, error) {
	conf, err := LoadConfig()
	if err != nil {
		fmt.Println("An error occurred while reading config file. Please try again -- ", err)
	}

	var profiles []string
	for i := 0; i < len(conf.Configs); i++ {
		profiles = append(profiles, conf.Configs[i].ConfigName)
	}

	if len(profiles) > 0 {
		return profiles, nil
	}

	return nil, errors.New("No profiles found in config file")
}

// fileExists validates if a file exists
func fileExists(file string) (bool, error) {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false, err
	}

	return !info.IsDir(), nil
}

// createConfig prepare Structs to write into file
func createConfig(newConf *models.Configs, conf *models.Configuration) (*models.Configuration, error) {
	// detect if a config with same name already exists
	for i := 0; i < len(conf.Configs); i++ {
		fmt.Println(conf.Configs[i].ConfigName)
		fmt.Println(newConf.ConfigName)
		if conf.Configs[i].ConfigName == newConf.ConfigName {
			return conf, errors.New("A configurtation with this name already exists")
		}
	}

	conf.Configs = append(conf.Configs, models.Configs{
		ConfigName: newConf.ConfigName,
		AwsConfig:  newConf.AwsConfig,
	})

	return conf, nil
}
