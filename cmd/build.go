package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/flosch/pongo2"
	"github.com/spf13/cobra"
)

var (
	buildCmd = &cobra.Command{
		Use:   "build",
		Short: "Building HTML site",
		Long:  `Building HTML site`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {

			loadConfig()
			loadNav()
			runBuild()
		},
	}
)

func init() {
	rootCmd.AddCommand(buildCmd)
}

func runBuild() {
	fmt.Println("BUILD")
	fmt.Println(nav["模板制作手册"])

	tpl, err := pongo2.FromFile("template/default/body.html")
	if err != nil {
		panic(err)
	}
	// Now you can render the template with the given
	// pongo2.Context how often you want to.
	out, err := tpl.Execute(pongo2.Context{
		"config": cfg,
		"nav":    nav,
	})
	if err != nil {
		panic(err)
	}
	writeFile("site/index.html", out)
}

func writeFile(path string, data string) {
	ioutil.WriteFile(path, []byte(data), 0666)
}
