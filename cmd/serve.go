package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start web services",
		Long:  `Start web services`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runBuild()
			runServe()
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

func runServe() {
	fmt.Println("Start server:", "http://127.0.0.1"+cfg.DevAddr)
	err := http.ListenAndServe(cfg.DevAddr, http.FileServer(http.Dir(cfg.SiteDir)))
	if err != nil {
		log.Fatal("ListenAndServe fail:", err)
	}
}
