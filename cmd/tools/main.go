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
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/fatih/color"
)

func main() {

}

// NewCmdProcess cria um novo processo
func NewCmdProcess(cmdStr string, processName string) *exec.Cmd {
	args := strings.Split(cmdStr, " ")
	commandName := args[0]

	i := 0                     // Remove first item from args
	copy(args[i:], args[i+1:]) // Shift a[i+1:] left one index
	args[len(args)-1] = ""     // Erase last element (write zero value)
	args = args[:len(args)-1]  // Truncate slice

	shellCmd := exec.Command(commandName, args...)
	shellCmd.Stdout = os.Stdout
	shellCmd.Stderr = os.Stderr

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		// Force destroy docker container
		if commandName == "docker" {
			exec.Command("docker", "rm", "-f", processName).Run()
		}
		// Exit message
		color.Cyan("\n\n[EXIT] Ctrl+C pressed in Terminal\n")
		// Kill main command process
		shellCmd.Process.Kill()
		// Main exit
		time.Sleep(300 * time.Millisecond)
		os.Exit(0)
	}()

	return shellCmd
}

// GetDockerCommand run dependencies and return basic Docker command
func GetDockerCommand(processName string) string {
	// Get current path
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Virtualenv, source code and output path
	envPath := currentPath + "/.brasilio/env"
	srcPath := currentPath + "/src"
	packagePath := currentPath + "/package"

	// Create all paths if not exist
	os.MkdirAll(srcPath, os.ModePerm)
	os.MkdirAll(envPath, os.ModePerm)
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
		envPath = "brasilio_env"
		cmd := exec.Command("docker", "volume", "inspect", envPath)
		if err := cmd.Run(); err != nil {
			color.Cyan(fmt.Sprintf("\n\n[WARN] Volume '%s' não encontrado! \n", envPath))
			cmd2 := exec.Command("docker", "volume", "create", envPath)
			if err := cmd2.Run(); err != nil {
				panic(fmt.Sprintf("Não foi possível criar o volume '%s' do Docker!", envPath))
			}
			color.Cyan(fmt.Sprintf("\n\n[INFO] Volume '%s' criado com sucesso. \n", envPath))
		}

		// DEBUG
		// fmt.Println("Is Windows.")
		// fmt.Println("envPath:", envPath)
		// fmt.Println("srcPath:", srcPath)
		// fmt.Println("packagePath:", packagePath)
	}

	// Prepare command
	cmdStr := "docker run --rm" +
		fmt.Sprintf(" --name=%s", processName) +
		fmt.Sprintf(" -v=%s:/app/src", srcPath) +
		fmt.Sprintf(" -v=%s:/app/env", envPath) +
		fmt.Sprintf(" -v=%s:/app/package", packagePath) +
		" thenets/brasilio:latest"

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

// GetSupportMessage return the Brasil.io default support message
func GetSupportMessage() string {
	return "# Se precisar de suporte, acesse: https://chat.brasil.io/"
}
