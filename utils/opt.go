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
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/zalando/go-keyring"
)

// ServiceName is used to save values to keychain
const ServiceName string = "go_cloud_auth"

// ShowToken does foo
func ShowToken() {
	// TODO: list token profiles?
	t := promptui.Prompt{
		Label: "Which profile?",
	}
	name, err := t.Run()
	if err != nil {
		log.Fatal(err)
	}

	key, err := keyring.Get(ServiceName, name)
	if err != nil {
		log.Fatal(err)
	}

	token := generateToptToken(key)

	fmt.Println(token)
}

// SetSerial saves a new token to system keychain
func SetSerial() {
	// Name of token
	t := promptui.Prompt{
		Label: "Token Name",
	}
	name, err := t.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Serial of token
	s := promptui.Prompt{
		Label: "Secret / Serial",
	}
	secret, err := s.Run()
	if err != nil {
		log.Fatal(err)
	}

	key := keyring.Set(ServiceName, name, secret)
	if key != nil {
		log.Fatal(key)
	}

	fmt.Printf("Token for %v set successfully\n", name)
}

// DeleteSerial saves a new token to system keychain
func DeleteSerial() {
	// TODO: list token profiles?
	t := promptui.Prompt{
		Label: "Which profile?",
	}
	name, err := t.Run()
	if err != nil {
		log.Fatal(err)
	}

	// first check if this key exists
	_, err = keyring.Get(ServiceName, name)
	if err != nil {
		log.Fatal(err)
	}

	err = keyring.Delete(ServiceName, name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Token for %v deleted successfully\n", name)
}

// GetAuthToken used for get token based on profile from config
func GetAuthToken(profile string) string {
	// get secret by profile name
	secret, err := keyring.Get(ServiceName, profile)
	if err != nil {
		log.Fatal(err)
	}

	return generateToptToken(secret)
}

// generateToptToken adds 30sec to call hotp
func generateToptToken(secret string) string {
	// actual unix timestamp
	now := time.Now().Unix()
	// TODO: make me variable
	intervalLength := 30
	counter := uint64(math.Floor(float64(now) / float64(intervalLength)))

	return generateHotpToken(secret, counter)
}

// generateHotpToken generates a hotp token from secret
func generateHotpToken(secret string, t uint64) (pass string) {
	// make sure to trim whitespace and transform to uppercase
	secret = strings.TrimSpace(strings.ToUpper(secret))
	secretAsBytes, _ := base32.StdEncoding.DecodeString(secret)

	buffer := make([]byte, 8)
	binary.BigEndian.PutUint64(buffer, t)

	// TODO: make use of dynamic algorithms
	hash := sha1.New
	mac := hmac.New(hash, secretAsBytes)
	mac.Write(buffer)
	h := mac.Sum(nil)
	offset := (h[19] & 15)

	var header uint32
	reader := bytes.NewReader(h[offset : offset+4])
	err := binary.Read(reader, binary.BigEndian, &header)
	if err != nil {
		log.Fatal(err)
	}

	h12 := (int(header) & 0x7fffffff) % 1000000
	otp := int(h12)

	return padleft0(otp, 6)
}

// padleft0 prepends 0 to defined length
func padleft0(val int, length int) string {
	return fmt.Sprintf("%0*d", length, val)
}
