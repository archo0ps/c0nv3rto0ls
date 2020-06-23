/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"
	"strings"
)

var rot13Ruler1 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var rot13Ruler2 = "nopqrstuvwxyzabcdefghijklmNOPQRSTUVWXYZABCDEFGHIJKLM"

// rot13Cmd represents the rot13 command
var rot13Cmd = &cobra.Command{
	Use:   "rot13",
	Short: "Rot13 decode & encode tool",
	Run: func(cmd *cobra.Command, args []string) {
		if !(strings.EqualFold(pattern, "decode") || strings.EqualFold(pattern, "encode")) {
			_ = cmd.Help()
			return
		}
		result, _ := rot13(text)
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(rot13Cmd)
	rot13Cmd.Flags().StringVarP(&pattern, "pattern", "p", "", "Decode or encode")
	rot13Cmd.Flags().StringVarP(&text, "text", "t", "", "Text to be processed")
}

func rot13(ror13String string) (string, error) {
	result := ""
	for _, value := range ror13String {
		if strings.Index(rot13Ruler1, string(value)) != -1 {
			result += string(rot13Ruler2[strings.Index(rot13Ruler1, string(value))])
		} else {
			result += string(value)
		}
	}
	return result, nil
}
