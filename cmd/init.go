package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init [name]",
		Short: "Initialize the document configuration",
		Long:  `Initialize the document configuration`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			makeConfig(args)
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
}

func makeConfig(args []string) {
	name := args[0]
	cfg = &Config{
		SiteName: name,
		Author:   "YOUNAME",
		Theme:    "default",
		DocsDir:  "docs",
		SiteDir:  "site",
	}
	cbyte, _ := yaml.Marshal(cfg)
	err := ioutil.WriteFile(configFile, cbyte, 0666)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Initialization is complete, you can manually modify config file:", configFile)
}

func loadConfig() {
	cbyte, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	cfg = &Config{}
	err = yaml.Unmarshal(cbyte, cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)
}
