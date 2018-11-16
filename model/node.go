package model

import (
	"path/filepath"
	"strings"
)

type (
	Node struct {
		Title    string  `json:"title" yaml:"title"`           // 标题
		Path     string  `json:"path" yaml:"path,omitempty"`   // 源文件地址
		Child    []*Node `json:"child" yaml:"child,omitempty"` // 子目录
		Parent   *Node   `json:"parent" yaml:"-"`
		Prev     *Node   `json:"prev" yaml:"-"`
		Next     *Node   `json:"next" yaml:"-"`
		Ext      string  `json:"ext" yaml:"-"`
		Link     string  `json:"link" yaml:"-"`
		BaseDir  string  `json:"base_dir" yaml:"-"`
		IsIndex  bool    `json:"is_index" yaml:"-"`
		IsFile   bool    `json:"is_file" yaml:"-"`
		IsActive bool    `json:"is_active" yaml:"-"`
	}
)

func NewNode() *Node {
	return &Node{}
}

func (this *Node) Init() {
	this.Ext = filepath.Ext(this.Path)
	if this.Path != "" {
		this.IsFile = true
	}
	if this.Path == "index.md" {
		this.IsIndex = true
	}
	if !this.IsIndex {
		this.Link = strings.TrimSuffix(this.Path, this.Ext)
		dep := len(strings.Split(this.Path, "/"))
		for i := 0; i < dep; i++ {
			this.BaseDir += "../"
		}
	}
}

func (this *Node) SetActive(isActive bool) {
	this.IsActive = isActive
	if this.Parent != nil {
		this.Parent.SetActive(isActive)
	}
}
