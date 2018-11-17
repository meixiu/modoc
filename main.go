package main

import "github.com/wbsifan/modoc/cmd"

//go:generate go-bindata -o asset/asset.go -pkg=asset template/...
func main() {
	cmd.Execute()
}
