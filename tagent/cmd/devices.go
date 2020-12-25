/*
Copyright © 2020 Quinn <zxk.albert@foxmail.com>

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
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

// devicesCmd represents the devices command
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "list devices UDID",
	Long: `list all attached devices UDID`,
	Run: func(cmd *cobra.Command, args []string) {

		devices := getDevices()

		fmt.Println("List of devices attached")
		for _, device := range devices {
			fmt.Println(device)
		}
	},
}

func init() {
	rootCmd.AddCommand(devicesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// devicesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// devicesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getDevices() (devices []string) {
	out, err := exec.Command("cfgutil", "list").Output()
	if err != nil {
		log.Println(err)
		fmt.Printf("%s\b", err)
		os.Exit(1)
	}

	r, _ := regexp.Compile(`UDID:\s([0-9a-z]+)\s`)
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		deviceText := scanner.Text()
		deviceUDID := r.FindString(deviceText)
		devices = append(devices, deviceUDID)
	}
	return
}
