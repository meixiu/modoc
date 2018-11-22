package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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
	themePath string
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
}
