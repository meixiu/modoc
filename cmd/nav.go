package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/wbsifan/modoc/helper"

	"github.com/wbsifan/modoc/model"

	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

var (
	navCmd = &cobra.Command{
		Use:   "nav",
		Short: "Generate the navigation configuration file",
		Long:  `Generate the navigation configuration file`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			loadConfig()
			makeNav()
		},
	}
)

func init() {
	rootCmd.AddCommand(navCmd)
}

func loadNav() {
	cbyte, err := ioutil.ReadFile(navFile)
	if err != nil {
		log.Fatal(err)
	}
	nav = model.NewNode()
	err = yaml.Unmarshal(cbyte, &nav)
	if err != nil {
		log.Fatal(err)
	}
}

func makeNav() {
	info, err := os.Lstat(cfg.DocsDir)
	if err != nil {
		log.Fatal(err)
	}
	node := model.NewNode()
	walkFile(cfg.DocsDir, info, node)
	tree, _ := yaml.Marshal(node)
	err = ioutil.WriteFile(navFile, tree, 0666)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generate the navigation configuration file:", navFile)
}

func walkFile(path string, info os.FileInfo, node *model.Node) {
	// 列出当前目录下的所有目录、文件
	files := helper.ListFiles(path)
	// 遍历这些文件
	for _, filename := range files {
		// 以 `_` 开头不加载
		if strings.HasPrefix(filename, "__") {
			continue
		}
		fpath := filepath.Join(path, filename)
		finfo, _ := os.Lstat(fpath)
		isDir := finfo.IsDir()
		// 如果遍历的当前文件是个目录，则进入该目录进行递归
		if isDir {
			child := &model.Node{
				Title: filename,
			}
			walkFile(fpath, finfo, child)
			node.Child = append(node.Child, child)
		} else {
			ext := filepath.Ext(fpath)
			name := strings.TrimSuffix(filename, ext)
			if ext != ".md" {
				continue
			}
			if fpath == filepath.Join(cfg.DocsDir, "index.md") {
				name = cfg.IndexTitle
			}
			link, _ := filepath.Rel(cfg.DocsDir, fpath)
			link = strings.Replace(link, "\\", "/", -1)
			child := &model.Node{
				Title: name,
				Path:  link,
			}
			node.Child = append(node.Child, child)
			fmt.Println("find:", fpath)
		}
	}
}
