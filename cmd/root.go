package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/wbsifan/modoc/helper"

	"github.com/wbsifan/modoc/asset"

	"github.com/wbsifan/modoc/model"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	configFile = "config.yaml"
	navFile    = "nav.yaml"
	rootCmd    = &cobra.Command{
		Use:   "mkdoc",
		Short: "Generate HTML web site for MD file directory",
		Long:  `Generate HTML web site for MD file directory`,
	}
	home      string
	customTpl bool
	cfg       *model.Config
	nav       *model.Node
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	h, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	// Modoc Cache path
	home = filepath.Join(h, ".modoc")
	// custom template
	if helper.IsDir("template") {
		customTpl = true
	} else {
		if !helper.IsDir(filepath.Join(home, "template")) {
			fmt.Println("First run, Installing...")
			unpackAsset()
		}
	}
}

// unpackAsset 解压静态资源
func unpackAsset() {
	dirs := []string{"template"}
	isSuccess := true
	for _, dir := range dirs {
		if err := asset.RestoreAssets(home, dir); err != nil {
			isSuccess = false
			fmt.Println(err)
			break
		}
	}
	if !isSuccess {
		for _, dir := range dirs {
			os.RemoveAll(filepath.Join(home, dir))
		}
		log.Fatal("unzip asset failure")
	}
}

func getAsset(path string) string {
	if customTpl {
		return filepath.Join("template", path)
	} else {
		return filepath.Join(home, "template", path)
	}

}
