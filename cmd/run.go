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
	"path/filepath"
	"os"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Executa o(s) script(s) para obtenção de dados",
	Long: `Executa o script de obtenção de dados e gera os 
arquivos CSV como saída.

Verifique se todos os arquivos estão dentro do diretório 'src':
(o número no início do arquivo determina a ordem de execução)
- ./meu-projeto/src/01-extrator-magistrados.py
- ./meu-projeto/src/02-extrator-juizes.py
- ./meu-projeto/src/03-extrator-politicos.py

Além disso, os arquivos deverão ser enviados para o diretório 'data':
- ./meu-projeto/data/magistrados.csv
- ./meu-projeto/data/juizes.csv
- ./meu-projeto/data/politicos.csv`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")

		// Get current path
		currentPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}

		// Temp path
		envPath := currentPath+"/.brasilio/env"
		os.MkdirAll(envPath, os.ModePerm)

		// Source code path
		sourcePath := currentPath+"/src"

		// Output data path
		dataPath := currentPath+"/data"

		// Run command
		cmdStr := "docker run --rm"+
			fmt.Sprintf(" -v='%s:/app/src'", sourcePath) +
			fmt.Sprintf(" -v='%s:/app/env'", envPath) +
			fmt.Sprintf(" -v='%s:/app/data'", dataPath) +
			" thenets/brasilio"
		shellCmd := exec.Command("/bin/sh", "-c", cmdStr)

		// Redirect output data
		shellCmd.Stdout = os.Stdout
		shellCmd.Stderr = os.Stderr
		shellCmd.Run()

		// Output
		out, _ := shellCmd.Output()  
		fmt.Printf("%s", out)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
