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
	"go_cloud_auth/utils"
	"os"

	"github.com/manifoldco/promptui"
)

// MfaToken uses existing profiles to assume AWS roles
func MfaToken() {
	prompt := promptui.Select{
		Label: "Please select",
		Items: []string{
			"Show single Token",
			"New / Change Token",
			"Delete Token",
			// TODO: list existing tokens
		},
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	switch i {
	case 0:
		// show
		utils.ShowToken()
		break
	case 1:
		// new / change
		utils.SetSerial()
		break
	case 2:
		// delete
		utils.DeleteSerial()
		break
	default:
		os.Exit(0)
		break
	}
}
