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
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)



// disconnectCmd represents the disonnect command
var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "kill xcodebuild and iproxy",
	Long: `find xcodebuild and iproxy process pid, then kill them`,
	Run: func(cmd *cobra.Command, args []string) {
		disconnect()
	},
}

func init() {
	rootCmd.AddCommand(disconnectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// disconnectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// disconnectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



func disconnect() {
	iproxyPid, err := getProcessPid("iproxy")
	if err != nil {
		if !errors.Is(err, ErrProcessNotFound) {
			fmt.Println(err)
		}
	} else {
		killProcess(iproxyPid)
	}

	xcodePid, err := getProcessPid("xcodebuild")
	if err != nil {
		if !errors.Is(err, ErrProcessNotFound) {
			fmt.Println(err)
		}
	} else {
		killProcess(xcodePid)
	}
}