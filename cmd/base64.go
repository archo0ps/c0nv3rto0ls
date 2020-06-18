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
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const ERROR_BASE64_TEXT = "There were invalid base64."
const ERROR_NIL_TEXT = "Please input text."

var pattern string
var text string

var base64 = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V",
	"W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i",
	"j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
	"w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8",
	"9", "+", "/"}

var base64String = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Base64 decode & encode tool",
	Run: func(cmd *cobra.Command, args []string) {
		if !(strings.EqualFold(pattern, "decode") || strings.EqualFold(pattern, "encode")) {
			_ = cmd.Help()
			return
		}
		if pattern == "decode" {
			result, err := decode(text)
			if err != nil {
				fmt.Println(err)
				_ = cmd.Help()
				return
			}
			fmt.Println(result)
		} else {
			result, err := encode(text)
			if err != nil {
				fmt.Println(err)
				_ = cmd.Help()
				return
			}
			fmt.Println(result)
		}
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.Flags().StringVarP(&pattern, "pattern", "p", "", "Decode or encode")
	base64Cmd.Flags().StringVarP(&text, "text", "t", "", "Text to be processed")

}

func decode(unDecode string) (string, error) {
	temp := ""
	plaintext := ""
	errorText := `[^A-Za-z0-9+/=]`
	matched, _ := regexp.MatchString(errorText, unDecode)
	if matched {
		return "", errors.New(ERROR_BASE64_TEXT)
	}
	if len(unDecode)%4 != 0 {
		return "", errors.New(ERROR_BASE64_TEXT)
	}
	matchedEquals := regexp.MustCompile("=")
	equals := matchedEquals.FindAllSubmatch([]byte(unDecode), -1)
	if len(equals) > 2 {
		return "", errors.New(ERROR_BASE64_TEXT)
	}

	for _, value := range unDecode {
		if string(value) == "=" {
			break
		}
		//fmt.Println(key, strings.Index(base64String, string(value)),
		//	convertToBin(strings.Index(base64String, string(value)),0),
		//	convertToInt(convertToBin(strings.Index(base64String, string(value)),0)))
		temp += convertToBin(strings.Index(base64String, string(value)), 0)
	}
	//fmt.Println(result)
	//Zero padding at low position
	for len(temp)%8 != 0 {
		temp = temp[:len(temp)-1]
	}
	for i := 0; i < len(temp); i++ {
		if i%8 != 0 {
			continue
		}
		plaintext += string(convertToInt(temp[i : i+8]))
	}
	return plaintext, nil
}

func encode(unEncode string) (string, error) {
	temp := ""
	ciphertext := ""
	if unEncode == "" {
		return "", errors.New(ERROR_NIL_TEXT)
	}
	for _, value := range unEncode {
		temp += convertToBin(int(value), 1)
	}
	//Zero padding at low position
	for len(temp)%6 != 0 {
		temp += "0"
	}
	for i := 0; i < len(temp); i++ {
		if i%6 != 0 {
			continue
		}
		ciphertext += base64[convertToInt(temp[i:i+6])]
	}

	//= padding
	for len(ciphertext)%4 != 0 {
		ciphertext += "="
	}
	return ciphertext, nil
}

//type 1:encode 0:decode
func convertToBin(num int, types int) string {
	bin := ""
	if num == 0 {
		return "0"
	}

	for ; num > 0; num /= 2 {
		lsb := num % 2
		bin = strconv.Itoa(lsb) + bin
	}
	// Zero padding at high position
	if types == 1 {
		if len(bin) != 8 {
			for ; len(bin) < 8; {
				bin = "0" + bin
			}
		}
	} else {
		if len(bin) != 6 {
			for ; len(bin) < 6; {
				bin = "0" + bin
			}
		}
	}

	return bin
}

func convertToInt(string string) int {
	var result float64 = 0
	for key, value := range string {
		if value == 49 {
			result += math.Pow(2, float64(len(string)-key-1)) * 1
		} else {
			result += math.Pow(2, float64(len(string)-key-1)) * 0
		}
	}
	return int(result)
}
