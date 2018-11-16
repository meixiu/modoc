package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/flosch/pongo2"
	"github.com/wbsifan/modoc/helper"
	"github.com/wbsifan/modoc/model"

	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init [name]",
		Short: "Initialize the document configuration",
		Long:  `Initialize the document configuration`,
		Run: func(cmd *cobra.Command, args []string) {
			makeConfig(args)
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
}

func makeConfig(args []string) {
	var siteName = "DEMO"
	if len(args) > 0 {
		siteName = args[0]
	} else {
		cwd, err := os.Getwd()
		if err == nil {
			siteName = filepath.Base(cwd)
		}
	}
	configTpl, err := pongo2.FromString(model.DefaultConfig)
	if err != nil {
		log.Fatal(err)
	}
	out, err := configTpl.Execute(pongo2.Context{
		"siteName": siteName,
	})
	if err != nil {
		log.Fatal(err)
	}
	err = helper.WriteFile("config.yaml", out)
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
	cfg = model.NewConfig()
	err = yaml.Unmarshal(cbyte, cfg)
	if err != nil {
		log.Fatal(err)
	}
}
