/*
Package cmd for go_cloud_auth
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
package cmd

import (
	"fmt"
	"go_cloud_auth/awsconfig"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go_cloud_auth",
	Short: "A go based CLI wrapper to assume an AWS session for interact with different AWS accounts",
	Long: `GO_Cloud Auth is a CLI wrapper to assume specific AWS roles for any provided subaccount.
Configure your Account IDs, MFA profiles and role names.

The information will be stored in your OS-KeyChain and used to assume different AWS-roles.
GO_Cloud Auth will generate different profiles in your ~/.aws/credentials (custom credentials file will be added soon)`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// TODO: implement cli authentication (cobra style)
			awsInit(cmd)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go_cloud_auth.json)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// awsInit prettyfied user input
func awsInit(cmd *cobra.Command) {
	prompt := promptui.Select{
		Label: "Please select",
		Items: []string{
			"Authenticate with existing profile",
			"Update existing profile",
			"Add new profile",
			"MFA Token",
			"Show help",
			"Exit",
		},
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	switch i {
	case 0:
		awsconfig.Authenticate()
		break
	case 1:
		awsconfig.UpdateProfile()
		break
	case 2:
		awsconfig.NewProfile()
		break
	case 3:
		awsconfig.MfaToken()
		os.Exit(0)
		break
	case 4:
		cmd.Help()
		os.Exit(0)
		break
	case 5:
		os.Exit(0)
		break
	default:
		cmd.Help()
		os.Exit(0)
		break
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// TODO: implement config load
	// 	if cfgFile != "" {
	// 		// Use config file from the flag.
	// 		viper.SetConfigFile(cfgFile)
	// 	} else {
	// 		// Find home directory.
	// 		home, err := homedir.Dir()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			os.Exit(1)
	// 		}

	// 		// Search config in home directory with name ".go_cloud_auth.yml".
	// 		viper.AddConfigPath(home)
	// 		viper.SetConfigName(".go_cloud_auth.go_cloud_auth")
	// 	}

	// 	viper.AutomaticEnv() // read in environment variables that match

	// 	// If a config file is found, read it in.
	// 	// if err := viper.ReadInConfig(); err == nil {
	// 	// 	fmt.Println("Using config file:", viper.ConfigFileUsed())
	// 	// }
}
