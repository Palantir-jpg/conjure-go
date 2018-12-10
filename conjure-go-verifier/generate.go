// Copyright (c) 2018 Palantir Technologies. All rights reserved.
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

// +build ignore

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/pkg/errors"

	"github.com/palantir/conjure-go/cmd"
)

const conjureVerifierVersion = "0.8.0"

func main() {
	const versionFilePath = "version_test.go"
	newVersionFileContent := fmt.Sprintf(`// This is a generated file: do not edit by hand.
// To update this file, run the generator in conjure-go-verifier.
package verifier_test

const verificationServerVersion = "%s"
`, conjureVerifierVersion)

	// version file exists and is in desired state: assume that all generated content is in desired state
	if currVersionFileContent, err := ioutil.ReadFile(versionFilePath); err == nil && string(currVersionFileContent) == newVersionFileContent {
		return
	}

	const testCasesFile = "test-cases.json"
	if err := downloadFile(testCasesFile, fmt.Sprintf("https://palantir.bintray.com/releases/com/palantir/conjure/verification/test-cases/%s/test-cases-%s.json", conjureVerifierVersion, conjureVerifierVersion)); err != nil {
		panic(err)
	}

	const verificationAPIFile = "verification-api.conjure.json"
	if err := downloadFile(verificationAPIFile, fmt.Sprintf("https://palantir.bintray.com/releases/com/palantir/conjure/verification/verification-api/%s/verification-api-%s.conjure.json", conjureVerifierVersion, conjureVerifierVersion)); err != nil {
		panic(err)
	}

	if err := cmd.Generate(verificationAPIFile, "."); err != nil {
		panic(err)
	}

	// update version in circle.yml
	if err := updateVersionInCircleConfig("../.circleci/config.yml", conjureVerifierVersion); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(versionFilePath, []byte(newVersionFileContent), 0644); err != nil {
		panic(err)
	}
}

var dockerImgRegexp = regexp.MustCompile(`(?m)^(.*?- image: palantirtechnologies/conjure-verification-server):(.+)$`)

func updateVersionInCircleConfig(cfgFile, version string) error {
	bytes, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return errors.WithStack(err)
	}
	cfgContent := string(bytes)

	updated := dockerImgRegexp.ReplaceAllString(cfgContent, fmt.Sprintf("$1:%s", version))
	if cfgContent == updated {
		return nil
	}
	if err := ioutil.WriteFile(cfgFile, []byte(updated), 0644); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("GET %s received bad response status: %s", url, resp.Status)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}
	return nil
}
