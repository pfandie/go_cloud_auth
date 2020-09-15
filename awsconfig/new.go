/*
Package awsconfig for go_cloud_auth
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
package awsconfig

import (
	"fmt"
	"go_cloud_auth/models"
	"go_cloud_auth/utils"

	"github.com/AlecAivazis/survey/v2"
)

// NewProfile creates new temp profiles to assume AWS roles
func NewProfile() {
	configName := []*survey.Question{
		{
			Name: "ConfigName",
			Prompt: &survey.Input{
				Message: "How should this config be named?",
				Help:    "Name of config for internal storage",
			},
			Validate: survey.Required,
		},
	}
	awsConf := []*survey.Question{
		{
			Name: "AccountDestination",
			Prompt: &survey.Input{
				Message: "Destination/Sub account?",
				Help:    "ID of the sub-account where you want to assume the role",
			},
			Validate: survey.Required,
		},
		{
			Name: "AccountSource",
			Prompt: &survey.Input{
				Message: "Source/Master account?",
				Help:    "ID of the master-account which holds the sub-account",
			},
			Validate: survey.Required,
		},
		{
			Name: "PeronalProfile",
			Prompt: &survey.Input{
				Message: "Personal AWS profile to use?",
				Help:    "AWS profile, which holds your personal AWS credentials/config",
			},
			Validate: survey.Required,
		},
		{
			Name: "AwsProfileName",
			Prompt: &survey.Input{
				Message: "How should this profile be called?",
				Help:    "How the profile name should be named in AWS credentials/config",
			},
			Validate: survey.Required,
		},
		{
			Name: "TokenProfile",
			Prompt: &survey.Input{
				Message: "Which OTP profile to use?",
				Help:    "Your MFA profile to use",
			},
			Validate: survey.Required,
		},
		{
			Name: "Duration",
			Prompt: &survey.Input{
				Message: "How long to assume this role (in seconds)?",
				Default: "3600",
				Help:    "Durataion of assuming the role, defaults to 3600 seconds",
			},
		},
		{
			Name: "AwsUserName",
			Prompt: &survey.Input{
				Message: "What is your AWS username?",
				Help:    "Usually the account name you are using to login to AWS console",
			},
			Validate: survey.Required,
		},
		{
			Name: "RoleName",
			Prompt: &survey.Input{
				Message: "Which role should be assumed?",
				Help:    "The role name you want to assume. e.g. admin or path_to_roles/admin",
			},
			Validate: survey.Required,
		},
		{
			Name: "Region",
			Prompt: &survey.Input{
				Message: "Which region for the profile?",
				Help:    "The region which your AWS profile defaults to",
			},
			Validate: survey.Required,
		},
	}

	userConf := models.AwsUserConfig{
		AccountDestination: "survey:AccountDestination",
		AccountSource:      "survey:AccountSource",
		PeronalProfile:     "survey:PeronalProfile",
		AwsProfileName:     "survey:AwsProfileName",
		TokenProfile:       "survey:TokenProfile",
		Duration:           "survey:Duration",
		AwsUserName:        "survey:AwsUserName",
		RoleName:           "survey:RoleName",
		Region:             "survey:Region",
	}

	configs := models.Configs{
		ConfigName: "survey:ConfigName",
	}
	cname := survey.Ask(configName, &configs)
	err := survey.Ask(awsConf, &userConf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if cname != nil {
		fmt.Println(cname.Error())
		return
	}

	configs.AwsConfig = userConf

	if showConfig(&configs) != true {
		fmt.Println("Aborted, exiting.")
		return
	}
	utils.SaveConfig(&configs)
}

// showConfig displays the entered settings for validation
func showConfig(c *models.Configs) bool {
	fmt.Printf("\nProfile name: %s\n", c.ConfigName)
	fmt.Printf("Account Destination: %s\n", c.AwsConfig.AccountDestination)
	fmt.Printf("Account Source: %s\n", c.AwsConfig.AccountSource)
	fmt.Printf("Personal root profile: %s\n", c.AwsConfig.PeronalProfile)
	fmt.Printf("AWS profile name: %s\n", c.AwsConfig.AwsProfileName)
	fmt.Printf("OTP profile: %s\n", c.AwsConfig.TokenProfile)
	fmt.Printf("Session duration: %s\n", c.AwsConfig.Duration)
	fmt.Printf("AWS username: %s\n", c.AwsConfig.AwsUserName)
	fmt.Printf("AWS role name: %s\n", c.AwsConfig.RoleName)
	fmt.Printf("AWS region: %s\n\n", c.AwsConfig.Region)

	prompt := &survey.Confirm{
		Message: "Are these infos correct?",
		Default: false,
	}
	answer := false

	err := survey.AskOne(prompt, &answer)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}

	return answer
}
