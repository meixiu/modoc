package model

type (
	Config struct {
		SiteName   string `json:"site_name" yaml:"site_name"`
		SiteAddr   string `json:"site_addr" yaml:"site_addr"`
		Favicon    string `json:"favicon" yaml:"favicon"`
		Author     string `json:"author" yaml:"author"`
		DocsDir    string `json:"docs_dir" yaml:"docs_dir"`
		SiteDir    string `json:"site_dir" yaml:"site_dir"`
		Theme      string `json:"theme" yaml:"theme"`
		IndexTitle string `json:"index_title" yaml:"index_title"`
		LinkPinyin bool   `json:"link_pinyin" yaml:"link_pinyin"`
		PrevNext   bool   `json:"prev_next" yaml:"prev_next"`
		Search     bool   `json:"search" yaml:"search"`
	}
)

var DefaultConfig = `# 网站名称
site_name: {{siteName}}
# 网站端口
site_addr: ":9000"
# 网站图标
favicon: 
# 网站作者
author: YOUNAME
# 源文档目录
docs_dir: docs
# 生成的HTML网站目录
site_dir: site
# 模板主题
theme: default
# 首页名称
index_title: 首页
# 是否将链接里的汉字转成拼音
link_pinyin: true
# 是否启用上一页下一页
prev_next: true
# 是否启用搜索
search: true`

func NewConfig() *Config {
	return &Config{}
}
