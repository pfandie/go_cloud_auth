/*
Package models for go_cloud_auth
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
package models

// AwsUserConfig struct containing Userdata and Account Information to assume role
type AwsUserConfig struct {
	// TODO: implement validation
	AccountDestination string "yaml:\"account_destination\"" // must have exact 12 digits
	AccountSource      string "yaml:\"account_source\""      // must have exact 12 digits 826934983476
	PeronalProfile     string "yaml:\"personal_profile\""    // appgile_root
	AwsProfileName     string "yaml:\"aws_profile\""         // appgile_master
	TokenProfile       string "yaml:\"token_profile\""       // otp - appgile
	Duration           string "yaml:\"duration\""            // 36000
	AwsUserName        string "yaml:\"aws_username\""        // hans.mayer
	RoleName           string "yaml:\"role_name\""           // appgile/appgile_master
	Region             string "yaml:\"region\""              // eu-central-1
}
