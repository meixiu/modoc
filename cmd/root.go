package cmd

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

type (
	Config struct {
		SiteName string `json:"site_name" yaml:"site_name"`
		Author   string `json:"author" yaml:"author"`
		DocsDir  string `json:"docs_dir" yaml:"docs_dir"`
		SiteDir  string `json:"site_dir" yaml:"site_dir"`
		Theme    string `json:"theme" yaml:"theme"`
	}

	FileNode struct {
		Title string      `json:"name" yaml:"title"`
		Path  string      `json:"path" yaml:"path"`
		Nodes []*FileNode `json:"nodes" yaml:"nodes"`
	}

	FileTree map[string]interface{}
)

var (
	configFile = "config.yaml"
	navFile    = "nav.yaml"
	rootCmd    = &cobra.Command{
		Use:   "mkdoc",
		Short: "Project documentation with Markdown",
		Long:  `Project documentation with Markdown by GO`,
	}
	appPath string
	cfg     *Config
	nav     FileTree
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
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	appPath = home
}
