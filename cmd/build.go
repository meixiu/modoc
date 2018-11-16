package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/wbsifan/modoc/helper"

	"github.com/wbsifan/modoc/model"

	"github.com/PuerkitoBio/goquery"
	"github.com/flosch/pongo2"
	"github.com/mozillazg/go-slugify"
	mdparser "github.com/russross/blackfriday"
	"github.com/spf13/cobra"
)

type (
	Toc struct {
		Title string
		Link  string
		Child []*Toc
	}
)

var (
	bodyTpl  *pongo2.Template
	search   *model.Search
	index    *model.Node
	pinyin   = make(map[string]int)
	nodeList = make([]*model.Node, 0)
	buildCmd = &cobra.Command{
		Use:   "build",
		Short: "Building HTML site",
		Long:  `Building HTML site`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runBuild()
		},
	}
)

func init() {
	rootCmd.AddCommand(buildCmd)
}

func runBuild() {
	loadConfig()
	loadNav()
	loadTpl()
	initSearch()
	initNode(nav)
	makeNode()
	makeStatic()
	makeSearch()
}

func loadTpl() {
	var err error
	bodyTplFile := getAsset(filepath.Join(cfg.Theme, "body.html"))
	bodyTpl, err = pongo2.FromFile(bodyTplFile)
	if err != nil {
		log.Fatal(err)
	}
}

func initSearch() {
	search = &model.Search{
		Config: &model.SearchConfig{
			Lang:          []string{"en"},
			PrebuildIndex: false,
			Separator:     "[\\s\\-]+",
		},
	}
}

func makeSearch() {
	cbyte, _ := json.Marshal(search)
	dstPath := filepath.Join(cfg.SiteDir, "static/search/search_index.json")
	err := helper.WriteFile(dstPath, string(cbyte))
	if err != nil {
		log.Fatal(err)
	}
}

func makeStatic() {
	srcPath := getAsset(filepath.Join(cfg.Theme, "static"))
	dstPath := filepath.Join(cfg.SiteDir, "static")
	fmt.Println("copy:", srcPath, "==>", dstPath)
	err := helper.CopyDir(dstPath, srcPath)
	if err != nil {
		log.Fatal(err)
	}
}

func initNode(node *model.Node) {
	node.Init()
	// 首页
	if node.IsIndex {
		index = node
	}
	// 链接里的汉字转成拼音
	if cfg.LinkPinyin && node.Link != "" {
		links := strings.Split(node.Link, "/")
		var slug []string
		for _, v := range links {
			slug = append(slug, slugify.Slugify(v))
		}
		newlink := strings.Join(slug, "/")
		_, has := pinyin[newlink]
		if has {
			pinyin[newlink] += 1
			newlink = fmt.Sprintf("%s-%v", newlink, pinyin[newlink])
		} else {
			pinyin[newlink] = 1
		}
		node.Link = newlink
	}
	// 如果是文件加入到列表
	if node.IsFile {
		nodeList = append(nodeList, node)
	}
	for _, n := range node.Child {
		n.Parent = node
		initNode(n)
	}
}

func makeNode() {
	max := len(nodeList) - 1
	for i, node := range nodeList {
		node.SetActive(true)
		// 上一页下一页
		prevId := i - 1
		nextId := i + 1
		if prevId < 0 {
			prevId = max
		}
		if nextId > max {
			nextId = 0
		}
		node.Prev = nodeList[prevId]
		node.Next = nodeList[nextId]
		makeBody(node)
		node.SetActive(false)
	}
	// if node.IsFile {
	// 	node.SetActive(true)
	// 	makeBody(node)
	// 	node.SetActive(false)
	// } else {
	// 	for _, node := range node.Child {
	// 		makeNode(node)
	// 	}
	// }
}

func makeBody(node *model.Node) {
	srcPath := filepath.Join(cfg.DocsDir, node.Path)
	dstPath := filepath.Join(cfg.SiteDir, "index.html")
	if node.Path != "index.md" {
		dstPath = filepath.Join(cfg.SiteDir, node.Link, "index.html")
	}
	output := parseMd(srcPath)
	html, text, tocs := parseToc(output)
	// Add search doc
	doc := &model.SearchDoc{
		Location: node.Link,
		Text:     text,
		Title:    node.Title,
	}
	search.AddDoc(doc)
	out, err := bodyTpl.Execute(pongo2.Context{
		"config":  cfg,
		"nav":     nav,
		"tocs":    tocs,
		"content": html,
		"node":    node,
		"index":   index,
		"baseDir": node.BaseDir,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("build:", dstPath)
	err = helper.WriteFile(dstPath, out)
	if err != nil {
		log.Fatal(err)
	}
}

func parseMd(path string) []byte {
	input, _ := ioutil.ReadFile(path)
	renderer := mdparser.NewHTMLRenderer(mdparser.HTMLRendererParameters{
		Flags: mdparser.CommonHTMLFlags | mdparser.TOC,
	})
	output := mdparser.Run(input, mdparser.WithRenderer(renderer))
	return output
}

func parseToc(input []byte) (string, string, *Toc) {
	tocs := &Toc{}
	buf := bytes.NewReader(input)
	dom, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		log.Fatal(err)
	}
	// Find h1 h2
	dom.Find("nav>ul>li").Each(func(i int, li *goquery.Selection) {
		a := li.Find("nav>ul>li>a")
		link, has := a.Attr("href")
		h1 := &Toc{
			Title: a.Text(),
			Link:  link,
		}
		if has {
			tocs.Child = append(tocs.Child, h1)
		}
		li.Find("nav>ul>li>ul>li>a").Each(func(j int, a *goquery.Selection) {
			link, _ := a.Attr("href")
			h2 := &Toc{
				Title: a.Text(),
				Link:  link,
			}
			if has {
				h1.Child = append(h1.Child, h2)
			} else {
				tocs.Child = append(tocs.Child, h2)
			}
		})
	})
	// Remove nav
	dom.Find("nav").Remove()
	html, err := dom.Html()
	if err != nil {
		log.Println(err)
	}
	text := strings.Replace(dom.Text(), "\n\n", "", -1)
	return html, text, tocs
}
