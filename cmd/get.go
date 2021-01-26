/*
Copyright © 2020 Mazedur Rahman <mazedur.rahman.litn@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	reader "github.com/yamaszone/gcp-secret-manager-buddy/internal"
	"os"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch latest value of secrets from GCP Secret Manager",
	Long: `Fetch secrets from GCP Secret Manager given a JSON input file of the
following format:
{
	"KEY1":"secret-name1",
	"KEY2":"secret-name2"
}

Output:
{
	"KEY1":"secret-value1",
	"KEY2":"secret-value2"
}`,
	RunE: getSecrets,
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&input, "input", "i", "", "Input JSON file name")
	getCmd.MarkFlagRequired("input")
	// May make this PersistentFlags in the future
	getCmd.Flags().StringVarP(&project, "project", "p", "", "GCP project name")
	getCmd.MarkFlagRequired("project")
	// May support version selection in the future. For now, keep it simple and stupid.
	//getCmd.Flags().StringVarP(&version, "version", "v", "latest", "Secret version (default: latest)")
	//getCmd.MarkFlagRequired("version")
}

func getSecrets(cmd *cobra.Command, args []string) error {
	_, present := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
	if !present {
		fmt.Println("Error: GOOGLE_APPLICATION_CREDENTIALS not set in the environment.")
		os.Exit(1)
	}

	inputFile, err := cmd.Flags().GetString("input")
	if err != nil {
		return err
	}
	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}

	reader.GetSecrets(inputFile, project, "latest")

	return nil
}
