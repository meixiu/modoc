package model

import (
	"encoding/json"
)

type (
	Skin struct {
		CustomCss string `json:"custom_css" yaml:"custom_css"`
		BaseCss   string `json:"base_css" yaml:"base_css"`
		ExtraCss  string `json:"extra_css" yaml:"extra_css"`
		HljsCss   string `json:"hljs_css" yaml:"hljs_css"`
	}

	Theme struct {
		Extend string
		Skin   map[string]*Skin
	}
)

func NewTheme() *Theme {
	return &Theme{}
}

func (this *Theme) GetSkinJSON() string {
	cbyte, _ := json.Marshal(this.Skin)
	return string(cbyte)
}
