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
	"github.com/spf13/viper"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect ios device",
	Long: `connect ios device.
if there more than one devices attached, 
will select first device to connect.

if tagent already connect,
will run disconnect pre.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		disconnect()
	},
	Run: func(cmd *cobra.Command, args []string) {

		var connectDevice string
		devices := getDevices()

		switch len(devices) {
		case 0:
			fmt.Println("no devices attached")
			os.Exit(1)
		case 1:
			connectDevice = devices[0]
			fmt.Println("connecting to", connectDevice)
		default:
			connectDevice = devices[0]
			fmt.Println("more than one device attached, will select first to connect", connectDevice)
		}

		if viper.IsSet("tagent") == false {
			fmt.Println("not get tagent path")
			os.Exit(1)
		}
		tagentPath := viper.GetString("tagent")

		command := exec.Command("xcodebuild",
			"-project",
			tagentPath,
			"-scheme",
			"WebDriverAgentRunner",
			"-destination",
			fmt.Sprintf("id=%s", connectDevice),
			"test")
		err := command.Start()
		if err != nil {
			panic(err)
		}
		fmt.Printf("xcode connected, [PID] %d running...\n", command.Process.Pid)
		//ioutil.WriteFile("tagent.lock", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)

		command = exec.Command("iproxy", "8100", "8100")
		err = command.Start()
		if err != nil {
			panic(err)
		}
		fmt.Printf("iproxy connected, [PID] %d running...\n", command.Process.Pid)

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
