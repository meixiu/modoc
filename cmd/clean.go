package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	cleanCmd = &cobra.Command{
		Use:   "clean",
		Short: "Clear site content and cache",
		Long:  `Clear site content and cache`,
		Run: func(cmd *cobra.Command, args []string) {
			runClean()
		},
	}
)

func init() {
	rootCmd.AddCommand(cleanCmd)
}

func runClean() {
	loadConfig()
	dirs := []string{cfg.SiteDir}
	for _, dir := range dirs {
		fmt.Println("clean:", dir)
		err := os.RemoveAll(dir)
		if err != nil {
			log.Fatal(err)
		}
	}
}
