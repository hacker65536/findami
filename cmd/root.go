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
	"context"
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type appConfig struct {
	AMIFilters []AMIFilter
}

type AMIFilter struct {
	Name   string
	Values string
}

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "findami",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		debug, _ := cmd.Flags().GetBool("debug")
		verbose, _ := cmd.Flags().GetBool("verbose")
		if debug {
			log.SetLevel(log.DebugLevel)
		}

		params := &ec2.DescribeImagesInput{}
		if len(args) == 0 {
			c := &appConfig{}
			log.WithFields(
				log.Fields{
					"c=": c,
				}).Debug("[debug]")
			if err := viper.Unmarshal(&c); err != nil {
				log.WithFields(log.Fields{
					"err": err,
				}).Fatal("Can't Unmarshal")
			}

			fs := []types.Filter{}
			for _, v := range c.AMIFilters {
				fs = append(fs, types.Filter{
					Name:   aws.String(v.Name),
					Values: []string{v.Values},
				},
				)

			}
			log.WithFields(
				log.Fields{
					"fs=": fs,
				}).Debug("[debug]")
			params.Filters = fs
			log.Debug("no args")
		} else {
			params.ImageIds = args
			//ImageIds: []string{"ami-fc3b1199"},
			log.WithFields(
				log.Fields{
					"args": args,
				}).Debug("[debug]")

		}
		findami(params, verbose)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.findami.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("debug", "d", false, "show debug message")
	rootCmd.Flags().BoolP("verbose", "v", false, "show detal with json. Only ami-id sepcified")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".findami" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".findami")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		log.WithFields(log.Fields{
			"file:": viper.ConfigFileUsed(),
		}).Debug("Using config")
	}
}

func findami(params *ec2.DescribeImagesInput, verbose bool) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("unable to load SDK config")
	}

	svc := ec2.NewFromConfig(cfg)

	log.WithFields(log.Fields{
		"params=": params,
	}).Debug("[debug]")
	resp, err := svc.DescribeImages(context.TODO(), params)

	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Can't list images")
	}

	if params.ImageIds != nil && verbose  {
		j, _ := json.Marshal(resp)
		fmt.Println(string(j))
	} else {
		w := tabwriter.NewWriter(os.Stdout, 2, 0, 3, ' ', 0)
		for _, v2 := range resp.Images {

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
				aws.ToString(v2.ImageId),
				aws.ToString(v2.PlatformDetails),
				v2.Architecture,
				v2.BlockDeviceMappings[0].Ebs.VolumeType,
				aws.ToString(v2.Name),
			)
		}
		w.Flush()
	}
}
