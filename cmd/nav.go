package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

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

func makeNav() {
	info, err := os.Lstat(cfg.DocsDir)
	if err != nil {
		log.Fatal(err)
	}
	node := make(FileTree, 0)
	walkFile(cfg.DocsDir, info, node)
	//tree, _ := json.MarshalIndent(node, " ", "  ")
	tree, _ := yaml.Marshal(node)
	err = ioutil.WriteFile(navFile, tree, 0666)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generate the navigation configuration file:", navFile)
}

func loadNav() {
	cbyte, err := ioutil.ReadFile(navFile)
	if err != nil {
		log.Fatal(err)
	}
	nav = make(FileTree, 0)
	err = yaml.Unmarshal(cbyte, &nav)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", nav)
}

func walkFile(path string, info os.FileInfo, node FileTree) {
	// 列出当前目录下的所有目录、文件
	files := listFiles(path)
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
			child := make(FileTree, 0)
			node[filename] = child
			walkFile(fpath, finfo, child)
		} else {
			ext := filepath.Ext(fpath)
			name := strings.TrimSuffix(filename, ext)
			if ext != ".md" {
				continue
			}
			link, _ := filepath.Rel(cfg.DocsDir, fpath)
			node[name] = link
			fmt.Println("find:", fpath)
		}
	}
	return
}

func listFiles(dirname string) []string {
	f, _ := os.Open(dirname)
	names, _ := f.Readdirnames(-1)
	f.Close()
	sort.Strings(names)
	return names
}

func isFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
