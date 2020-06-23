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
	"strings"

	"github.com/spf13/cobra"
)

var caesarString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// caesarCmd represents the caesar command
var caesarCmd = &cobra.Command{
	Use:   "caesar",
	Short: "Caesar decode & encode tool",
	Run: func(cmd *cobra.Command, args []string) {
		if !(strings.EqualFold(pattern, "decode") || strings.EqualFold(pattern, "encode")) {
			_ = cmd.Help()
			return
		}
		if offset <= 0 || offset >= 26 {
			fmt.Println("offset Error!")
			_ = cmd.Help()
			return
		}
		var result = ""
		if pattern == "decode" {
			result, _ = caesarDecode(text)
			fmt.Println(result)
		} else {
			result, _ = caesarEncode(text)
			fmt.Println(result)
		}

	},
}

func init() {
	rootCmd.AddCommand(caesarCmd)
	caesarCmd.Flags().StringVarP(&pattern, "pattern", "p", "", "Decode or encode")
	caesarCmd.Flags().StringVarP(&text, "text", "t", "", "Text to be processed")
	caesarCmd.Flags().IntVarP(&offset, "offset", "o", 0, "Offset of the plaintext or ciphertext")
}

func caesarDecode(unDecode string) (string, error) {
	plainText := ""
	for _, value := range unDecode {
		if strings.Index(caesarString, string(value)) == -1 {
			plainText += string(value)
		} else if strings.Index(caesarString, string(value))-offset < 0 {
			plainText += string(caesarString[strings.Index(caesarString, string(value))-offset+25])
		} else {
			plainText += string(caesarString[strings.Index(caesarString, string(value))-offset])
		}
	}
	return plainText, nil
}

func caesarEncode(unEncode string) (string, error) {
	cipherText := ""
	for _, value := range unEncode {
		if strings.Index(caesarString, string(value)) == -1 {
			cipherText += string(value)
		} else if strings.Index(caesarString, string(value))+offset > len(caesarString) {
			cipherText += string(caesarString[strings.Index(caesarString, string(value))+offset-len(caesarString)])
		} else {
			cipherText += string(caesarString[strings.Index(caesarString, string(value))+offset])
		}
	}
	return cipherText, nil
}
