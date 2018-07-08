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
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {

}

// NewCmdProcess cria um novo processo
func NewCmdProcess(cmdStr string, processName string) int {
	args := strings.Split(cmdStr, " ")
	commandName := args[0]

	// Get the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Transfer stdin, stdout, and stderr to the new process
	// and also set target directory for the shell to start in.
	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}

	// Get absolute path of command
	commandPath, err := exec.LookPath(args[0])
	if err != nil {
		panic(err)
	}

	// Start up a new shell.
	// Note that we supply "login" twice.
	// -fpl means "don't prompt for PW and pass through environment."
	color.Cyan("[START] Iniciando Docker container...\n")
	proc, err := os.StartProcess(commandPath, args, &pa)
	if err != nil {
		panic(err)
	}

	// Wait until user exits the shell
	state, err := proc.Wait()
	if err != nil {
		panic(err)
	}

	// Force destroy docker container
	if commandName == "docker" {
		exec.Command("docker", "rm", "-f", processName).Run()
	}

	exitCode, err := strconv.Atoi(strings.Replace(state.String(), "exit status ", "", 1))

	return exitCode
}

// GetDockerCommand run dependencies and return basic Docker command
func GetDockerCommand(processName string) string {
	// Get current path
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Virtualenv, source code and output path
	configPath := currentPath + "/.brasilio"
	srcPath := currentPath + "/src"
	packagePath := currentPath + "/package"

	// Create all paths if not exist
	os.MkdirAll(srcPath, os.ModePerm)
	os.MkdirAll(configPath, os.ModePerm)
	os.MkdirAll(packagePath, os.ModePerm)

	// Fix path for Windows
	if runtime.GOOS == "windows" {
		// Recreate paths
		// TODO add support to all Windows disk letters
		srcPath = strings.Replace(srcPath, "\\", "/", -1)
		packagePath = strings.Replace(packagePath, "\\", "/", -1)
		srcPath = strings.Replace(srcPath, "C:", "//c", -1)
		packagePath = strings.Replace(packagePath, "C:", "//c", -1)

		// Create volume if don't exist
		configPath = "brasilio"
		cmd := exec.Command("docker", "volume", "inspect", configPath)
		if err := cmd.Run(); err != nil {
			color.HiYellow(fmt.Sprintf("[WARN] Volume '%s' não encontrado! \n", configPath))
			cmd2 := exec.Command("docker", "volume", "create", configPath)
			if err := cmd2.Run(); err != nil {
				panic(fmt.Sprintf("Não foi possível criar o volume '%s' do Docker!", configPath))
			}
			color.Cyan(fmt.Sprintf("[INFO] Volume '%s' criado com sucesso. \n", configPath))
		}

		// DEBUG
		// fmt.Println("Is Windows.")
		// fmt.Println("configPath:", configPath)
		// fmt.Println("srcPath:", srcPath)
		// fmt.Println("packagePath:", packagePath)
	}

	// Prepare command
	cmdStr := "docker run --rm -it" +
		fmt.Sprintf(" --name=%s", processName) +
		fmt.Sprintf(" -v=%s:/app/src", srcPath) +
		fmt.Sprintf(" -v=%s:/app/.brasilio", configPath) +
		fmt.Sprintf(" -v=%s:/app/package", packagePath) +
		" thenets/brasilio:latest sh"

	return cmdStr
}

// IsCommandAvailable check true if command is available
func IsCommandAvailable(name string) bool {
	cmd := exec.Command(name, "--version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
