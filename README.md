# Modoc
***
将Markdown文件目录生成html网站

## 什么是Modoc

Modoc 是一个用于创建项目文档的 HTML 静态站点生成工具. 文档源码使用 Markdown 来撰写, 用 YAML 文件作为配置文档.

## 安装

 - 从源码安装
 - 直接下载可执行程序
    - Windows平台: [modoc-win.exe](modoc-win.exe)
    - Mac平台: [modoc-mac](modoc-mac)
    - Linux平台: [modoc-linux](modoc-linux)

## 开始

### 1.初始化项目
创建一个目录, 有如下结构
```
demo
└── docs
    └── index.md
```
- `demo` 为项目根目录
- `docs` 为存放文档的目录

在 `demo` 目录下, 执行初始化命令
```
modoc init
// output: 
// Initialization is complete, you can manually modify config file: config.yaml
// win下执行: modoc-win.exe init
```
自动生成`config.yaml`配置文件, 可以手动修改该文件以适配不同的需求
```
demo
├── config.yaml
└── docs
    └── index.md
```

### 2.生成导航菜单
将您的文档和目录按规划放到`docs`目录下.

执行以下命令会根据您的目录结构生成导航菜单:
```
modoc nav
// output:
// find: docs/index.md
// Generate the navigation configuration file: nav.yaml
```

### 3.开启服务
Modoc 包含了一个内建的服务器以预览当前文档. 

执行以下命令以启动内建服务器:
```
modoc serve
// output:
// build: site/index.html
// Start server: http://127.0.0.1:9000
```
访问 http://127.0.0.1:9000 即可预览网站.

### 4.发布项目
执行build命令, 在当前目录生成`site`文件夹.

将`site`文件夹部署到任何一个`web`服务器目录中即可, 不需要任何脚本环境支持.
```
modoc build
// output:
// build: site/index.html
```

## 配置
编辑`config.yaml`文件, 重新build或者重启serve.
```
# 网站名称
site_name: modoc
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
search: true
```

## 自定义导航菜单
Modoc会自动根据`docs`下的目录结构生成对应的导航菜单.

如果您需要定制菜单, 可以编辑 `nav.yaml`文件.

**注意**: 执行`modoc nav`命令会覆盖`nav.yaml`文件

```
title: ""
child:
- title: 把首页的名子改了HOME
  path: index.md
```