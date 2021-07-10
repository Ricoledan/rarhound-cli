/*
Copyright © 2021 Ricardo Ledan <ricoledan@gmail.com>

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
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Woof! Woof! 🐕",
	Long:  `Unpacks RAR files recursively within a specified directory so you can concentrate on enjoying your content.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		directories := args
		fmt.Println("🐕 sniff...sniff...Woof! " + strings.Join(args, " "))
		fmt.Println(directories)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
