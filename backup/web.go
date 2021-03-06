/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"bzz-collector/service"

	"github.com/spf13/cobra"
)

var (
	Addr string
	Port string
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "bzz web service",
	Long:  `bzz web service`,
	Run: func(cmd *cobra.Command, args []string) {
		service.StartService(Addr, Port)
	},
}

func init() {
	webCmd.Flags().StringVarP(&Addr, "addr", "a", "addr", "listen address.")
	webCmd.Flags().StringVarP(&Port, "port", "p", "port", "listen port.")
	rootCmd.AddCommand(webCmd)

}
