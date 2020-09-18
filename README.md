# GO Cloud Auth

[![Build Status](https://api.travis-ci.org/pfandie/go_cloud_auth.svg?branch=master)](https://travis-ci.org/pfandie/go_cloud_auth)
[![Go Report Card](https://goreportcard.com/badge/github.com/pfandie/go_cloud_auth)](https://goreportcard.com/report/github.com/pfandie/go_cloud_auth)
[![Coverage Status](https://coveralls.io/repos/github/pfandie/go_cloud_auth/badge.svg?branch=master)](https://coveralls.io/github/pfandie/go_cloud_auth?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg)](https://pkg.go.dev/github.com/pfandie/go_cloud_auth)
[![MIT License](https://img.shields.io/badge/license-Apache-blueviolet.svg)][license]
[![GitHub release](https://img.shields.io/github/release/pfandie/go_cloud_auth.svg)][release]
[![Maintainability](https://api.codeclimate.com/v1/badges/3f963565edd4aa310d27/maintainability)][maintain]

[release]: https://github.com/pfandie/go_cloud_auth/releases
[license]: https://github.com/pfandie/go_cloud_auth/blob/master/LICENSE
[maintain]: https://codeclimate.com/github/pfandie/go_cloud_auth/maintainability

A tool to get temporary AWS credentials

## What does

`go_cloud_auth` makes assume roles easier. With multuple accounts it gets even more easier. No more struggle with messing around with multiple account_ids and different roles to assume. It even comes with it´s own basic MFA (HOTP/TOTP) generation.

## Installation

```bash
go get github.com/pfandie/go_cloud_auth
```

Or download from the release pages <https://github.com/pfandie/go_cloud_auth/releases>

## Usage

```bash
go install go_cloud_auth
go_cloud_auth
```

Or use the downloade binary

## Example config

```yaml
configs:
- config_name: CONFIG-NAME
  settings:
    account_destination: sub-account_id # where you want to assume role in
    account_source: master-account_id # your master/main/parent account
    personal_profile: main-aws-profile # profile configured with 'aws configure' (most likely it´s default in ~/.aws/credentials)
    aws_profile: aws-profile-to-save # name of the new profile (in ~/.aws/credentials, will be shown on profile select)
    token_profile: token-keyring_id # name of your mfa-profile (saved in os-keyring)
    duration: "3600" # how long to assume role
    aws_username: john.dow # your personal aws account-name (which you use to login in aws console ui)
    role_name: role-path/role-name # the role you want to assume in the sub-account
    region: eu-central-1 # the default region, which should be set in ~/.aws/config
```

## Contribution

1. Fork ([https://github.com/ppfandie/go_cloud_auth/fork](https://github.com/pfandie/go_cloud_auth/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `make test` command and confirm that it passes
6. Run `gofmt -s -w .`
7. Create new Pull Request

## Packages used

- github.com/aws/aws-sdk-go/aws
- github.com/spf13/cobra (viper)
- github.com/manifoldco/promptui
- github.com/AlecAivazis/survey/v2
- github.com/zalando/go-keyring
- github.com/mitchellh/go-homedir

### Open ToDos

- add examples
- check if default config exists (config may include another config file)
- implement tests
- allow custom setting for shared credentials
- allow custom config files
- stick to one: manifoldco/promptui or AlecAivazis/survey/v2
- add profile update/change/delete
- add validation for settings
- add 'real' cobra commands for direct auth/configs
- add viper for config handling
- cleanup code
