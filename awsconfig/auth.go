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
	"log"
	"os/exec"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

// Authenticate uses existing profiles to assume AWS roles
func Authenticate() {
	profiles, err := utils.GetAllProfileNames()
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	profile := []*survey.Question{
		{
			Name: "profile",
			Prompt: &survey.Select{
				Message: "Select profile to use:",
				Options: profiles,
			},
		},
	}

	p := struct {
		Profile string "survey:\"profile\""
	}{}

	surv := survey.Ask(profile, &p)
	if surv != nil {
		log.Fatalln(surv.Error())
	}

	fmt.Printf("Profile selected: %v\n", p.Profile)
	c, err := utils.LoadConfigByName(p.Profile)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	// token, err := getAwsSession(c)
	token := utils.GetAuthToken(c.TokenProfile)

	creds, err := initializeAwsSession(c, token)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	setProfileSimple(creds, c.AwsProfileName, c.Region)
	duration, _ := strconv.Atoi(c.Duration)
	duration = duration / 60 / 60

	fmt.Printf("Session successfully set. Session is valid for %d hours.\n", duration)
}

// setProfileSimple simple function to set profile in aws/credentials file, has to be changed
func setProfileSimple(c *credentials.Value, profile string, region string) {
	cmdArgs := []string{
		"configure",
		"--profile",
		profile,
		"set",
		"aws_access_key_id",
		c.AccessKeyID,
	}

	cmd := exec.Command("aws", cmdArgs...)
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("Set successful aws_access_key_id for profile %s\n", profile)

	cmdArgs[4] = "aws_secret_access_key"
	cmdArgs[5] = c.SecretAccessKey
	cmd = exec.Command("aws", cmdArgs...)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("Set successful aws_secret_access_key for profile %s\n", profile)

	cmdArgs[4] = "aws_session_token"
	cmdArgs[5] = c.SessionToken
	cmd = exec.Command("aws", cmdArgs...)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("Set successful aws_session_token for profile %s\n", profile)

	cmdArgs[4] = "region"
	cmdArgs[5] = region
	cmd = exec.Command("aws", cmdArgs...)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("Set successful region %s for profile %s\n", region, profile)
}

// initializeAwsSession gets credentials for temp session
func initializeAwsSession(c *models.AwsUserConfig, token string) (*credentials.Value, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: c.PersonalProfile,
		Config: aws.Config{
			Region: aws.String(c.Region),
		},
	}))
	svc := sts.New(sess)

	duration, _ := strconv.ParseInt(c.Duration, 10, 64)
	params := &sts.AssumeRoleInput{
		DurationSeconds: aws.Int64(duration),
		RoleArn:         aws.String(fmt.Sprintf("arn:aws:iam::%v:role/%v", c.AccountDestination, c.RoleName)),
		RoleSessionName: aws.String("aws_session"),
		TokenCode:       aws.String(token),
		SerialNumber:    aws.String(fmt.Sprintf("arn:aws:iam::%v:mfa/%v", c.AccountSource, c.AwsUserName)),
	}

	response, err := svc.AssumeRole(params)
	if err != nil {
		return nil, err
	}

	awsCredentials := credentials.Value{
		AccessKeyID:     *response.Credentials.AccessKeyId,
		SecretAccessKey: *response.Credentials.SecretAccessKey,
		SessionToken:    *response.Credentials.SessionToken,
	}

	return &awsCredentials, nil
}
