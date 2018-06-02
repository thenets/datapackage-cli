// Copyright © 2018 Luiz Felipe F M Costa <luiz@thenets.org>
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

package tools

import (
	"fmt"
	"path/filepath"
	"os"
)

func getDockerCommand() string{
	// Get current path
	currentPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	// Virtualenv, source code and output path
	envPath 	:= currentPath+"/.brasilio/env"
	srcPath 	:= currentPath+"/src"
	packagePath := currentPath+"/package"

	// Create all paths if not exist
	os.MkdirAll(srcPath, os.ModePerm)
	os.MkdirAll(envPath, os.ModePerm)
	os.MkdirAll(packagePath, os.ModePerm)

	// Prepare command
	cmdStr := "docker run --rm"+
		fmt.Sprintf(" -v='%s:/app/src'", srcPath) +
		fmt.Sprintf(" -v='%s:/app/env'", envPath) +
		fmt.Sprintf(" -v='%s:/app/package'", packagePath) +
		" thenets/brasilio"

	return cmdStr
}
