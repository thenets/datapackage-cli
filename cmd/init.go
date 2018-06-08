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

package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/thenets/brasilio-cli/cmd/tools"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inicia um novo projeto do Brasil.io",
	Long: `O novo projeto do Brasil.io terá tudo que será necessário para começar
a desenvolver um novo código de extração e/ou tratamento de dados.

Um código em Python será criado como exemplo para ajudar na criação de um 'datapackage'.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if has argument
		if len(args) != 1 {
			fmt.Println("ERROR: É necessário informar o nome do novo projeto.")
			fmt.Println("Execute 'brasilio init --help' para mais detalhes.")
			os.Exit(0)
		}

		// Check if Docker is available
		if tools.IsCommandAvailable("docker") == false {
			fmt.Println("ERROR: 'docker' não está instalado.")
			fmt.Println("Acesse https://docs.docker.com/install/ para mais informações.")
			os.Exit(0)
		}

		// Check if Git is available
		if tools.IsCommandAvailable("git") == false {
			fmt.Println("ERROR: 'git' não está instalado.")
			fmt.Println("Acesse https://git-scm.com/downloads para mais informações.")
			os.Exit(0)
		}

		// Check if project already exist
		if _, err := os.Stat(args[0]); err == nil {
			fmt.Printf("ERROR: o projeto '%s' já existe!\n", args[0])
			fmt.Println(tools.GetSupportMessage())
			os.Exit(0)
		}

		// Clone new project
		shellCmd := exec.Command("git", "clone", "https://github.com/thenets/brasilio-package", args[0])
		shellCmd.Stdout = os.Stdout
		shellCmd.Stderr = os.Stderr
		if err := shellCmd.Run(); err != nil {
			fmt.Println(tools.GetSupportMessage())
			os.Exit(0)
		}

		// Remove git reference
		os.RemoveAll(args[0]+"/.git")

		fmt.Printf("\n[DONE] Projeto '%s' criado!\n", args[0])
		fmt.Printf("[TIP ] Entre no diretório 'cd %s' e execute 'brasilio run' para começar! :)\n", args[0])

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
